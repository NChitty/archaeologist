package actors

import (
	"errors"
	"log/slog"
	"math"
	"time"

	"github.com/NChitty/archaeologist/pkg/artifacts"
	"github.com/NChitty/archaeologist/pkg/character"
	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
)

type CraftingActor struct {
	GoalItem       *artifactsmmo.ItemSchema
	GoalQuantity   int
	craftingRecipe *artifactsmmo.CraftSchema
	character      *character.Character
}

var workshop string = "workshop"

func NewCraftingActor(character *character.Character, goalCode string, goalQuantity int) (*CraftingActor, error) {
	item, err := artifacts.GetItem(goalCode)
	if err != nil {
		slog.Error("Could not retrieve item info:", err)
		return nil, err
	}

	actor := &CraftingActor{
		GoalItem:     item,
		GoalQuantity: goalQuantity,
		character:    character,
	}

	isCraftable := IsCraftable(item)

	if !isCraftable {
		return nil, errors.ErrUnsupported
	}

	craftingRecipe, err := item.Craft.AsCraftSchema()
	if err != nil {
		slog.Error("Could not parse crafting recipe", "content", item.Craft)
		return nil, err
	}
	actor.craftingRecipe = &craftingRecipe

	slog.Info("Created crafting actor", "item", actor.GoalItem, "qty", actor.GoalQuantity)

	return actor, nil
}

func IsCraftable(item *artifactsmmo.ItemSchema) bool {
	if item.Craft == nil {
		return false
	}
	return true
}

func (actor *CraftingActor) Do() error {
	for _, item := range *actor.craftingRecipe.Items {
		slot, isPresent := actor.character.GetInventory()[item.Code]
		needed := item.Quantity * actor.GoalQuantity / *actor.craftingRecipe.Quantity
		if isPresent && needed > slot.Quantity {
			slog.Debug("Do not have sufficient material", "item", item.Code, "qtyNeeded", needed, "qtyHave", slot.Quantity)
			itemSchema, err := artifacts.GetItem(item.Code)
			if err != nil {
				slog.Error("Could not retrieve item info:", err)
				return err
			}
			if !IsCraftable(itemSchema) && !IsGatherable(itemSchema) {
				return errors.ErrUnsupported
			}
			if IsGatherable(itemSchema) && IsCraftable(itemSchema) {
				// todo optimizer
				prereqActor, err := NewGatherActor(actor.character, item.Code, needed)
				if err != nil {
					return err
				}
				err = prereqActor.Do()
				if err != nil {
					return err
				}
			}
			if IsGatherable(itemSchema) {
				prereqActor, err := NewGatherActor(actor.character, item.Code, needed)
				if err != nil {
					return err
				}
				err = prereqActor.Do()
				if err != nil {
					return err
				}
			}
			prereqActor, err := NewCraftingActor(actor.character, item.Code, needed - slot.Quantity)
			if err != nil {
				return err
			}
			err = prereqActor.Do()
			if err != nil {
				return err
			}
		}
	}

	maps, err := artifacts.GetAllMaps(&workshop, (*string)(actor.craftingRecipe.Skill))
	if err != nil {
		slog.Error("Could not retrieve all map tiles potentially relevant to gathering resources.", err)
		return err
	}
	if len(maps) == 0 {
		slog.Error("No maps with type needed to craft resource.", "skill", actor.craftingRecipe.Skill)
		return errors.New("No maps with subtype needed to gather resource.")
	}
	charX, charY := actor.character.GetLocation()
	minDist := math.MaxInt
	var dest *artifactsmmo.MapSchema
	for _, cell := range maps {
		distX := charX - cell.X
		distY := charY - cell.Y
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

	moveRes, err := actor.character.Move(dest.X, dest.Y)
	if err != nil {
		return err
	}
	time.Sleep(moveRes.CooldownRemaining)

	neededQty := actor.GoalQuantity
	slot, isPresent := actor.character.GetInventory()[actor.GoalItem.Code]
	if isPresent {
		neededQty = actor.GoalQuantity - slot.Quantity
	}

	craftRes, err := actor.character.Craft(actor.GoalItem.Code, neededQty)
	if err != nil {
		slog.Error("Failed to craft")
		return err
	}
	time.Sleep(craftRes.CooldownRemaining)

	return nil
}
