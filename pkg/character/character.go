package character

import (
	"context"
	"errors"
	"log/slog"
	"time"

	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
)

type Character struct {
	Name      string
	Character *artifactsmmo.CharacterSchema
	client    *artifactsmmo.ClientWithResponses
}

func New(client *artifactsmmo.ClientWithResponses, name string) (*Character, error) {
	character := &Character{
		Name:   name,
		client: client,
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
		slog.Error("Could not retrieve character:", err)
		return nil, err
	}
	if characterResp.StatusCode() == 404 {
		slog.Error("Could not retrieve character", "name", character.Name)
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
		slog.Debug(
			"Waiting for cooldown",
			"expiration", charSchema.CooldownExpiration,
			"timeRemaining", charSchema.CooldownExpiration.Sub(time.Now()),
		)
		time.Sleep(charSchema.CooldownExpiration.Sub(time.Now()))
	}
  return nil
}

func (character *Character) GetLocation() (int, int) {
  return character.Character.X, character.Character.Y
}
