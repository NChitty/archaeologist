package fights

import (
	"testing"

	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	type testCase struct {
		character   *artifactsmmo.CharacterSchema
		monster     *artifactsmmo.MonsterSchema
		expectedDmg uint16
	}
	testCases := []testCase{
		testCase{
			character: &artifactsmmo.CharacterSchema{
				AttackFire:  0,
				AttackEarth: 0,
				AttackWater: 0,
				AttackAir:   16,
				DmgFire:     12,
				DmgEarth:    12,
				DmgWater:    7,
				DmgAir:      7,
				ResFire:     4,
				ResEarth:    4,
				ResWater:    4,
				ResAir:      4,
			},
			monster: &artifactsmmo.MonsterSchema{
				Hp:          70,
				AttackFire:  0,
				AttackEarth: 8,
				AttackWater: 0,
				AttackAir:   0,
				ResFire:     0,
				ResEarth:    25,
				ResWater:    0,
				ResAir:      0,
			},
			expectedDmg: 17,
		},
		testCase{
			character: &artifactsmmo.CharacterSchema{
				AttackFire:  0,
				AttackEarth: 0,
				AttackWater: 0,
				AttackAir:   16,
				DmgFire:     12,
				DmgEarth:    12,
				DmgWater:    7,
				DmgAir:      7,
				ResFire:     4,
				ResEarth:    4,
				ResWater:    4,
				ResAir:      4,
			},
			monster: &artifactsmmo.MonsterSchema{
				Hp:          70,
				AttackFire:  0,
				AttackEarth: 0,
				AttackWater: 0,
				AttackAir:   12,
				ResFire:     0,
				ResEarth:    0,
				ResWater:    0,
				ResAir:      25,
			},
			expectedDmg: 13,
		},
	}
	t.Run("Damage calculations", func(t *testing.T) {
		for _, testCase := range testCases {
			assert.Equal(t, testCase.expectedDmg, calculateCharacterDamage(testCase.character, testCase.monster))
		}
	})
}
