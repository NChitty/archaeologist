package actions

import "github.com/NChitty/archaeologist/pkg/schemas"

type TaskSchema struct {
	Code  string           `json:"code"`
	Type  schemas.TaskType `json:"type"`
	Total int32            `json:"total"`
}

type TaskRewardSchema schemas.SimpleItemSchema

type TaskDataSchema struct {
	Cooldown  schemas.CooldownSchema  `json:"cooldown"`
	Task      TaskSchema              `json:"task"`
	Character schemas.CharacterSchema `json:"character"`
}

type TaskRewardDataSchema struct {
	Cooldown  schemas.CooldownSchema  `json:"cooldown"`
	Reward    TaskRewardSchema        `json:"reward"`
	Character schemas.CharacterSchema `json:"character"`
}
