package effects

import (
	"log/slog"

	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
)

type RestoreEffect struct {
  quantity int
  value int
}

type EffectAccumulator struct {
  effects map[ItemEffect]int
  restore []RestoreEffect
  logger *slog.Logger
}

func New(logger *slog.Logger) *EffectAccumulator {
  effectAccumulator := EffectAccumulator {
    effects: map[ItemEffect]int{},
    restore: []RestoreEffect{},
    logger: logger,
  }
  return &effectAccumulator
}

func (accumulator *EffectAccumulator) Accumulate(itemEffects *[]artifactsmmo.ItemEffectSchema, quantity int) {
  for _, effectSchema := range *itemEffects {
    effect := GetEffect(effectSchema.Name)
    if effect == Restore {
      accumulator.restore = append(accumulator.restore, RestoreEffect{quantity, effectSchema.Value})
      continue
    }
    accumulator.effects[effect] += effectSchema.Value
  }
}

func (accumulator *EffectAccumulator) GetEffects() map[ItemEffect]int {
  return accumulator.effects
}
