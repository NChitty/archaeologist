package effects_test

import (
	"testing"

	"github.com/NChitty/archaeologist/pkg/items/effects"
	"github.com/stretchr/testify/assert"
)

func TestItemEffectsMapping(t *testing.T) {
  for k, v := range effects.Effects() {
    assert.Equal(t, v, k.GetEffectName())
    assert.Equal(t, k, effects.GetEffect(v))
  }
}
