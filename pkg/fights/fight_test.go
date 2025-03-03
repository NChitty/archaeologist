package fights_test

import (
	"log/slog"
	"testing"

	"github.com/NChitty/archaeologist/pkg/actors"
	"github.com/NChitty/archaeologist/pkg/fights"
	"github.com/NChitty/archaeologist/pkg/items"
	"github.com/NChitty/archaeologist/pkg/items/effects"
	"github.com/NChitty/archaeologist/pkg/models/character"
	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	character      *character.Character
	monster        *artifactsmmo.MonsterSchema
	expectedResult actors.FightResult
}

var simpleChickenFight testCase = testCase{
	character: character.FromSchema(artifactsmmo.CharacterSchema{
		MaxHp:      120,
		WeaponSlot: "copper_dagger",
	}),
	monster: &artifactsmmo.MonsterSchema{
		Hp:          60,
		AttackWater: 4,
		Effects:     &[]artifactsmmo.SimpleEffectSchema{},
	},
	expectedResult: actors.FightResult{
		Win:             true,
		Turns:           19,
		CharacterTurns:  10,
		CharacterHpLoss: 36,
		RestoreTurns:    0,
		CharacterDmg:    6,
		MonsterDmg:      4,
	},
}

var chickenFightWithReconstitution testCase = testCase{
	character: character.FromSchema(artifactsmmo.CharacterSchema{
		MaxHp:      120,
		WeaponSlot: "copper_dagger",
	}),
	monster: &artifactsmmo.MonsterSchema{
		Hp:          38,
		AttackWater: 4,
		Effects: &[]artifactsmmo.SimpleEffectSchema{
			{
				Code:  "reconstitution",
				Value: 3,
			},
		},
	},
	expectedResult: actors.FightResult{
		Win:             true,
		Turns:           19,
		CharacterTurns:  10,
		CharacterHpLoss: 32,
		RestoreTurns:    0,
		CharacterDmg:    6,
		MonsterDmg:      4,
	},
}

