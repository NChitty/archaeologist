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

type GatherActor struct {
	GoalItem     *artifactsmmo.ItemSchema
	GoalQuantity int
	character    *character.Character
}

var GatherSkills = [4]string{"mining", "woodcutting", "fishing", "alchemy"}

func NewGatherActor(character *character.Character, goalCode string, goalQuantity int) (*GatherActor, error) {
	item, err := artifacts.GetItem(goalCode)
	if err != nil {
		slog.Error("Could not retrieve item info:", err)
		return nil, err
	}

	if item.Subtype == "mob" {
		slog.Error("Fighting is not a currently supported operation")
		return nil, errors.ErrUnsupported
	}

	if !IsGatherable(item) {
		slog.Error("This item is not gatherable", "item", *item)
		return nil, errors.ErrUnsupported
	}

	slog.Info("Created gathering actor", "item", item, "qty", goalQuantity)

	return &GatherActor{GoalItem: item, GoalQuantity: goalQuantity, character: character}, nil
}

func IsGatherable(item *artifactsmmo.ItemSchema) bool {
	if item.Type != "resource" {
		return false
	}

	resources, err := artifacts.GetAllResources(nil, &item.Code)
	if err != nil {
		slog.Error("Failed to retrieve resources", "code", item.Code, "error", err)
		return false
	}

	if len(resources) == 0 {
		return false
	}

	return true
}

func (gatherActor *GatherActor) Do() error {
	resources, err := artifacts.GetAllResources(
		(*artifactsmmo.GatheringSkill)(&gatherActor.GoalItem.Subtype),
		&gatherActor.GoalItem.Code,
	)
	if err != nil {
		slog.Error("Could not retreive resource with given drop and skill.", err)
		return err
	}

	// NOTE: we might want to change this behavior given leveling (optimization todo)
	resourceCodeMap := make(map[string]artifactsmmo.DropRateSchema)
	for _, resource := range resources {
		for _, drop := range resource.Drops {
			if drop.Code == gatherActor.GoalItem.Code {
				resourceCodeMap[resource.Code] = drop
			}
		}
	}
	var resource string
	lowestRate := math.MaxInt
	for k, v := range resourceCodeMap {
		if v.Rate < lowestRate {
			resource = k
			lowestRate = v.Rate
		}
	}

	maps, err := artifacts.GetAllMaps(nil, &resource)
	if err != nil {
		slog.Error("Could not retrieve all map tiles potentially relevant to gathering resources.", err)
		return err
	}
	if len(maps) == 0 {
		slog.Error("No maps with type needed to gather resource.", "type", gatherActor.GoalItem.Type)
		return errors.New("No maps with subtype needed to gather resource.")
	}

	var x, y *int
	// todo another place for an optimizer
	for _, cell := range maps {
		x = &cell.X
		y = &cell.Y
		break
	}

	if x == nil || y == nil {
		slog.Error(
			"Could not try finding map cell for gathering",
			"itemCode",
			gatherActor.GoalItem.Code,
		)
		return errors.New("Could not find a map cell for gathering resource")
	}

	if gatherActor.character.Character.X != *x || gatherActor.character.Character.Y != *y {
		result, err := gatherActor.character.Move(*x, *y)
		if err != nil {
			return err
		}
		time.Sleep(result.CooldownRemaining)
	}

	for {
		_, err := gatherActor.character.Gather()
		if err != nil {
			slog.Error("Could not gather resource", err)
			return err
		}
		if gatherActor.character.GetInventory()[gatherActor.GoalItem.Code].Quantity == gatherActor.GoalQuantity {
			slog.Info("Finished gathering", "item", gatherActor.GoalItem, "qty", gatherActor.GoalQuantity)
			break
		}
	}

	return nil
}
