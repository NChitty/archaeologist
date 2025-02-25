package effects

import (
	"log/slog"

	"github.com/NChitty/archaeologist/pkg/models/character"
	"github.com/NChitty/archaeologist/pkg/models/item"
)

type RestoreEffect struct {
	Quantity int
	Value    int
}

type EffectAccumulator struct {
	effects map[item.Effect]int
	restore []RestoreEffect
	logger  *slog.Logger
}

func New(logger *slog.Logger) *EffectAccumulator {
	effectAccumulator := EffectAccumulator{
		effects: map[item.Effect]int{},
		restore: []RestoreEffect{},
		logger:  logger,
	}
	return &effectAccumulator
}

func (accumulator *EffectAccumulator) Accumulate(equipment *[]character.Equipment) {
	for _, equipmentItem := range *equipment {
		for _, effectSchema := range *equipmentItem.Item.Effects {
			itemEffect := item.GetEffect(effectSchema.Code)
			if itemEffect == item.RestoreEffect {
				accumulator.restore = append(accumulator.restore, RestoreEffect{equipmentItem.Quantity, effectSchema.Value})
				continue
			}
			accumulator.effects[itemEffect] += effectSchema.Value
		}
	}
}

func (accumulator *EffectAccumulator) Reset() {
  accumulator.effects = map[item.Effect]int{}
  accumulator.restore = []RestoreEffect{}
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

func (accumulator *EffectAccumulator) GetRestoreEffect(restoreTurns int) int {
  for _, effect := range accumulator.restore {
    if effect.Quantity > restoreTurns {
      return effect.Value
    }
  }
  return 0
}

func (accumulator *EffectAccumulator) CanRestore() bool {
  return len(accumulator.restore) > 0
}
