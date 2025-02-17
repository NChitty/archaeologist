package effects_test

import (
	"log/slog"
	"testing"

	"github.com/NChitty/archaeologist/pkg/items/effects"
	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
	"github.com/stretchr/testify/assert"
)

func TestEffectAccumulator(t *testing.T) {
	type testCase struct {
		items          *[]artifactsmmo.ItemEffectSchema
		expectedResult map[effects.ItemEffect]int
	}
	testCases := []testCase{
		testCase{
			items: &[]artifactsmmo.ItemEffectSchema{
				artifactsmmo.ItemEffectSchema{
					Name:  effects.Hp.GetEffectName(),
					Value: 25,
				},
				artifactsmmo.ItemEffectSchema{
					Name:  effects.Hp.GetEffectName(),
					Value: 25,
				},
			},
			expectedResult: map[effects.ItemEffect]int{effects.Hp: 50},
		},
	}
	t.Run("Accumulate", func(t *testing.T) {
		for _, testCase := range testCases {
			accumulator := effects.New(slog.Default())
			accumulator.Accumulate(testCase.items, 0)
			assert.Equal(t, testCase.expectedResult, accumulator.GetEffects())
		}
	})
}
