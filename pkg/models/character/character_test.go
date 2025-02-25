package character_test

import (
	"testing"

	"github.com/NChitty/archaeologist/pkg/models/character"
	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
	"github.com/stretchr/testify/assert"
)

func TestFromSchema(t *testing.T) {
  expected := character.Character{
    Hp: 60,

    Inventory: map[string]artifactsmmo.InventorySlot{},

    Token: "myToken",
  }
  initial := &character.Character{Token: "myToken"}
  assert.NotEqual(t, expected, initial)
  initial.FromSchema(artifactsmmo.CharacterSchema{
    Hp: 60,
  })
  assert.Equal(t, expected, *initial)
  expected = character.Character{
    Hp: 60,

    Inventory: map[string]artifactsmmo.InventorySlot{},
  }
  assert.Equal(t, expected, *character.FromSchema(artifactsmmo.CharacterSchema{
    Hp: 60,
  }))
}
