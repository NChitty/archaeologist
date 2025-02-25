package characters

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/NChitty/archaeologist/pkg/models/character"
	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
)

type CharacterService struct {
	client artifactsmmo.ClientWithResponsesInterface
	logger *slog.Logger
}

func NewCharacterService(
	logger *slog.Logger,
	client artifactsmmo.ClientWithResponsesInterface,
) *CharacterService {
	character := &CharacterService{
		client: client,
		logger: logger,
	}
	return character
}

func (service *CharacterService) UpdateCharacter(character *character.Character) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	characterResp, err := service.client.GetCharacterCharactersNameGetWithResponse(
		ctx,
		character.Name,
		artifactsmmo.NewBearerAuthorizationRequestFunc(character.Token),
	)
	if err != nil {
		service.logger.Error("Could not retrieve character", "error", err)
		return err
	}
	if characterResp.StatusCode() != 200 {
		service.logger.Error("Could not retrieve character", "name", character.Name, "body", string(characterResp.Body))
		return errors.New("Could not retrieve character with name: " + character.Name)
	}

	character.FromSchema(characterResp.JSON200.Data)

	return nil
}

func (service *CharacterService) GetInventory(character *character.Character) (map[string]artifactsmmo.InventorySlot, error) {
	err := service.UpdateCharacter(character)
	if err != nil {
		return nil, err
	}
	return character.Inventory, nil
}

func (service *CharacterService) WaitCooldown(character *character.Character) error {
	err := service.UpdateCharacter(character)
	if err != nil {
		return err
	}
	if time.Now().Before(*character.CooldownExpiration) {
		service.logger.Debug(
			"Waiting for cooldown",
			"expiration", character.CooldownExpiration,
			"timeRemaining", character.CooldownExpiration.Sub(time.Now().UTC()),
		)
		time.Sleep(character.CooldownExpiration.Sub(time.Now()))
	}
	return nil
}
