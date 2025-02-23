package actors

import (
	"errors"
	"log/slog"
	"math"
	"sync"
	"time"

	"github.com/NChitty/archaeologist/pkg/characters"
	"github.com/NChitty/archaeologist/pkg/fights"
	"github.com/NChitty/archaeologist/pkg/items"
	"github.com/NChitty/archaeologist/pkg/items/effects"
	"github.com/NChitty/archaeologist/pkg/maps"
	"github.com/NChitty/archaeologist/pkg/monsters"
	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
)

type TaskFightingActor struct {
	monster          string
	quantity         int
	exitOnRest       bool
	characterService *characters.CharacterService
	fightService     *fights.FightService
	itemAccessor     items.ItemAccessor
	mapAccessor      maps.MapAccessor
	monsterAccessor  monsters.MonsterAccessor
	logger           *slog.Logger
}

func NewTaskFightingActor(
	character *characters.CharacterWrapper,
	exitOnRest bool,
	characterService *characters.CharacterService,
	fightService *fights.FightService,
	itemAccessor items.ItemAccessor,
	mapAccessor maps.MapAccessor,
	monsterAccessor monsters.MonsterAccessor,
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
		monsterAccessor:  monsterAccessor,
		logger:           logger,
	}, nil
}

func (actor *TaskFightingActor) Do(character *characters.CharacterWrapper) error {
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

	maps, err := actor.mapAccessor.GetAllMaps(nil, &actor.monster, nil, nil)
	if err != nil {
		actor.logger.Error("Could not retrieve all map tiles potentially relevant to monster.", "monster", actor.monster, "error", err)
		return err
	}
	if len(maps) == 0 {
		actor.logger.Error("No maps with monster.", "monster", actor.monster)
		return errors.New("No maps with monster.")
	}

	var x, y int
	// todo another place for an optimizer
	for _, cell := range maps {
		x = cell.X
		y = cell.Y
		break
	}

	if x == 0 && y == 0 {
		actor.logger.Error("Could not find map cell to go to")
		return errors.New("Could not find map cell")
	}

	if character.X != x || character.Y != y {
		result, err := actor.characterService.Move(character, x, y)
		if err != nil {
			return err
		}
		time.Sleep(result.CooldownRemaining)
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
			return err
		}
	}

	return nil
}

func buildUseSchema(character *characters.CharacterWrapper, healingItems map[string]*artifactsmmo.ItemSchema, targetHealing int) *artifactsmmo.SimpleItemSchema {
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
	for code, item := range healingItems {
		effectMap := mapEffects(item.Effects)
		healing, present := effectMap[effects.Heal]
		if present && float64(targetHealing)/float64(healing) > 1 {
			hasQty := character.Inventory[item.Code].Quantity
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

func mapEffects(effectsSchema *[]artifactsmmo.SimpleEffectSchema) map[effects.ItemEffect]int {
	values := map[effects.ItemEffect]int{}
	for _, effect := range *effectsSchema {
		values[effects.GetEffect(effect.Code)] = effect.Value
	}
	return values
}

func getHealingItems(itemAccessor items.ItemAccessor, character *characters.CharacterWrapper, logger *slog.Logger) map[string]*artifactsmmo.ItemSchema {
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
			item, err := itemAccessor.GetItem(slot.Code)
			if err != nil {
				logger.Warn("Could not get item", "code", slot.Code, "error", err)
			}
			canHeal := false
			for _, effect := range *item.Effects {
				if effect.Code == effects.Heal.GetEffectName() {
					canHeal = true
					break
				}
			}
			if canHeal {
				lock.Lock()
				consumables[item.Code] = item
				lock.Unlock()
			}
		}()
	}
	wg.Wait()
	return consumables
}

func (actor *TaskFightingActor) heal(character *characters.CharacterWrapper, fightResult *fights.FightResult) error {
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
			return errors.New("Exit on rest")
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
