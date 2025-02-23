package monsters_test

import (
	"log/slog"
	"testing"

	"github.com/NChitty/archaeologist/pkg/monsters"
	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
	"github.com/stretchr/testify/assert"
)

var chicken artifactsmmo.MonsterSchema = artifactsmmo.MonsterSchema{
	Name:           "Chicken",
	Code:           "chicken",
	Level:          1,
	Hp:             60,
	AttackFire:     0,
	AttackEarth:    0,
	AttackWater:    4,
	AttackAir:      0,
	ResFire:        0,
	ResEarth:       0,
	ResWater:       0,
	ResAir:         0,
	CriticalStrike: 0,
	Effects:        &[]artifactsmmo.SimpleEffectSchema{},
	MinGold:        0,
	MaxGold:        3,
	Drops: []artifactsmmo.DropRateSchema{
		artifactsmmo.DropRateSchema{
			Code:        "raw_chicken",
			Rate:        10,
			MinQuantity: 1,
			MaxQuantity: 1,
		},
		artifactsmmo.DropRateSchema{
			Code:        "egg",
			Rate:        12,
			MinQuantity: 1,
			MaxQuantity: 1,
		},
		artifactsmmo.DropRateSchema{
			Code:        "feather",
			Rate:        8,
			MinQuantity: 1,
			MaxQuantity: 1,
		},
		artifactsmmo.DropRateSchema{
			Code:        "golden_egg",
			Rate:        1000,
			MinQuantity: 1,
			MaxQuantity: 1,
		},
	},
}

func TestGetAllMonsters(t *testing.T) {
	client, err := artifactsmmo.NewClientWithResponses("https://api.artifactsmmo.com/")
	assert.NoError(t, err)
	monsterAccessor := monsters.NewClientMonsterAccessor(client, slog.Default())
	t.Run("Get All with Drop", func(t *testing.T) {
		rawChicken := "raw_chicken"
		monsters, err := monsterAccessor.GetAllMonsters(nil, nil, &rawChicken, nil, nil)
		assert.NoError(t, err)
		assert.Contains(t, *monsters, chicken)
	})
	t.Run("Get All with Error", func(t *testing.T) {
		maxLevel := -1
		_, err := monsterAccessor.GetAllMonsters(&maxLevel, nil, nil, nil, nil)
		assert.Error(t, err)
	})
}

func TestGetMonster(t *testing.T) {
	client, err := artifactsmmo.NewClientWithResponses("https://api.artifactsmmo.com/")
	assert.NoError(t, err)
	monsterAccessor := monsters.NewClientMonsterAccessor(client, slog.Default())
  t.Run("Get Chicken", func(t *testing.T) {
		monster, err := monsterAccessor.GetMonster("chicken")
		assert.NoError(t, err)
		assert.Equal(t, *monster, chicken)
  })
  t.Run("Get DNE", func(t *testing.T) {
		_, err := monsterAccessor.GetMonster("DNE")
		assert.Error(t, err)
  })
}
