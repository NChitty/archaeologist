package effects_test

import (
	"log/slog"
	"testing"

	"github.com/NChitty/archaeologist/pkg/items/effects"
	"github.com/NChitty/archaeologist/pkg/models/character"
	"github.com/NChitty/archaeologist/pkg/models/item"
	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
	"github.com/stretchr/testify/assert"
)

func TestEffectAccumulator(t *testing.T) {
	type testCase struct {
		items                 *[]character.Equipment
		expectedResult        map[item.Effect]int
		expectedRestoreResult []effects.RestoreEffect
	}
	testCases := []testCase{
		{
			items: &[]character.Equipment{
				{
					Item: artifactsmmo.ItemSchema{
						Effects: &[]artifactsmmo.SimpleEffectSchema{
							{
								Code:  item.HpEffect.GetEffectCode(),
								Value: 25,
							},
						},
					},
					Quantity: 1,
				},
				{
					Item: artifactsmmo.ItemSchema{
						Effects: &[]artifactsmmo.SimpleEffectSchema{
							{
								Code:  item.HpEffect.GetEffectCode(),
								Value: 25,
							},
						},
					},
					Quantity: 1,
				},
			},
			expectedResult:        map[item.Effect]int{item.HpEffect: 50},
			expectedRestoreResult: []effects.RestoreEffect{},
		},
		{
			items: &[]character.Equipment{
				{
					Item: artifactsmmo.ItemSchema{
						Effects: &[]artifactsmmo.SimpleEffectSchema{
							{
								Code:  item.RestoreEffect.GetEffectCode(),
								Value: 20,
							},
						},
					},
					Quantity: 5,
				},
			},
			expectedResult: map[item.Effect]int{},
			expectedRestoreResult: []effects.RestoreEffect{
				{
					Quantity: 5,
					Value:    20,
				},
			},
		},
	}
	t.Run("Accumulate", func(t *testing.T) {
		for _, testCase := range testCases {
			accumulator := effects.New(slog.Default())
			accumulator.Accumulate(testCase.items)
			assert.Equal(t, testCase.expectedResult, accumulator.GetEffects())
			assert.Equal(t, testCase.expectedRestoreResult, accumulator.GetRestoreEffects())
		}
	})
}
