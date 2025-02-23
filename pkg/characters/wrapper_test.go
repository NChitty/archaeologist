package characters_test

import (
	"testing"

	"github.com/NChitty/archaeologist/pkg/characters"
	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
	"github.com/stretchr/testify/assert"
)

func TestFromSchema(t *testing.T) {
  expected := characters.CharacterWrapper{
    Hp: 60,

    Inventory: map[string]artifactsmmo.InventorySlot{},

    Token: "myToken",
  }
  initial := &characters.CharacterWrapper{Token: "myToken"}
  assert.NotEqual(t, expected, initial)
  initial.FromSchema(artifactsmmo.CharacterSchema{
    Hp: 60,
  })
  assert.Equal(t, expected, *initial)
  expected = characters.CharacterWrapper{
    Hp: 60,

    Inventory: map[string]artifactsmmo.InventorySlot{},
  }
  assert.Equal(t, expected, *characters.FromSchema(artifactsmmo.CharacterSchema{
    Hp: 60,
  }))
}
