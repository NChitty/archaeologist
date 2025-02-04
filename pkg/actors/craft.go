package actors

import (
	"github.com/NChitty/archaeologist/pkg/character"
	"github.com/NChitty/archaeologist/pkg/artifacts"
)

type CraftingActor struct {
  GoalCode string
  GoalQuantity string
  character *character.Character
  stack []Actor
}

func (actor *CraftingActor) Do() {
  // lookup goal
  artifacts.GetItem(actor.GoalCode)
}
