package effects

import (
	"log/slog"

	"github.com/NChitty/archaeologist/pkg/items"
)

type RestoreEffect struct {
	Quantity int
	Value    int
}

type EffectAccumulator struct {
	effects map[ItemEffect]int
	restore []RestoreEffect
	logger  *slog.Logger
}

func New(logger *slog.Logger) *EffectAccumulator {
	effectAccumulator := EffectAccumulator{
		effects: map[ItemEffect]int{},
		restore: []RestoreEffect{},
		logger:  logger,
	}
	return &effectAccumulator
}

func (accumulator *EffectAccumulator) Accumulate(equipment *[]items.Equipment) {
	for _, equipmentItem := range *equipment {
		for _, effectSchema := range *equipmentItem.Item.Effects {
			effect := GetEffect(effectSchema.Code)
			if effect == Restore {
				accumulator.restore = append(accumulator.restore, RestoreEffect{equipmentItem.Quantity, effectSchema.Value})
				continue
			}
			accumulator.effects[effect] += effectSchema.Value
		}
	}
}

func (accumulator *EffectAccumulator) Reset() {
  accumulator.effects = map[ItemEffect]int{}
  accumulator.restore = []RestoreEffect{}
}

func (accumulator *EffectAccumulator) GetEffects() map[ItemEffect]int {
	return accumulator.effects
}

func (accumulator *EffectAccumulator) GetEffect(effect ItemEffect) int {
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
