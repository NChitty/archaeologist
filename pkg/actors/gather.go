package actors

import (
	"errors"
	"log/slog"
	"math"
	"slices"
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

func New(character *character.Character, goalCode string, goalQuantity int) (*GatherActor, error) {
	item, err := artifacts.GetItem(goalCode)
	if err != nil {
		slog.Error("Could not retrieve item info:", err)
		return nil, err
	}

	if item.Type != "resource" {
		slog.Error("Item is not a resource", "type", item.Type, "subtype", item.Subtype)
		return nil, errors.New("Item is not resource.")
	}

	if item.Subtype == "mob" {
		return nil, errors.ErrUnsupported
	}

	if !slices.Contains(GatherSkills[:], item.Subtype) {
		return nil, errors.ErrUnsupported
	}

	return &GatherActor{GoalItem: item, GoalQuantity: goalQuantity, character: character}, nil
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

	inventoryMap, err := gatherActor.character.GetInventory()
	if err != nil {
		slog.Error("Could not retrieve character info:", err)
		return err
	}

	needed := gatherActor.GoalQuantity
	slot, exists := inventoryMap[gatherActor.GoalItem.Code]
	if exists {
		needed = gatherActor.GoalQuantity - slot.Quantity
	}

	for i := 0; i < needed; i++ {
		slog.Debug("Gathering resource", "gatherActionsTaken", i, "needed", needed)
		actionResult, err := gatherActor.character.Gather()
		if err != nil {
			slog.Error("Could not gather resource", err)
			return err
		}
		time.Sleep(actionResult.CooldownRemaining)
	}

	return nil
}
