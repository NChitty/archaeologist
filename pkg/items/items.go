package items

import (
	"context"
	"errors"
	"log/slog"
	"sync"
	"time"

	"github.com/NChitty/archaeologist/pkg/characters"
	artifactsErrors "github.com/NChitty/archaeologist/pkg/errors"
	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
)

type ItemService struct {
	client *artifactsmmo.ClientWithResponses
	logger *slog.Logger
}

func NewItemService(logger *slog.Logger, server string, opts ...artifactsmmo.ClientOption) (*ItemService, error) {
	client, err := artifactsmmo.NewClientWithResponses("https://api.artifactsmmo.com/", opts...)
	if err != nil {
		logger.Error("Could not create authentication-less client", "error", err)
		return nil, err
	}

	return &ItemService{
		client: client,
		logger: logger,
	}, nil
}

var defaultItemService *ItemService

func DefaultItemService() *ItemService {
	if defaultItemService == nil {
		service, err := NewItemService(slog.Default(), "https://api.artifactsmmo.com/")
		for err != nil {
			service, err = NewItemService(slog.Default(), "https://api.artifactsmmo.com/")
		}
		defaultItemService = service
	}
	return defaultItemService
}

func (service *ItemService) GetItem(name string) (*artifactsmmo.ItemSchema, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	itemResp, err := service.client.GetItemItemsCodeGetWithResponse(ctx, name)
	if err != nil {
		service.logger.Error("Could not retrieve item", "error", err)
		return nil, err
	}
	if itemResp.StatusCode() == artifactsErrors.NotFound {
		service.logger.Error("Item not found", "code", name, "response", string(itemResp.Body))
		return nil, errors.New("Item not found.")
	}

	return &itemResp.JSON200.Data, nil
}

func (service *ItemService) GetAllResources(skill *artifactsmmo.GatheringSkill, code *string) ([]artifactsmmo.ResourceSchema, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resourceResp, err := service.client.GetAllResourcesResourcesGetWithResponse(
		ctx,
		&artifactsmmo.GetAllResourcesResourcesGetParams{
			MinLevel: nil,
			MaxLevel: nil,
			Skill:    skill,
			Drop:     code,
			Page:     nil,
			Size:     nil,
		},
	)
	if err != nil {
		service.logger.Error("Could not retrieve gather resource", "error", err)
		return nil, err
	}
	if results, err := resourceResp.JSON200.Total.AsDataPageResourceSchemaTotal0(); err != nil {
		return nil, err
	} else if results < 1 {
		return nil, errors.New("No results.")
	}

	return resourceResp.JSON200.Data, nil
}

type equipment struct {
	code     string
	quantity int
}

func (service *ItemService) GetCharacterEquipment(character *characters.Character) *[]Equipment {
	equipmentCodes := []equipment{}
	equipmentCodes = append(equipmentCodes, equipment{character.Character.WeaponSlot, 1})
	equipmentCodes = append(equipmentCodes, equipment{character.Character.ShieldSlot, 1})
	equipmentCodes = append(equipmentCodes, equipment{character.Character.HelmetSlot, 1})
	equipmentCodes = append(equipmentCodes, equipment{character.Character.BodyArmorSlot, 1})
	equipmentCodes = append(equipmentCodes, equipment{character.Character.LegArmorSlot, 1})
	equipmentCodes = append(equipmentCodes, equipment{character.Character.BootsSlot, 1})
	equipmentCodes = append(equipmentCodes, equipment{character.Character.Ring1Slot, 1})
	equipmentCodes = append(equipmentCodes, equipment{character.Character.Ring2Slot, 1})
	equipmentCodes = append(equipmentCodes, equipment{character.Character.AmuletSlot, 1})
	equipmentCodes = append(equipmentCodes, equipment{character.Character.Artifact1Slot, 1})
	equipmentCodes = append(equipmentCodes, equipment{character.Character.Artifact2Slot, 1})
	equipmentCodes = append(equipmentCodes, equipment{character.Character.Artifact3Slot, 1})
	equipmentCodes = append(equipmentCodes, equipment{
		character.Character.Utility1Slot,
		character.Character.Utility1SlotQuantity,
	})
	equipmentCodes = append(equipmentCodes, equipment{
		character.Character.Utility2Slot,
		character.Character.Utility2SlotQuantity,
	})

	equipmentSlots := []Equipment{}

	var wg sync.WaitGroup
	var lock sync.Mutex
	for _, equipment := range equipmentCodes {
		if equipment.code == "" {
			continue
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			item, err := service.GetItem(equipment.code)
			if err != nil {
				service.logger.Warn("Could not retrieve item", "code", equipment.code, "error", err)
				return
			}
			lock.Lock()
			equipmentSlots = append(equipmentSlots, Equipment{
				Item:     *item,
				Quantity: equipment.quantity,
			})
			lock.Unlock()
		}()
	}
	wg.Wait()
	return &equipmentSlots
}

type Equipment struct {
	Item     artifactsmmo.ItemSchema
	Quantity int
}
