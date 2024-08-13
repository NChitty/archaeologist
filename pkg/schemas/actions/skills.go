package actions

import "github.com/NChitty/archaeologist/pkg/schemas"

type SkillInfoSchema struct {
	Experience int32                `json:"xp"`
	Items      []schemas.DropSchema `json:"items"`
}

type SkillDataSchema struct {
	Cooldown  schemas.CooldownSchema  `json:"cooldown"`
	Details   SkillInfoSchema         `json:"details"`
	Character schemas.CharacterSchema `json:"character"`
}

type CraftingRequestSchema schemas.SimpleItemSchema
