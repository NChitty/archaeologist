package actors

import (
	"errors"
	"log/slog"
	"math"

	"github.com/NChitty/archaeologist/pkg/models/character"
	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
)

type GatheringCharacterAdapter interface {
	CharacterAdapter
	FightingCharacterAdapter
	Gather(character *character.Character) (*ActionResult, error)
}

type GatherActor struct {
	GoalItem         *artifactsmmo.ItemSchema
	GoalQuantity     int
	characterService GatheringCharacterAdapter
	itemService      ItemAdapter
	mapService       MapAdapter
	logger           *slog.Logger
}

var GatherSkills = [4]string{"mining", "woodcutting", "fishing", "alchemy"}

func NewGatherActor[C GatheringCharacterAdapter](
	character *character.Character,
	goalCode string,
	goalQuantity int,
	config *Config[C],
) (Actor, error) {
	item, err := config.itemAdapter.GetItem(goalCode)
	if err != nil {
		config.logger.Error("Could not retrieve item info", "error", err)
		return nil, err
	}

	if item.Subtype == "mob" {
		return NewDropsFightingActor(character, goalCode, goalQuantity, config)
	}

	if !config.itemAdapter.IsGatherable(item) {
		config.logger.Error("This item is not gatherable", "item", *item)
		return nil, errors.ErrUnsupported
	}

	config.logger.Info("Created gathering actor", "item", *item, "qty", goalQuantity)

	return &GatherActor{
		GoalItem:         item,
		GoalQuantity:     goalQuantity,
		characterService: config.characterAdapter,
		itemService:      config.itemAdapter,
		mapService:       config.mapAdapter,
		logger:           config.logger,
	}, nil
}

func (actor *GatherActor) Do(character *character.Character) error {
	actor.characterService.UpdateCharacter(character)
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

	err = move(actor.mapService, actor.characterService, character, nil, &resource)
	if err != nil {
		actor.logger.Error("Could not move character", "error", err)
		return err
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
