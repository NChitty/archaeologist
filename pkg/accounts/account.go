package accounts

import (
	"context"
	"log/slog"
	"time"

	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
)

type Account struct {
	client *artifactsmmo.ClientWithResponses
}

func New(client *artifactsmmo.ClientWithResponses) (*Account, error) {
	account := &Account{client: client}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	_, err := account.client.GetMyCharactersMyCharactersGetWithResponse(ctx)
	if err != nil {
		slog.Error("Could not retrieve account details", "error", err)
		return nil, err
	}

	return account, nil
}

func (account *Account) GetCharacters(token string) ([]artifactsmmo.CharacterSchema, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	charactersResponse, err := account.client.GetMyCharactersMyCharactersGetWithResponse(ctx, artifactsmmo.NewBearerAuthorizationRequestFunc(token))
	if err != nil {
		return nil, err
	}

	return charactersResponse.JSON200.Data, nil
}
