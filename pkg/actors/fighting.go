package actors

import (
	"errors"
	"log/slog"
	"math"
	"sync"

	"github.com/NChitty/archaeologist/pkg/models/character"
	"github.com/NChitty/archaeologist/pkg/models/item"
	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
)

var ExitOnRest error = errors.New("Exit on rest")

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

func buildUseSchema(
	character *character.Character,
	healingItems map[string]*artifactsmmo.ItemSchema,
	minHealing int,
	targetHealing int,
) *artifactsmmo.SimpleItemSchema {
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
		// healing > min and closest to target
		if present && (minHealing/healing)+1 >= character.Inventory[code].Quantity {
			// this item cannot heal enough
			continue
		}
		neededQty := targetHealing / healing
		qty := neededQty
		if neededQty > character.Inventory[code].Quantity {
			// do not have enough to reach targetHealing
			qty = character.Inventory[code].Quantity
		}
		delta := targetHealing - healing*qty
		if math.Abs(float64(min.delta)) > math.Abs(float64(delta)) {
			min.delta = delta
			min.code = code
			min.qty = qty
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
