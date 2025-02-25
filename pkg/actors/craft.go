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

type CraftingActor struct {
	GoalItem         *artifactsmmo.ItemSchema
	GoalQuantity     int
	craftingRecipe   *artifactsmmo.CraftSchema
	characterService CraftingCharacterAdapter
	itemService      ItemAdapter
	mapService       MapAdapter
	logger           *slog.Logger
}

var workshop string = "workshop"

func NewCraftingActor(
	goalCode string,
	goalQuantity int,
	characterService CraftingCharacterAdapter,
	itemService ItemAdapter,
	mapService MapAdapter,
	logger *slog.Logger,
) (Actor, error) {
	item, err := itemService.GetItem(goalCode)
	if err != nil {
		logger.Error("Could not retrieve item info", "error", err)
		return nil, err
	}

	actor := &CraftingActor{
		GoalItem:         item,
		GoalQuantity:     goalQuantity,
		characterService: characterService,
		itemService:      itemService,
		mapService:       mapService,
		logger:           logger,
	}

	isCraftable := itemService.IsCraftable(item)

	if !isCraftable {
		logger.Error("Item is not craftable", "item", *item)
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

func (actor *CraftingActor) Do(character *character.Character) error {
	actor.characterService.UpdateCharacter(character)
	for _, item := range *actor.craftingRecipe.Items {
		slot, isPresent := character.Inventory[item.Code]
		needed := item.Quantity * actor.GoalQuantity / *actor.craftingRecipe.Quantity
		actor.logger.Debug("Checking prereq", "code", item.Code, "qtyNeeded", needed)
		if isPresent && needed > slot.Quantity {
			actor.logger.Debug("Do not have sufficient material", "item", item.Code, "qtyNeeded", needed, "qtyHave", slot.Quantity)
			itemSchema, err := actor.itemService.GetItem(item.Code)
			if err != nil {
				actor.logger.Error("Could not retrieve item info", "error", err)
				return err
			}
			if !actor.itemService.IsCraftable(itemSchema) &&
				!actor.itemService.IsGatherable(itemSchema) {
				actor.logger.Error("Item is not attainable through gathering or crafting", "item", *itemSchema)
				return errors.ErrUnsupported
			}
			if actor.itemService.IsGatherable(itemSchema) &&
				actor.itemService.IsCraftable(itemSchema) {
				// todo optimizer
				prereqActor, err := NewGatherActor(
					item.Code,
					needed,
					actor.characterService,
					actor.itemService,
					actor.mapService,
					actor.logger,
				)
				if err != nil {
					return err
				}
				err = prereqActor.Do(character)
				if err != nil {
					return err
				}
				continue
			}
			if actor.itemService.IsGatherable(itemSchema) &&
				!actor.itemService.IsCraftable(itemSchema) {
				prereqActor, err := NewGatherActor(
					item.Code,
					needed,
					actor.characterService,
					actor.itemService,
					actor.mapService,
					actor.logger,
				)
				if err != nil {
					return err
				}
				err = prereqActor.Do(character)
				if err != nil {
					return err
				}
				continue
			}
			prereqActor, err := NewCraftingActor(item.Code,
				needed-slot.Quantity,
				actor.characterService,
				actor.itemService,
				actor.mapService,
				actor.logger,
			)
			if err != nil {
				return err
			}
			err = prereqActor.Do(character)
			if err != nil {
				return err
			}
		} else if !isPresent {
			actor.logger.Debug("Do not have sufficient material", "item", item.Code, "qtyNeeded", needed, "qtyHave", 0)
			itemSchema, err := actor.itemService.GetItem(item.Code)
			if err != nil {
				actor.logger.Error("Could not retrieve item info", "error", err)
				return err
			}
			if !actor.itemService.IsCraftable(itemSchema) &&
				!actor.itemService.IsGatherable(itemSchema) {
				actor.logger.Error("This item is not gatherable", "item", *itemSchema)
				return errors.ErrUnsupported
			}
			if actor.itemService.IsGatherable(itemSchema) &&
				actor.itemService.IsCraftable(itemSchema) {
				// todo optimizer
				prereqActor, err := NewGatherActor(
					item.Code,
					needed,
					actor.characterService,
					actor.itemService,
					actor.mapService,
					actor.logger,
				)
				if err != nil {
					return err
				}
				err = prereqActor.Do(character)
				if err != nil {
					return err
				}
				continue
			}
			if actor.itemService.IsGatherable(itemSchema) &&
				!actor.itemService.IsCraftable(itemSchema) {
				prereqActor, err := NewGatherActor(
					item.Code,
					needed,
					actor.characterService,
					actor.itemService,
					actor.mapService,
					actor.logger,
				)
				if err != nil {
					return err
				}
				err = prereqActor.Do(character)
				if err != nil {
					return err
				}
				continue
			}
			prereqActor, err := NewCraftingActor(
				item.Code,
				needed-slot.Quantity,
				actor.characterService,
				actor.itemService,
				actor.mapService,
				actor.logger,
			)
			if err != nil {
				return err
			}
			err = prereqActor.Do(character)
			if err != nil {
				return err
			}
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
