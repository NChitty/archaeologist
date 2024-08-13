package actions

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/NChitty/archaeologist/pkg/schemas"
)

type TaskType int

const (
	Monsters = iota
	Resources
	Crafts
)

var taskTypeName = map[TaskType]string{
	Monsters:  "monsters",
	Resources: "resources",
	Crafts:    "crafts",
}

var taskTypeValue = map[string]TaskType{
	"monsters":  Monsters,
	"resources": Resources,
	"crafts":    Crafts,
}

func ParseTaskType(string string) (TaskType, error) {
	string = strings.TrimSpace(strings.ToLower(string))
	value, ok := fightResultValue[string]
	if !ok {
		return TaskType(0), fmt.Errorf("%q is not a valid TaskType", string)
	}
	return TaskType(value), nil
}

func (s *TaskType) UnmarshalJSON(data []byte) error {
	var taskTypes string
	if err := json.Unmarshal(data, &taskTypes); err != nil {
		return err
	}
	parsed, err := ParseTaskType(taskTypes)
	if err != nil {
		return err
	}
	*s = parsed
	return nil
}

type TaskSchema struct {
	Code  string   `json:"code"`
	Type  TaskType `json:"type"`
	Total int32    `json:"total"`
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
