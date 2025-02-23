package characters_test

import (
	"log/slog"
	"net/http"
	"testing"

	"github.com/NChitty/archaeologist/pkg/characters"
	"github.com/promiseofcake/artifactsmmo-go-client/client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var desireauxCharacter client.CharacterSchema = client.CharacterSchema{
	Name: "Desireaux",
	Hp:   120,
}

func TestUpdateCharacter(t *testing.T) {
	mockClient := new(mockClient)
	character := &characters.CharacterWrapper{
		Name:  "Desireaux",
		Token: "myToken",
	}
	t.Run("Update character", func(t *testing.T) {
		mockCall := mockClient.On(
			"GetCharacterCharactersNameGetWithResponse",
			mock.Anything,
			"Desireaux",
			mock.Anything,
		).Return(
			&client.GetCharacterCharactersNameGetResponse{
				HTTPResponse: &http.Response{Status: "200 Ok", StatusCode: 200},
				JSON200:      &client.CharacterResponseSchema{Data: desireauxCharacter}},
			nil,
		)
		service := characters.NewCharacterService(slog.Default(), mockClient)
		service.UpdateCharacter(character)
		assert.Equal(t, 120, character.Hp)
		assert.Equal(t, "myToken", character.Token)
		mockCall.Unset()
	})
	t.Run("Update character error", func(t *testing.T) {
		mockCall := mockClient.On(
			"GetCharacterCharactersNameGetWithResponse",
			mock.Anything,
			"Desireaux",
			mock.Anything,
		).Return(
			&client.GetCharacterCharactersNameGetResponse{
				Body:         []byte(`{"error":{"code": 404,"message": "Character not found."}}`),
				HTTPResponse: &http.Response{Status: "404 Not Found", StatusCode: 404},
				JSON200:      nil,
			}, nil)

		service := characters.NewCharacterService(slog.Default(), mockClient)
		err := service.UpdateCharacter(character)
		assert.Error(t, err)
		mockCall.Unset()
	})
}
