package character

import (
	"context"
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

	return &characterResp.JSON200.Data, nil
}
