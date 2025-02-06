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
	logger       *slog.Logger
}

var GatherSkills = [4]string{"mining", "woodcutting", "fishing", "alchemy"}

func NewGatherActor(logger *slog.Logger, character *character.Character, goalCode string, goalQuantity int) (*GatherActor, error) {
	item, err := artifacts.GetItem(logger, goalCode)
	if err != nil {
		logger.Error("Could not retrieve item info:", err)
		return nil, err
	}

	if item.Subtype == "mob" {
		logger.Error("Fighting is not a currently supported operation")
		return nil, errors.ErrUnsupported
	}

	if !IsGatherable(logger, item) {
		logger.Error("This item is not gatherable", "item", *item)
		return nil, errors.ErrUnsupported
	}

	logger.Info("Created gathering actor", "item", item, "qty", goalQuantity)

	return &GatherActor{GoalItem: item, GoalQuantity: goalQuantity, character: character, logger: logger}, nil
}

func IsGatherable(logger *slog.Logger, item *artifactsmmo.ItemSchema) bool {
	if item.Type != "resource" {
		return false
	}

	resources, err := artifacts.GetAllResources(logger, nil, &item.Code)
	if err != nil {
		logger.Error("Failed to retrieve resources", "code", item.Code, "error", err)
		return false
	}

	if len(resources) == 0 {
		return false
	}

	return true
}

func (actor *GatherActor) Do() error {
	resources, err := artifacts.GetAllResources(
		actor.logger,
		(*artifactsmmo.GatheringSkill)(&actor.GoalItem.Subtype),
		&actor.GoalItem.Code,
	)
	if err != nil {
		actor.logger.Error("Could not retreive resource with given drop and skill.", err)
		return err
	}

	// NOTE: we might want to change this behavior given leveling (optimization todo)
	resourceCodeMap := make(map[string]artifactsmmo.DropRateSchema)
	for _, resource := range resources {
		for _, drop := range resource.Drops {
			if drop.Code == actor.GoalItem.Code {
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

	maps, err := artifacts.GetAllMaps(actor.logger, nil, &resource)
	if err != nil {
		actor.logger.Error("Could not retrieve all map tiles potentially relevant to gathering resources.", err)
		return err
	}
	if len(maps) == 0 {
		actor.logger.Error("No maps with type needed to gather resource.", "type", actor.GoalItem.Type)
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
		actor.logger.Error(
			"Could not try finding map cell for gathering",
			"itemCode",
			actor.GoalItem.Code,
		)
		return errors.New("Could not find a map cell for gathering resource")
	}

	if actor.character.Character.X != *x || actor.character.Character.Y != *y {
		result, err := actor.character.Move(*x, *y)
		if err != nil {
			return err
		}
		time.Sleep(result.CooldownRemaining)
	}

	for {
		_, err := actor.character.Gather()
		if err != nil {
			actor.logger.Error("Could not gather resource", err)
			return err
		}
		if actor.character.GetInventory()[actor.GoalItem.Code].Quantity == actor.GoalQuantity {
			actor.logger.Info("Finished gathering", "item", actor.GoalItem, "qty", actor.GoalQuantity)
			break
		}
	}

	return nil
}
