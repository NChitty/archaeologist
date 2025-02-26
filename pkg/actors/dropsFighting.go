package actors

import (
	"errors"
	"fmt"
	"log/slog"
	"math"
	"sync"
	"time"

	"github.com/NChitty/archaeologist/pkg/models/character"
	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
)

type DropsFightingActor struct {
	code             string
	qty              int
	characterService FightingCharacterAdapter
	fightService     FightSimulator
	itemAccessor     ItemAdapter
	mapAccessor      MapAdapter
	monsterAccessor  MonsterAdapter
	logger           *slog.Logger
}

func NewDropsFightingActor[C FightingCharacterAdapter](
	character *character.Character,
	code string,
	qty int,
	config *Config[C],
) (Actor, error) {
	monsters, err := config.monsterAdapter.GetAllMonsters(nil, maxLevel(*character), &code, nil, nil)
	if err != nil {
		return nil, err
	}
	if len(*monsters) == 0 {
		config.logger.Error("Could not find any monsters within a reasonable level range with drop.")
		return nil, errors.New("No monsters found")
	}
	return &DropsFightingActor{
		code,
		qty,
		config.characterAdapter,
		config.fightSimulator,
		config.itemAdapter,
		config.mapAdapter,
		config.monsterAdapter,
		config.logger,
	}, nil
}

func (actor *DropsFightingActor) Do(character *character.Character) error {
	err := actor.characterService.UpdateCharacter(character)
	if err != nil {
		actor.logger.Error("Failed to update character", slog.Any("error", err))
		return err
	}
	monsters, err := actor.monsterAccessor.GetAllMonsters(nil, maxLevel(*character), &actor.code, nil, nil)
	if err != nil {
		return err
	}
	if len(*monsters) == 0 {
		actor.logger.Error("Could not find any monsters within a reasonable level range with drop.")
		return errors.New("No monsters found")
	}
	monster, err := actor.optimizeMonsterSelection(*monsters, character)
	if err != nil {
		return err
	}
	if monster.Code == "" {
		return errors.New("No monster found")
	}

	fightResult, err := actor.fightService.CalculateFightResult(character, &monster)
	if err != nil {
		actor.logger.Error("Could not simulate fight", "error", err)
		return err
	}

	if !fightResult.Win {
		return errors.New("Fight is not winnable in the worst-case")
	}

	err = actor.heal(character, fightResult)
	if err != nil {
		return err
	}

	err = move(actor.mapAccessor, actor.characterService, character, nil, &(monster.Code))
	if err != nil {
		actor.logger.Error("Could not move character", "error", err)
		return err
	}

	for character.Inventory[actor.code].Quantity < actor.qty {
		fightRes, err := actor.characterService.Fight(character)
		if err != nil {
			return err
		}
		time.Sleep(fightRes.CooldownRemaining)

		if fightResult.RestoreTurns > 0 {
			fightResult, err = actor.fightService.CalculateFightResult(character, &monster)
			if err != nil {
				actor.logger.Error("Could not simulate fight", "error", err)
				return err
			}
			if !fightResult.Win {
				return errors.New("Fight is not winnable in the worst-case")
			}
		}
		err = actor.heal(character, fightResult)
		if err != nil {
			if errors.Is(err, ExitOnRest) {
				return nil
			}
			return err
		}
	}

	return nil
}

func maxLevel(character character.Character) *int {
	maxLevel := character.Level + 2
	return &maxLevel
}

func (actor *DropsFightingActor) optimizeMonsterSelection(
	monsters []artifactsmmo.MonsterSchema,
	character *character.Character,
) (artifactsmmo.MonsterSchema, error) {
	type optimization struct {
		monster       artifactsmmo.MonsterSchema
		hpLoss        int
		expectedValue float64
		restoreTurns  int
	}
	min := optimization{
		hpLoss:        math.MaxInt,
		expectedValue: 0,
		restoreTurns:  math.MaxInt,
	}
	var wg sync.WaitGroup
	var lock sync.Mutex
	for _, monster := range monsters {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fightResult, err := actor.fightService.CalculateFightResult(character, &monster)
			if err != nil {
				actor.logger.Error("Could not get fight result", "monster", monster.Code)
				return
			}
			if !fightResult.Win {
				return
			}

			drops := mapDrops(monster.Drops)
			expectedValue := (float64(drops[actor.code].MinQuantity) + float64(drops[actor.code].MaxQuantity)/float64(2)) * 1 / float64(drops[actor.code].Rate)
			actor.logger.Debug("Fight result", "monster", monster.Code, "expectedValue", expectedValue, "fightResult", fmt.Sprintf("%+v", fightResult))

			lock.Lock()
			if min.expectedValue < expectedValue && min.hpLoss >= fightResult.CharacterHpLoss && min.restoreTurns > fightResult.RestoreTurns {
				min.monster = monster
				min.expectedValue = expectedValue
				min.hpLoss = fightResult.CharacterHpLoss
				min.restoreTurns = fightResult.RestoreTurns
			}
			lock.Unlock()
		}()
	}

	wg.Wait()

	if min.monster.Code == "" {
		return min.monster, errors.New("No monsters found")
	}

	return min.monster, nil
}

func mapDrops(drops []artifactsmmo.DropRateSchema) map[string]artifactsmmo.DropRateSchema {
	retVal := map[string]artifactsmmo.DropRateSchema{}
	for _, drop := range drops {
		retVal[drop.Code] = drop
	}
	return retVal
}

func (actor *DropsFightingActor) heal(character *character.Character, fightResult *FightResult) error {
	if fightResult.CharacterHpLoss < character.Hp {
		return nil
	}

	actor.logger.Warn(
		"Expected HP loss of fight is greater than current HP - healing...",
		"hp",
		character.Hp,
		"hpLoss",
		fightResult.CharacterHpLoss,
	)

	healingItems := getHealingItems(actor.itemAccessor, character, actor.logger)

	if len(healingItems) == 0 {
		actor.logger.Info("No healing items, resting...")
		healRes, err := actor.characterService.Rest(character)
		if err != nil {
			return err
		}
		time.Sleep(healRes.CooldownRemaining)
		return nil
	}

	targetHealing := character.MaxHp - character.Hp
	minHealing := 1 + fightResult.CharacterHpLoss - character.Hp
	useSchema := buildUseSchema(character, healingItems, minHealing, targetHealing)
	if useSchema.Code == "" || useSchema.Quantity == 0 {
		actor.logger.Warn("No item found to heal to minimum health")
		actor.logger.Info("Resting...")
		healRes, err := actor.characterService.Rest(character)
		if err != nil {
			return err
		}
		time.Sleep(healRes.CooldownRemaining)
		return nil
	}

	healRes, err := actor.characterService.Use(character, useSchema)
	if err != nil {
		actor.logger.Error("Could not use item", "use", *useSchema, "error", err)
		return err
	}
	time.Sleep(healRes.CooldownRemaining)

	return nil
}
