package effects

import (
	"log/slog"

	"github.com/NChitty/archaeologist/pkg/models/character"
	"github.com/NChitty/archaeologist/pkg/models/item"
	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
)

type RestoreEffect struct {
	Quantity int
	Value    int
}

type EffectAccumulator struct {
	effects    map[item.Effect]int
	restore    []RestoreEffect
	antipoison []RestoreEffect
	logger     *slog.Logger
}

func New(logger *slog.Logger) *EffectAccumulator {
	effectAccumulator := EffectAccumulator{
		effects:    map[item.Effect]int{},
		restore:    []RestoreEffect{},
		antipoison: []RestoreEffect{},
		logger:     logger,
	}
	return &effectAccumulator
}

func GetEffect(effects *[]artifactsmmo.SimpleEffectSchema, search item.Effect) int {
	for _, effect := range *effects {
		if item.GetEffect(effect.Code) == search {
			return effect.Value
		}
	}
	return 0
}

func (accumulator *EffectAccumulator) Accumulate(equipment *[]character.Equipment) {
	for _, equipmentItem := range *equipment {
		for _, effectSchema := range *equipmentItem.Item.Effects {
			itemEffect := item.GetEffect(effectSchema.Code)
			if itemEffect == item.RestoreEffect {
				accumulator.restore = append(accumulator.restore, RestoreEffect{equipmentItem.Quantity, effectSchema.Value})
				continue
			}
			if itemEffect == item.AntipoisonEffect {
				accumulator.antipoison = append(accumulator.antipoison, RestoreEffect{equipmentItem.Quantity, effectSchema.Value})
				continue
			}
			accumulator.effects[itemEffect] += effectSchema.Value
		}
	}
}

func (accumulator *EffectAccumulator) Reset() {
	accumulator.effects = map[item.Effect]int{}
	accumulator.restore = []RestoreEffect{}
  accumulator.antipoison = []RestoreEffect{}
}

func (accumulator *EffectAccumulator) GetEffects() map[item.Effect]int {
	return accumulator.effects
}

func (accumulator *EffectAccumulator) GetEffect(effect item.Effect) int {
	return accumulator.effects[effect]
}

func (accumulator *EffectAccumulator) GetRestoreEffects() []RestoreEffect {
	return accumulator.restore
}

func (accumulator *EffectAccumulator) GetCureEffect(cureTurns int) int {
	for _, effect := range accumulator.antipoison {
		if effect.Quantity > cureTurns {
			return effect.Value
		}
	}
	return 0
}

func (accumulator *EffectAccumulator) GetRestoreEffect(restoreTurns int) int {
	for _, effect := range accumulator.restore {
		if effect.Quantity > restoreTurns {
			return effect.Value
		}
	}
	return 0
}

func (accumulator *EffectAccumulator) IsSimpleEffects() bool {
	_, hasBurnEffect := accumulator.effects[item.BurnEffect]
	_, hasHealingEffect := accumulator.effects[item.HealingEffect]
	_, hasLifestealEffect := accumulator.effects[item.LifestealEffect]
	_, hasPoisonEffect := accumulator.effects[item.PoisonEffect]
	_, hasReconstitutionEffect := accumulator.effects[item.ReconstitutionEffect]
	return !(hasBurnEffect ||
		hasHealingEffect ||
		hasLifestealEffect ||
		hasPoisonEffect ||
		hasReconstitutionEffect)
}

func (accumulator *EffectAccumulator) CanRestore() bool {
	return len(accumulator.restore) > 0
}

func (accumulator *EffectAccumulator) CanCure() bool {
	return len(accumulator.antipoison) > 0
}