var yellowSlimeballResistance testCase = testCase{
	character: character.FromSchema(artifactsmmo.CharacterSchema{
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

		Utility1Slot:         "",
		Utility1SlotQuantity: 0,
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
		Effects:     &[]artifactsmmo.SimpleEffectSchema{},
	},
	expectedResult: actors.FightResult{
		Win:             true,
		Turns:           9,
		CharacterTurns:  5,
		CharacterHpLoss: 32,
		RestoreTurns:    0,
		CharacterDmg:    14,
		MonsterDmg:      8,
	},
}

var cowFightNegativeResistanceAndRestore testCase = testCase{
	character: character.FromSchema(artifactsmmo.CharacterSchema{
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
		Effects:     &[]artifactsmmo.SimpleEffectSchema{},
	},
	expectedResult: actors.FightResult{
		Win:             true,
		Turns:           27,
		CharacterTurns:  14,
		CharacterHpLoss: 112,
		RestoreTurns:    7,
		CharacterDmg:    23,
		MonsterDmg:      21,
	},
}

var poison testCase = testCase{
	character: character.FromSchema(artifactsmmo.CharacterSchema{
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

		Utility1Slot:         "small_antidote",
		Utility1SlotQuantity: 10,
		Utility2Slot:         "",
		Utility2SlotQuantity: 0,
	}),
	monster: &artifactsmmo.MonsterSchema{
		Hp:          100,
		AttackFire:  0,
		AttackEarth: 0,
		AttackWater: 0,
		AttackAir:   0,
		ResFire:     0,
		ResEarth:    -30,
		ResWater:    30,
		ResAir:      0,
		Effects: &[]artifactsmmo.SimpleEffectSchema{
			{
				Code:  "poison",
				Value: 50,
			},
		},
	},
	expectedResult: actors.FightResult{
		Win:             true,
		Turns:           11,
		CharacterTurns:  6,
		CharacterHpLoss: 40,
		RestoreTurns:    0,
		CharacterDmg:    23,
		MonsterDmg:      0,
	},
}

var burn testCase = testCase{
	character: character.FromSchema(artifactsmmo.CharacterSchema{
		MaxHp:      220,
		WeaponSlot: "copper_dagger",
	}),
	monster: &artifactsmmo.MonsterSchema{
		Hp:          100,
		AttackFire:  5,
		AttackEarth: 5,
		AttackWater: 5,
		AttackAir:   5,
		ResFire:     0,
		ResEarth:    0,
		ResWater:    0,
		ResAir:      0,
		Effects: &[]artifactsmmo.SimpleEffectSchema{
			{
				Code:  "burn",
				Value: 10,
			},
		},
	},
	expectedResult: actors.FightResult{
		Win:             false,
		Turns:           24,
		CharacterTurns:  12,
		CharacterHpLoss: 231,
		RestoreTurns:    0,
		CharacterDmg:    6,
		MonsterDmg:      20,
	},
}

var healing testCase = testCase{
	character: character.FromSchema(artifactsmmo.CharacterSchema{
		MaxHp:      220,
		WeaponSlot: "copper_dagger",
		RuneSlot:   "healing_rune",
	}),
	monster: &artifactsmmo.MonsterSchema{
		Hp:          100,
		AttackFire:  5,
		AttackEarth: 0,
		AttackWater: 0,
		AttackAir:   0,
		ResFire:     0,
		ResEarth:    0,
		ResWater:    0,
		ResAir:      0,
		Effects:     &[]artifactsmmo.SimpleEffectSchema{},
	},
	expectedResult: actors.FightResult{
		Win:             true,
		Turns:           35,
		CharacterTurns:  18,
		CharacterHpLoss: 20,
		RestoreTurns:    0,
		CharacterDmg:    6,
		MonsterDmg:      5,
	},
}

var characterBurn testCase = testCase{
	character: character.FromSchema(artifactsmmo.CharacterSchema{
		MaxHp:      220,
		WeaponSlot: "wrathsword",
		RuneSlot:   "burn_rune",
	}),
	monster: &artifactsmmo.MonsterSchema{
		Hp:          1000,
		AttackFire:  0,
		AttackEarth: 0,
		AttackWater: 0,
		AttackAir:   0,
		ResFire:     0,
		ResEarth:    0,
		ResWater:    0,
		ResAir:      0,
		Effects:     &[]artifactsmmo.SimpleEffectSchema{},
	},
	expectedResult: actors.FightResult{
		Win:             true,
		Turns:           19,
		CharacterTurns:  10,
		CharacterHpLoss: 0,
		RestoreTurns:    0,
		CharacterDmg:    100,
		MonsterDmg:      0,
	},
}

func Test(t *testing.T) {
	itemService := items.DefaultItemService()
	effectsAccumulator := effects.New(slog.Default())
	service := fights.NewFightService(effectsAccumulator, slog.Default(), itemService)
	t.Run("Simple Chicken Fight", func(t *testing.T) {
		fightResult, err := service.CalculateFightResult(simpleChickenFight.character, simpleChickenFight.monster)
		assert.NoError(t, err)
		assert.EqualValues(t, simpleChickenFight.expectedResult, *fightResult)
	})
	t.Run("Chicken Fight with Reconstitution", func(t *testing.T) {
		fightResult, err := service.CalculateFightResult(chickenFightWithReconstitution.character, chickenFightWithReconstitution.monster)
		assert.NoError(t, err)
		assert.EqualValues(t, chickenFightWithReconstitution.expectedResult, *fightResult)
	})
	t.Run("Yellow slimeball resistance calculation", func(t *testing.T) {
		fightResult, err := service.CalculateFightResult(yellowSlimeballResistance.character, yellowSlimeballResistance.monster)
		assert.NoError(t, err)
		assert.EqualValues(t, yellowSlimeballResistance.expectedResult, *fightResult)
	})
	t.Run("Cow negative resistance and restore", func(t *testing.T) {
		fightResult, err := service.CalculateFightResult(cowFightNegativeResistanceAndRestore.character, cowFightNegativeResistanceAndRestore.monster)
		assert.NoError(t, err)
		assert.EqualValues(t, cowFightNegativeResistanceAndRestore.expectedResult, *fightResult)
	})
	t.Run("Poison with antidote", func(t *testing.T) {
		fightResult, err := service.CalculateFightResult(poison.character, poison.monster)
		assert.NoError(t, err)
		assert.EqualValues(t, poison.expectedResult, *fightResult)
	})
	t.Run("Burn", func(t *testing.T) {
		fightResult, err := service.CalculateFightResult(burn.character, burn.monster)
		assert.NoError(t, err)
		assert.EqualValues(t, burn.expectedResult, *fightResult)
	})
	t.Run("Healing", func(t *testing.T) {
		fightResult, err := service.CalculateFightResult(healing.character, healing.monster)
		assert.NoError(t, err)
		assert.EqualValues(t, healing.expectedResult, *fightResult)
	})
	t.Run("Character Burn", func(t *testing.T) {
		fightResult, err := service.CalculateFightResult(characterBurn.character, characterBurn.monster)
		assert.NoError(t, err)
		assert.EqualValues(t, characterBurn.expectedResult, *fightResult)
	})
}
