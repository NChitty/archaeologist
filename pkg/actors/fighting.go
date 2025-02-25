package actors

import (
	"errors"
	"log/slog"
	"math"
	"sync"
	"time"

	"github.com/NChitty/archaeologist/pkg/models/character"
	"github.com/NChitty/archaeologist/pkg/models/item"
	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
)

type FightingCharacterAdapter interface {
	CharacterAdapter
	Fight(character *character.Character) (*ActionResult, error)
	Rest(character *character.Character) (*ActionResult, error)
	Use(character *character.Character, item *artifactsmmo.SimpleItemSchema) (*ActionResult, error)
}

type FightResult struct {
	Win             bool
	Turns           int
	CharacterTurns  int
	CharacterHpLoss int
	RestoreTurns    int
	CharacterDmg    int
	MonsterDmg      int
}

type FightSimulator interface {
	CalculateFightResult(character *character.Character, monster *artifactsmmo.MonsterSchema) (*FightResult, error)
}

type TaskFightingActor struct {
	monster          string
	quantity         int
	exitOnRest       bool
	characterService FightingCharacterAdapter
	fightService     FightSimulator
	itemAccessor     ItemAdapter
	mapAccessor      MapAdapter
	monsterAccessor  MonsterAdapter
	logger           *slog.Logger
}

var ExitOnRest error = errors.New("Exit on rest")

func NewTaskFightingActor(
	character *character.Character,
	exitOnRest bool,
	characterService FightingCharacterAdapter,
	fightService FightSimulator,
	itemAccessor ItemAdapter,
	mapAccessor MapAdapter,
	monsterAccessor MonsterAdapter,
	logger *slog.Logger,
) (Actor, error) {
	if character.TaskType != "monsters" {
		return nil, errors.New("Unsupported task type: " + character.TaskType)
	}
	logger.Info("Created new Task Fighting Actor", "monster", character.Task, "taskTotal", character.TaskTotal, "taskProgress", character.TaskProgress)
	return &TaskFightingActor{
		monster:          character.Task,
		quantity:         character.TaskTotal,
		exitOnRest:       exitOnRest,
		characterService: characterService,
		fightService:     fightService,
		itemAccessor:     itemAccessor,
		mapAccessor:      mapAccessor,
		monsterAccessor:  monsterAccessor,
		logger:           logger,
	}, nil
}

func (actor *TaskFightingActor) Do(character *character.Character) error {
	actor.characterService.UpdateCharacter(character)
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

func buildUseSchema(character *character.Character, healingItems map[string]*artifactsmmo.ItemSchema, targetHealing int) *artifactsmmo.SimpleItemSchema {
	type optimizer struct {
		delta int
		code  string
		qty   int
	}
	min := optimizer{
		delta: math.MaxInt,
		code:  "",
		qty:   0,
	}
	for code, healItem := range healingItems {
		effectMap := mapEffects(healItem.Effects)
		healing, present := effectMap[item.HealEffect]
		if present && float64(targetHealing)/float64(healing) > 1 {
			hasQty := character.Inventory[healItem.Code].Quantity
			neededQty := targetHealing / healing
			qty := neededQty
			delta := (targetHealing % healing) * -1
			if hasQty < neededQty {
				delta = healing*hasQty - targetHealing
				qty = hasQty
			}
			if math.Abs(float64(min.delta)) > math.Abs(float64(delta)) {
				min.delta = delta
				min.qty = qty
				min.code = code
			}
		}
	}
	return &artifactsmmo.SimpleItemSchema{
		Code:     min.code,
		Quantity: min.qty,
	}
}

func mapEffects(effectsSchema *[]artifactsmmo.SimpleEffectSchema) map[item.Effect]int {
	values := map[item.Effect]int{}
	for _, effect := range *effectsSchema {
		values[item.GetEffect(effect.Code)] = effect.Value
	}
	return values
}

func getHealingItems(itemAccessor ItemAdapter, character *character.Character, logger *slog.Logger) map[string]*artifactsmmo.ItemSchema {
	var wg sync.WaitGroup
	var lock sync.Mutex
	consumables := make(map[string]*artifactsmmo.ItemSchema)
	for _, slot := range character.Inventory {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if len(slot.Code) == 0 {
				return
			}
			slotItem, err := itemAccessor.GetItem(slot.Code)
			if err != nil {
				logger.Warn("Could not get item", "code", slot.Code, "error", err)
			}
			canHeal := false
			for _, effect := range *slotItem.Effects {
				if effect.Code == item.HealEffect.GetEffectName() {
					canHeal = true
					break
				}
			}
			if canHeal {
				lock.Lock()
				consumables[slotItem.Code] = slotItem
				lock.Unlock()
			}
		}()
	}
	wg.Wait()
	return consumables
}

func (actor *TaskFightingActor) heal(character *character.Character, fightResult *FightResult) error {
	if fightResult.CharacterHpLoss >= character.Hp {
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
		useSchema := buildUseSchema(character, healingItems, targetHealing)
		healRes, err := actor.characterService.Use(character, useSchema)
		if err != nil {
			actor.logger.Error("Could not use item", "use", *useSchema, "error", err)
			return err
		}
		time.Sleep(healRes.CooldownRemaining)
	}
	return nil
}
