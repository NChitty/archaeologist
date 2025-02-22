package actors

import (
	"errors"
	"log/slog"
	"math"
	"time"

	"github.com/NChitty/archaeologist/pkg/characters"
	"github.com/NChitty/archaeologist/pkg/items"
	"github.com/NChitty/archaeologist/pkg/maps"
	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
)

type CraftingActor struct {
	GoalItem         *artifactsmmo.ItemSchema
	GoalQuantity     int
	craftingRecipe   *artifactsmmo.CraftSchema
	characterService *characters.CharacterService
	itemService      *items.ItemService
	logger           *slog.Logger
}

var workshop string = "workshop"

func NewCraftingActor(
	goalCode string,
	goalQuantity int,
	characterService *characters.CharacterService,
	itemService *items.ItemService,
	logger *slog.Logger,
) (*CraftingActor, error) {
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
		logger:           logger,
	}

	isCraftable := IsCraftable(item)

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

func IsCraftable(item *artifactsmmo.ItemSchema) bool {
	if item.Craft == nil {
		return false
	}
	return true
}

func (actor *CraftingActor) Do(character *characters.CharacterWrapper) error {
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
			if !IsCraftable(itemSchema) && !IsGatherable(actor.logger, itemSchema) {
				actor.logger.Error("Item is not attainable through gathering or crafting", "item", *itemSchema)
				return errors.ErrUnsupported
			}
			if IsGatherable(actor.logger, itemSchema) && IsCraftable(itemSchema) {
				// todo optimizer
				prereqActor, err := NewGatherActor(
					item.Code,
					needed,
					actor.characterService,
					actor.itemService,
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
			if IsGatherable(actor.logger, itemSchema) && !IsCraftable(itemSchema) {
				prereqActor, err := NewGatherActor(
					item.Code,
					needed,
					actor.characterService,
					actor.itemService,
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
			if !IsCraftable(itemSchema) && !IsGatherable(actor.logger, itemSchema) {
				actor.logger.Error("This item is not gatherable", "item", *itemSchema)
				return errors.ErrUnsupported
			}
			if IsGatherable(actor.logger, itemSchema) && IsCraftable(itemSchema) {
				// todo optimizer
				prereqActor, err := NewGatherActor(
					item.Code,
					needed,
					actor.characterService,
					actor.itemService,
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
			if IsGatherable(actor.logger, itemSchema) && !IsCraftable(itemSchema) {
				prereqActor, err := NewGatherActor(
					item.Code,
					needed,
					actor.characterService,
					actor.itemService,
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

	maps, err := maps.GetAllMaps(actor.logger, &workshop, (*string)(actor.craftingRecipe.Skill))
	if err != nil {
		actor.logger.Error("Could not retrieve all map tiles potentially relevant to gathering resources.", "error", err)
		return err
	}
	if len(maps) == 0 {
		actor.logger.Error("No maps with type needed to craft resource.", "skill", actor.craftingRecipe.Skill)
		return errors.New("No maps with subtype needed to gather resource.")
	}
	minDist := math.MaxInt
	var dest *artifactsmmo.MapSchema
	for _, cell := range maps {
		distX := character.X - cell.X
		distY := character.Y - cell.Y
		if distX < 0 {
			distX *= -1
		}
		if distY < 0 {
			distY *= -1
		}
		if minDist > (distX + distY) {
			minDist = distX + distY
			dest = &cell
		}
	}

	moveRes, err := actor.characterService.Move(character, dest.X, dest.Y)
	if err != nil {
		return err
	}
	time.Sleep(moveRes.CooldownRemaining)

	craftRes, err := actor.characterService.Craft(character, actor.GoalItem.Code, actor.GoalQuantity)
	if err != nil {
		actor.logger.Error("Failed to craft")
		return err
	}
	time.Sleep(craftRes.CooldownRemaining)

	return nil
}
