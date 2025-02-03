package character

import (
	"context"
	"errors"
	"log/slog"
	"time"

	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
)

type Character struct {
	Name   string
	client *artifactsmmo.ClientWithResponses
}

func New(client *artifactsmmo.ClientWithResponses, name string) (*Character, error) {
	character := &Character{
		Name:   name,
		client: client,
	}
	_, err := character.GetCharacter()
	return character, err
}

func (character *Character) GetCharacter() (*artifactsmmo.CharacterSchema, error) {
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

	return &characterResp.JSON200.Data, nil
}

func (character *Character) GetInventory() (map[string]artifactsmmo.InventorySlot, error) {
	characterInfo, err := character.GetCharacter()
	if err != nil {
		return nil, err
	}
  inventoryMap := make(map[string]artifactsmmo.InventorySlot)
  for _, slot := range(*characterInfo.Inventory) {
    inventoryMap[slot.Code] = slot
  }
  return inventoryMap, nil
}
