package actors

import (
	"errors"
	"log/slog"
	"strings"
	"time"

	character "github.com/NChitty/archaeologist/pkg"
	"github.com/NChitty/archaeologist/pkg/artifacts"
	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
)

type GatherActor struct {
	GoalItem     *artifactsmmo.ItemSchema
	GoalQuantity int
	character    *character.Character
}

func New(character *character.Character, goalCode string, goalQuantity int) (*GatherActor, error) {
	item, err := artifacts.GetItem(goalCode)
	if err != nil {
		slog.Error("Could not retrieve item info:", err)
		return nil, err
	}

	if item.Craft != nil {
		slog.Error("This item is not gatherable.")
		return nil, errors.New(goalCode + " is not gatherable")
	}

	return &GatherActor{GoalItem: item, GoalQuantity: goalQuantity, character: character}, nil
}

func (gatherActor *GatherActor) Do() error {
	inventoryMap, err := gatherActor.character.GetInventory()
	if err != nil {
		slog.Error("Could not retrieve character info:", err)
		return err
	}

	contentType, err := getRelevantContentType(*gatherActor.GoalItem)
	resourceType, _, _ := strings.Cut(gatherActor.GoalItem.Code, "_")
	if err != nil {
		slog.Error("Error getting content type for map query:", err)
		return err
	}

	maps, err := artifacts.GetAllMaps(contentType, &resourceType)
	if err != nil {
		slog.Error("Could not retrieve all map tiles potentially relevant to gathering resources.", err)
	}
	if len(maps) == 0 {
		slog.Error("No maps with type needed to gather resource.", "type", gatherActor.GoalItem.Type)
		return errors.New("No maps with subtype needed to gather resource.")
	}

	var x, y *int
	for _, cell := range maps {
		cellContent, err := cell.Content.AsMapContentSchema()
		if err != nil {
			slog.Error("Could not parse map content schema:", err)
			return err
		}

		if strings.Contains(cellContent.Code, resourceType) {
			x = &cell.X
			y = &cell.Y
			break
		}
	}

	if x == nil || y == nil {
		slog.Error(
			"Could not try finding map cell for gathering",
			"itemCode",
			gatherActor.GoalItem.Code,
		)
		return errors.New("Could not find a map cell for gathering resource")
	}

	character, err := gatherActor.character.GetCharacter()
	if err != nil {
		return err
	}

	if character.X != *x || character.Y != *y {
		result, err := gatherActor.character.Move(*x, *y)
		if err != nil {
			return err
		}
		time.Sleep(result.CooldownRemaining)
	}

	needed := gatherActor.GoalQuantity
	slot, exists := inventoryMap[gatherActor.GoalItem.Code]
	if exists {
		needed = gatherActor.GoalQuantity - slot.Quantity
	}

	if strings.Contains(*contentType, "resource") {
		for i := 0; i < needed; i++ {
			slog.Debug("Gathering resource", "gatherActionsTaken", i, "needed", needed)
			actionResult, err := gatherActor.character.Gather()
			if err != nil {
				slog.Error("Could not gather resource", err)
				return err
			}
			time.Sleep(actionResult.CooldownRemaining)
		}
	}

	return nil
}

func getRelevantContentType(item artifactsmmo.ItemSchema) (*string, error) {
	var contentType string
	switch item.Subtype {
	case "mob":
		contentType = "monster"
	case "woodcutting":
		contentType = "resource"
	case "mining":
		contentType = "resource"
	case "fishing":
		contentType = "resource"
	default:
		contentType = ""
	}

	if len(contentType) == 0 {
		return nil, errors.New("Unsupported subtype provided: " + item.Subtype)
	}

	return &contentType, nil
}
