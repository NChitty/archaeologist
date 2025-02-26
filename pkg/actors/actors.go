package actors

import (
	"errors"
	"log/slog"
	"math"
	"time"

	"github.com/NChitty/archaeologist/pkg/models/character"
	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
)

type ActionResult struct {
	CooldownRemaining time.Duration
	Success           bool
}

type Actor interface {
	Do(character *character.Character) error
}

type CharacterAdapter interface {
	UpdateCharacter(character *character.Character) error
	Move(character *character.Character, x int, y int) (*ActionResult, error)
}

type ItemAdapter interface {
	GetItem(name string) (*artifactsmmo.ItemSchema, error)
	GetAllResources(skill *artifactsmmo.GatheringSkill, code *string) ([]artifactsmmo.ResourceSchema, error)
	GetCharacterEquipment(character *character.Character) *[]character.Equipment
	IsGatherable(item *artifactsmmo.ItemSchema) bool
	IsCraftable(item *artifactsmmo.ItemSchema) bool
}

type MapAdapter interface {
	GetAllMaps(contentType *string, contentCode *string, page *int, size *int) ([]artifactsmmo.MapSchema, error)
	GetMap(x int, y int) (*artifactsmmo.MapSchema, error)
}

type MonsterAdapter interface {
	GetAllMonsters(
		minLevel *int,
		maxLevel *int,
		drop *string,
		page *int,
		size *int,
	) (*[]artifactsmmo.MonsterSchema, error)
	GetMonster(code string) (*artifactsmmo.MonsterSchema, error)
}

type Config[C CharacterAdapter] struct {
	characterAdapter C
	fightSimulator   FightSimulator
	itemAdapter      ItemAdapter
	mapAdapter       MapAdapter
	monsterAdapter   MonsterAdapter
	logger           *slog.Logger
}

func NewConfig[C CharacterAdapter](ca C, fs FightSimulator, ia ItemAdapter, ma MapAdapter, moa MonsterAdapter, logger *slog.Logger) *Config[C] {
	return &Config[C]{ca, fs, ia, ma, moa, logger}
}

func move(mapAdapter MapAdapter, characterAdapter CharacterAdapter, character *character.Character, contentType *string, contentCode *string) error {
	maps, err := mapAdapter.GetAllMaps(
		contentType,
		contentCode,
		nil,
		nil,
	)
	if err != nil {
		slog.Error("Could not retrieve all map tiles potentially relevant to actor.", "error", err)
		return err
	}
	if len(maps) == 0 {
		slog.Error("No maps with parameters.", "contentType", *contentType, "contentCode", *contentCode)
		return errors.New("No maps returned with given parameters.")
	}
	minDist := math.MaxInt
	var dest *artifactsmmo.MapSchema
	for _, cell := range maps {
		distX := character.X - cell.X
		distY := character.Y - cell.Y
		if distX < 0 {
			distX *= -1
		}
		if distY < 0 {
			distY *= -1
		}
		if minDist > (distX + distY) {
			minDist = distX + distY
			dest = &cell
		}
	}

	moveRes, err := characterAdapter.Move(character, dest.X, dest.Y)
	if err != nil {
		return err
	}
	time.Sleep(moveRes.CooldownRemaining)

	return nil
}
