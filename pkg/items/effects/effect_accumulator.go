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
			effect := GetEffect(effectSchema.Name)
			if effect == Restore {
				accumulator.restore = append(accumulator.restore, RestoreEffect{equipmentItem.Quantity, effectSchema.Value})
				continue
			}
			accumulator.effects[effect] += effectSchema.Value
		}
	}
}

func (accumulator *EffectAccumulator) GetEffects() map[ItemEffect]int {
	return accumulator.effects
}

func (accumulator *EffectAccumulator) GetRestoreEffects() []RestoreEffect {
	return accumulator.restore
}
