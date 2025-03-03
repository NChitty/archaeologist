package actors

import (
	"errors"
	"log/slog"
	"time"

	"github.com/NChitty/archaeologist/pkg/models/character"
	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
)

type CraftingCharacterAdapter interface {
	CharacterAdapter
	GatheringCharacterAdapter
	Craft(character *character.Character, itemCode string, quantity int) (*ActionResult, error)
}

type CraftingActor[C CraftingCharacterAdapter] struct {
	GoalItem         *artifactsmmo.ItemSchema
	GoalQuantity     int
	craftingRecipe   *artifactsmmo.CraftSchema
	characterService CraftingCharacterAdapter
	itemService      ItemAdapter
	mapService       MapAdapter
	config           *Config[C]
	logger           *slog.Logger
}

var workshop string = "workshop"

func NewCraftingActor[C CraftingCharacterAdapter](
	goalCode string,
	goalQuantity int,
	config *Config[C],
) (*CraftingActor[C], error) {
	item, err := config.itemAdapter.GetItem(goalCode)
	if err != nil {
		config.logger.Error("Could not retrieve item info", "error", err)
		return nil, err
	}

	actor := &CraftingActor[C]{
		GoalItem:         item,
		GoalQuantity:     goalQuantity,
		characterService: config.characterAdapter,
		itemService:      config.itemAdapter,
		mapService:       config.mapAdapter,
		config:           config,
		logger:           config.logger,
	}

	isCraftable := actor.itemService.IsCraftable(item)

	if !isCraftable {
		actor.logger.Error("Item is not craftable", "item", *item)
		return nil, errors.ErrUnsupported
	}

	craftingRecipe, err := item.Craft.AsCraftSchema()
	if err != nil {
		actor.logger.Error("Could not parse crafting recipe", "content", item.Craft)
		return nil, err
	}
	actor.craftingRecipe = &craftingRecipe

	actor.logger.Info("Created crafting actor", "item", *actor.GoalItem, "qty", actor.GoalQuantity)

	return actor, nil
}

func (actor *CraftingActor[C]) Do(character *character.Character) error {
	actor.characterService.UpdateCharacter(character)
	for _, item := range *actor.craftingRecipe.Items {
		slot, isPresent := character.Inventory[item.Code]
		needed := item.Quantity * actor.GoalQuantity / *actor.craftingRecipe.Quantity
		actor.logger.Debug("Checking prereq", "code", item.Code, "qtyNeeded", needed)

		itemSchema, err := actor.itemService.GetItem(item.Code)
		if err != nil {
			actor.logger.Error("Could not retrieve item info", "error", err)
			return err
		}

		if isPresent && needed <= slot.Quantity {
      continue
		}

		isGatherable := actor.itemService.IsGatherable(itemSchema)
		isCraftable := actor.itemService.IsCraftable(itemSchema)
		if !isCraftable && !isGatherable {
			actor.logger.Error("Item is not attainable through gathering or crafting", "item", *itemSchema)
			return errors.ErrUnsupported
		}

		if isGatherable && isCraftable {
			// todo optimizer
			prereqActor, err := NewGatherActor(character, item.Code, needed, actor.config)
			if err != nil {
				return err
			}
			err = prereqActor.Do(character)
			if err != nil {
				return err
			}
			continue
		}

		if isGatherable && !isCraftable {
			prereqActor, err := NewGatherActor(character, item.Code, needed, actor.config)
			if err != nil {
				return err
			}
			err = prereqActor.Do(character)
			if err != nil {
				return err
			}
			continue
		}

		if isPresent && needed > slot.Quantity {
			needed -= slot.Quantity
		}

		prereqActor, err := NewCraftingActor(item.Code, needed, actor.config)
		if err != nil {
			return err
		}

		err = prereqActor.Do(character)
		if err != nil {
			return err
		}
	}

	actor.logger.Debug("Finished gathering prereqs", "item", *actor.GoalItem)
	err := move(actor.mapService, actor.characterService, character, &workshop, (*string)(actor.craftingRecipe.Skill))
	if err != nil {
		actor.logger.Error("Could not move character", "error", err)
		return err
	}

	craftRes, err := actor.characterService.Craft(character, actor.GoalItem.Code, actor.GoalQuantity)
	if err != nil {
		actor.logger.Error("Failed to craft")
		return err
	}
	actor.logger.Info("Finished crafting", "item", actor.GoalItem.Code, "qty", actor.GoalQuantity)
	time.Sleep(craftRes.CooldownRemaining)

	return nil
}
