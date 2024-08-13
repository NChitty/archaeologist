package schemas

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Enum of allowable skill values
type Skill int

const (
	WeaponCrafting = iota
	GearCrafting
	JewelryCrafting
	Cooking
	Woodcutting
	Mining
	Fishing
)

var skillName = map[Skill]string{
	WeaponCrafting:  "weaponcrafting",
	GearCrafting:    "gearcrafting",
	JewelryCrafting: "jewelrycrafting",
	Cooking:         "cooking",
	Woodcutting:     "woodcutting",
	Mining:          "mining",
	Fishing:         "fishing",
}

var skillValue = map[string]Skill{
	"weaponcrafting":  WeaponCrafting,
	"gearcrafting":    GearCrafting,
	"jewelrycrafting": JewelryCrafting,
	"cooking":         Cooking,
	"woodcutting":     Woodcutting,
	"mining":          Mining,
	"fishing":         Fishing,
}

func (s Skill) String() string {
	return skillName[s]
}

func ParseSkill(string string) (Skill, error) {
	string = strings.TrimSpace(strings.ToLower(string))
	value, ok := skillValue[string]
	if !ok {
		return Skill(0), fmt.Errorf("%q is not a valid Skill", string)
	}
	return Skill(value), nil
}

func (s *Skill) UnmarshalJSON(data []byte) error {
	var skills string
	if err := json.Unmarshal(data, &skills); err != nil {
		return err
	}
	parsed, err := ParseSkill(skills)
	if err != nil {
		return err
	}
	*s = parsed
	return nil
}

// Enum of allowable skin values
type Skin int

const (
	Men1 = iota
	Men2
	Men3
	Women1
	Women2
	Women3
)

var skinName = map[Skin]string{
	Men1:   "men1",
	Men2:   "men2",
	Men3:   "men3",
	Women1: "women1",
	Women2: "women2",
	Women3: "women3",
}

var skinValue = map[string]Skin{
	"men1":   Men1,
	"men2":   Men2,
	"men3":   Men3,
	"women1": Women1,
	"women2": Women2,
	"women3": Women3,
}

func (s Skin) String() string {
	return skinName[s]
}

func ParseSkin(string string) (Skin, error) {
	string = strings.TrimSpace(strings.ToLower(string))
	value, ok := skinValue[string]
	if !ok {
		return Skin(0), fmt.Errorf("%q is not a valid skin", string)
	}
	return Skin(value), nil
}

func (s *Skin) UnmarshalJSON(data []byte) error {
	var skins string
	if err := json.Unmarshal(data, &skins); err != nil {
		return err
	}
	parsed, err := ParseSkin(skins)
	if err != nil {
		return err
	}
	*s = parsed
	return nil
}

// Enum of fight results (allowable values: "win" or "lose")
type FightResult int

const (
	Win = iota
	Lose
)

var fightResultName = map[FightResult]string{
	Win:  "win",
	Lose: "lose",
}

var fightResultValue = map[string]FightResult{
	"win":  Win,
	"lose": Lose,
}

func ParseFightResult(string string) (FightResult, error) {
	string = strings.TrimSpace(strings.ToLower(string))
	value, ok := fightResultValue[string]
	if !ok {
		return FightResult(0), fmt.Errorf("%q is not a valid FightResult", string)
	}
	return FightResult(value), nil
}

func (s *FightResult) UnmarshalJSON(data []byte) error {
	var fightResults string
	if err := json.Unmarshal(data, &fightResults); err != nil {
		return err
	}
	parsed, err := ParseFightResult(fightResults)
	if err != nil {
		return err
	}
	*s = parsed
	return nil
}

// Enum of task types
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
	value, ok := taskTypeValue[string]
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
