package characters

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/NChitty/archaeologist/pkg/items"
	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
)

type Character struct {
	Name      string
	Character *artifactsmmo.CharacterSchema
	client    *artifactsmmo.ClientWithResponses
	logger    *slog.Logger
}

func New(logger *slog.Logger, client *artifactsmmo.ClientWithResponses, name string) (*Character, error) {
	character := &Character{
		Name:   name,
		client: client,
		logger: logger,
	}
	_, err := character.UpdateCharacter()
	return character, err
}

func (character *Character) UpdateCharacter() (*artifactsmmo.CharacterSchema, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	characterResp, err := character.client.GetCharacterCharactersNameGetWithResponse(
		ctx,
		character.Name,
	)
	if err != nil {
		character.logger.Error("Could not retrieve character", "error", err)
		return nil, err
	}
	if characterResp.StatusCode() == 404 {
		character.logger.Error("Could not retrieve character", "name", character.Name)
		return nil, errors.New("Could not retrieve character with name: " + character.Name)
	}

	character.Character = &characterResp.JSON200.Data

	return &characterResp.JSON200.Data, nil
}

func (character *Character) GetInventory() map[string]artifactsmmo.InventorySlot {
	inventoryMap := make(map[string]artifactsmmo.InventorySlot)
	for _, slot := range *character.Character.Inventory {
		inventoryMap[slot.Code] = slot
	}
	return inventoryMap
}

func (character *Character) WaitCooldown() error {
	charSchema, err := character.UpdateCharacter()
	if err != nil {
		return err
	}
	character.Character = charSchema

	if time.Now().Before(*charSchema.CooldownExpiration) {
		character.logger.Debug(
			"Waiting for cooldown",
			"expiration", charSchema.CooldownExpiration,
			"timeRemaining", charSchema.CooldownExpiration.Sub(time.Now()),
		)
		time.Sleep(charSchema.CooldownExpiration.Sub(time.Now()))
	}
	return nil
}

type equipment struct {
	code     string
	quantity int
}

func (character *Character) GetEquipment() *[]items.Equipment {
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
	equipmentCodes = append(equipmentCodes, equipment{character.Character.Utility1Slot, character.Character.Utility1SlotQuantity})
	equipmentCodes = append(equipmentCodes, equipment{character.Character.Utility2Slot, character.Character.Utility2SlotQuantity})

	equipmentSlots := []items.Equipment{}

	for _, equipment := range equipmentCodes {
		if equipment.code == "" {
			continue
		}
		item, err := items.GetItem(character.logger, equipment.code)
		if err != nil {
			character.logger.Warn("Could not retrieve item", "code", equipment.code, "error", err)
			continue
		}
		equipmentSlots = append(equipmentSlots, items.Equipment{
			Item:     *item,
			Quantity: equipment.quantity,
		})
	}
	return &equipmentSlots
}

func (character *Character) GetLocation() (int, int) {
	return character.Character.X, character.Character.Y
}
