package fights_test

import (
	"log/slog"
	"testing"

	"github.com/NChitty/archaeologist/pkg/characters"
	"github.com/NChitty/archaeologist/pkg/fights"
	"github.com/NChitty/archaeologist/pkg/items"
	"github.com/NChitty/archaeologist/pkg/items/effects"
	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	type testCase struct {
		character      *characters.CharacterWrapper
		monster        *artifactsmmo.MonsterSchema
		expectedResult fights.FightResult
	}
	testCases := []testCase{
		testCase{
			character: characters.FromSchema(artifactsmmo.CharacterSchema{
				MaxHp:      120,
				WeaponSlot: "copper_dagger",
			}),
			monster: &artifactsmmo.MonsterSchema{
				Hp:          60,
				AttackWater: 4,
			},
			expectedResult: fights.FightResult{
				true,
				19,
				10,
				36,
				0,
				6,
				4,
			},
		},
		testCase{
			character: characters.FromSchema(artifactsmmo.CharacterSchema{
				MaxHp:         220,
				WeaponSlot:    "sticky_sword",
				ShieldSlot:    "wooden_shield",
				HelmetSlot:    "copper_helmet",
				BodyArmorSlot: "copper_armor",
				LegArmorSlot:  "copper_legs_armor",
				BootsSlot:     "copper_boots",
				Ring1Slot:     "copper_ring",
				Ring2Slot:     "copper_ring",
				AmuletSlot:    "",
				Artifact1Slot: "",
				Artifact2Slot: "",
				Artifact3Slot: "",

				Utility1Slot:         "small_health_potion",
				Utility1SlotQuantity: 10,
				Utility2Slot:         "",
				Utility2SlotQuantity: 0,
			}),
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
			expectedResult: fights.FightResult{
				true,
				9,
				5,
				32,
				0,
				14,
				8,
			},
		},
		testCase{
			character: characters.FromSchema(artifactsmmo.CharacterSchema{
				MaxHp:         220,
				WeaponSlot:    "sticky_sword",
				ShieldSlot:    "wooden_shield",
				HelmetSlot:    "copper_helmet",
				BodyArmorSlot: "copper_armor",
				LegArmorSlot:  "copper_legs_armor",
				BootsSlot:     "copper_boots",
				Ring1Slot:     "copper_ring",
				Ring2Slot:     "copper_ring",
				AmuletSlot:    "",
				Artifact1Slot: "",
				Artifact2Slot: "",
				Artifact3Slot: "",

				Utility1Slot:         "small_health_potion",
				Utility1SlotQuantity: 10,
				Utility2Slot:         "",
				Utility2SlotQuantity: 0,
			}),
			monster: &artifactsmmo.MonsterSchema{
				Hp:          80,
				AttackFire:  0,
				AttackEarth: 0,
				AttackWater: 0,
				AttackAir:   12,
				ResFire:     0,
				ResEarth:    0,
				ResWater:    0,
				ResAir:      25,
			},
			expectedResult: fights.FightResult{
				true,
				9,
				5,
				48,
				0,
				18,
				12,
			},
		},
		testCase{
			character: characters.FromSchema(artifactsmmo.CharacterSchema{
				MaxHp:         220,
				WeaponSlot:    "sticky_sword",
				ShieldSlot:    "wooden_shield",
				HelmetSlot:    "copper_helmet",
				BodyArmorSlot: "copper_armor",
				LegArmorSlot:  "copper_legs_armor",
				BootsSlot:     "copper_boots",
				Ring1Slot:     "copper_ring",
				Ring2Slot:     "copper_ring",
				AmuletSlot:    "",
				Artifact1Slot: "",
				Artifact2Slot: "",
				Artifact3Slot: "",

				Utility1Slot:         "small_health_potion",
				Utility1SlotQuantity: 10,
				Utility2Slot:         "",
				Utility2SlotQuantity: 0,
			}),
			monster: &artifactsmmo.MonsterSchema{
				Hp:          120,
				AttackFire:  18,
				AttackEarth: 0,
				AttackWater: 0,
				AttackAir:   0,
				ResFire:     25,
				ResEarth:    0,
				ResWater:    0,
				ResAir:      0,
			},
			expectedResult: fights.FightResult{
				true,
				13,
				7,
				108,
				0,
				18,
				18,
			},
		},
		testCase{
			character: characters.FromSchema(artifactsmmo.CharacterSchema{
				MaxHp:         220,
				WeaponSlot:    "sticky_sword",
				ShieldSlot:    "wooden_shield",
				HelmetSlot:    "copper_helmet",
				BodyArmorSlot: "copper_armor",
				LegArmorSlot:  "copper_legs_armor",
				BootsSlot:     "copper_boots",
				Ring1Slot:     "copper_ring",
				Ring2Slot:     "copper_ring",
				AmuletSlot:    "",
				Artifact1Slot: "",
				Artifact2Slot: "",
				Artifact3Slot: "",

				Utility1Slot:         "small_health_potion",
				Utility1SlotQuantity: 10,
				Utility2Slot:         "",
				Utility2SlotQuantity: 0,
			}),
			monster: &artifactsmmo.MonsterSchema{
				Hp:          280,
				AttackFire:  0,
				AttackEarth: 0,
				AttackWater: 0,
				AttackAir:   21,
				ResFire:     0,
				ResEarth:    -30,
				ResWater:    30,
				ResAir:      0,
			},
			expectedResult: fights.FightResult{
				true,
				25,
				13,
				112,
				7,
				23,
				21,
			},
		},
	}
	t.Run("Damage calculations", func(t *testing.T) {
		itemService := items.DefaultItemService()
		effectsAccumulator := effects.New(slog.Default())
		service := fights.NewFightService(effectsAccumulator, slog.Default(), itemService)
		for _, testCase := range testCases {
			fightResult, err := service.CalculateFightResult(testCase.character, testCase.monster)
			assert.NoError(t, err)
			assert.EqualValues(t, testCase.expectedResult, *fightResult)
		}
	})
}
