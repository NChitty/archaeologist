package item_test

import (
	"testing"

	"github.com/NChitty/archaeologist/pkg/models/item"
	"github.com/stretchr/testify/assert"
)

func TestItemEffectsMapping(t *testing.T) {
  for k, v := range item.Effects() {
    assert.Equal(t, v, k.GetEffectCode())
    assert.Equal(t, k, item.GetEffect(v))
  }
}
