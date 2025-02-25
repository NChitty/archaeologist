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
	Gather(character *character.Character) (*ActionResult, error)
}

type GatherActor[T GatheringCharacterAdapter] struct {
	GoalItem         *artifactsmmo.ItemSchema
	GoalQuantity     int
	characterService T
	itemService      ItemAdapter
	mapService       MapAdapter
	logger           *slog.Logger
}

var GatherSkills = [4]string{"mining", "woodcutting", "fishing", "alchemy"}

func NewGatherActor[T GatheringCharacterAdapter](
	goalCode string,
	goalQuantity int,
	characterService T,
	itemService ItemAdapter,
	mapService MapAdapter,
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

	if !itemService.IsGatherable(item) {
		logger.Error("This item is not gatherable", "item", *item)
		return nil, errors.ErrUnsupported
	}

	logger.Info("Created gathering actor", "item", *item, "qty", goalQuantity)

	return &GatherActor[GatheringCharacterAdapter]{
		GoalItem:         item,
		GoalQuantity:     goalQuantity,
		characterService: characterService,
		itemService:      itemService,
		mapService:       mapService,
		logger:           logger}, nil
}

func (actor *GatherActor[GatheringCharacterAdapter]) Do(character *character.Character) error {
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
