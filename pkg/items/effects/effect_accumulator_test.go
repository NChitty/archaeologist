package effects_test

import (
	"log/slog"
	"testing"

	"github.com/NChitty/archaeologist/pkg/items"
	"github.com/NChitty/archaeologist/pkg/items/effects"
	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
	"github.com/stretchr/testify/assert"
)

func TestEffectAccumulator(t *testing.T) {
	type testCase struct {
		items                 *[]items.Equipment
		expectedResult        map[effects.ItemEffect]int
		expectedRestoreResult []effects.RestoreEffect
	}
	testCases := []testCase{
		testCase{
			items: &[]items.Equipment{
				items.Equipment{
					Item: artifactsmmo.ItemSchema{
						Effects: &[]artifactsmmo.ItemEffectSchema{
							artifactsmmo.ItemEffectSchema{
								Name:  effects.Hp.GetEffectName(),
								Value: 25,
							},
						},
					},
					Quantity: 1,
				},
				items.Equipment{
					Item: artifactsmmo.ItemSchema{
						Effects: &[]artifactsmmo.ItemEffectSchema{
							artifactsmmo.ItemEffectSchema{
								Name:  effects.Hp.GetEffectName(),
								Value: 25,
							},
						},
					},
					Quantity: 1,
				},
			},
			expectedResult:        map[effects.ItemEffect]int{effects.Hp: 50},
			expectedRestoreResult: []effects.RestoreEffect{},
		},
		testCase{
			items: &[]items.Equipment{
				items.Equipment{
					Item: artifactsmmo.ItemSchema{
						Effects: &[]artifactsmmo.ItemEffectSchema{
							artifactsmmo.ItemEffectSchema{
								Name:  effects.Restore.GetEffectName(),
								Value: 20,
							},
						},
					},
					Quantity: 5,
				},
			},
			expectedResult: map[effects.ItemEffect]int{},
			expectedRestoreResult: []effects.RestoreEffect{
				effects.RestoreEffect{
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
