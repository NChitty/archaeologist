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

type GatherActor struct {
	GoalItem         *artifactsmmo.ItemSchema
	GoalQuantity     int
	characterService *characters.CharacterService
	itemService      items.ItemAccessor
	mapService       maps.MapAccessor
	logger           *slog.Logger
}

var GatherSkills = [4]string{"mining", "woodcutting", "fishing", "alchemy"}

func NewGatherActor(
	goalCode string,
	goalQuantity int,
	characterService *characters.CharacterService,
	itemService items.ItemAccessor,
	mapService maps.MapAccessor,
	logger *slog.Logger,
) (Actor, error) {
	item, err := itemService.GetItem(goalCode)
	if err != nil {
		logger.Error("Could not retrieve item info", "error", err)
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

	logger.Info("Created gathering actor", "item", *item, "qty", goalQuantity)

	return &GatherActor{
		GoalItem:         item,
		GoalQuantity:     goalQuantity,
		characterService: characterService,
		itemService:      itemService,
		mapService:       mapService,
		logger:           logger}, nil
}

func IsGatherable(logger *slog.Logger, item *artifactsmmo.ItemSchema) bool {
	if item.Type != "resource" {
		return false
	}

	resources, err := items.DefaultItemService().GetAllResources(nil, &item.Code)
	if err != nil {
		logger.Warn("Failed to retrieve resources", "code", item.Code, "error", err)
		return false
	}

	if len(resources) == 0 {
		return false
	}

	return true
}

func (actor *GatherActor) Do(character *characters.CharacterWrapper) error {
	resources, err := actor.itemService.GetAllResources(
		(*artifactsmmo.GatheringSkill)(&actor.GoalItem.Subtype),
		&actor.GoalItem.Code,
	)
	if err != nil {
		actor.logger.Error("Could not retreive resource with given drop and skill.", "error", err)
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
	actor.logger.Info("Found resource", "code", resource)

	maps, err := actor.mapService.GetAllMaps(nil, &resource, nil, nil)
	if err != nil {
		actor.logger.Error("Could not retrieve all map tiles potentially relevant to gathering resources.", "error", err)
		return err
	}
	if len(maps) == 0 {
		actor.logger.Error("No maps with type needed to gather resource.", "type", actor.GoalItem.Type)
		return errors.New("No maps with subtype needed to gather resource.")
	}

	var x, y int
	// todo another place for an optimizer
	for _, cell := range maps {
		x = cell.X
		y = cell.Y
		break
	}

	if x == 0 && y == 0 {
		actor.logger.Error(
			"Could not find map cell for gathering",
			"itemCode",
			actor.GoalItem.Code,
		)
		return errors.New("Could not find a map cell for gathering resource")
	}

	if character.X != x || character.Y != y {
		result, err := actor.characterService.Move(character, x, y)
		if err != nil {
			return err
		}
		time.Sleep(result.CooldownRemaining)
	}

	for {
		_, err := actor.characterService.Gather(character)
		if err != nil {
			actor.logger.Error("Could not gather resource", "error", err)
			return err
		}
		actor.logger.Info(
			"Gathered resource",
			"item",
			actor.GoalItem.Code,
			"goal",
			actor.GoalQuantity,
			"qty",
			character.Inventory[actor.GoalItem.Code].Quantity,
		)
		if character.Inventory[actor.GoalItem.Code].Quantity == actor.GoalQuantity {
			actor.logger.Info("Finished gathering", "item", actor.GoalItem.Code, "qty", actor.GoalQuantity)
			break
		}
	}

	return nil
}
