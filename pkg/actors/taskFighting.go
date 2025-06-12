package actors

import (
	"errors"
	"log/slog"
	"time"

	"github.com/NChitty/archaeologist/pkg/models/character"
)

type TaskFightingActor struct {
	monster          string
	exitOnRest       bool
	characterService FightingCharacterAdapter
	fightService     FightSimulator
	itemAccessor     ItemAdapter
	mapAccessor      MapAdapter
	monsterAccessor  MonsterAdapter
	logger           *slog.Logger
}

func NewTaskFightingActor[C FightingCharacterAdapter](
	character *character.Character,
	exitOnRest bool,
	config *Config[C],
) (Actor, error) {
	if character.TaskType != "monsters" {
		return nil, errors.New("Unsupported task type: " + character.TaskType)
	}
	config.logger.Info("Created new Task Fighting Actor", "monster", character.Task, "taskTotal", character.TaskTotal, "taskProgress", character.TaskProgress)
	return &TaskFightingActor{
		monster:          character.Task,
		exitOnRest:       exitOnRest,
		characterService: config.characterAdapter,
		fightService:     config.fightSimulator,
		itemAccessor:     config.itemAdapter,
		mapAccessor:      config.mapAdapter,
		monsterAccessor:  config.monsterAdapter,
		logger:           config.logger,
	}, nil
}


func (actor *TaskFightingActor) Do(character *character.Character) error {
	err := actor.characterService.UpdateCharacter(character)
	if err != nil {
		actor.logger.Error("Failed to update character", slog.Any("error", err))
		return err
	}
	monster, err := actor.monsterAccessor.GetMonster(actor.monster)
	if err != nil {
		actor.logger.Error("Could not find monster with code", "error", err)
		return err
	}

	fightResult, err := actor.fightService.CalculateFightResult(character, monster)
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

	err = move(actor.mapAccessor, actor.characterService, character, nil, &actor.monster)
	if err != nil {
		actor.logger.Error("Could not move character", "error", err)
		return err
	}

	for character.TaskProgress < character.TaskTotal {
		fightRes, err := actor.characterService.Fight(character)
		if err != nil {
			return err
		}
		time.Sleep(fightRes.CooldownRemaining)

		if fightResult.RestoreTurns > 0 {
			fightResult, err = actor.fightService.CalculateFightResult(character, monster)
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

func (actor *TaskFightingActor) heal(character *character.Character, fightResult *FightResult) error {
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

	if len(healingItems) == 0 && actor.exitOnRest {
		return ExitOnRest
	}

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
		actor.logger.Warn("No item found to heal to minimum health", "minHealing", minHealing)
		if actor.exitOnRest {
			return ExitOnRest
		}
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
