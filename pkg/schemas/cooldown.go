package schemas

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Action int

const (
	Movement = iota
)

var actionName = map[Action]string{
	Movement: "movement",
}

var actionValue = map[string]Action{
	"movement": Movement,
}

func (a Action) String() string {
	return actionName[a]
}

func ParseAction(string string) (Action, error) {
	string = strings.TrimSpace(strings.ToLower(string))
	value, ok := actionValue[string]
	if !ok {
		return Action(0), fmt.Errorf("%q is not a valid action", string)
	}
	return Action(value), nil
}

func (a *Action) UnmarshalJSON(data []byte) error {
	var string string
	if err := json.Unmarshal(data, &string); err != nil {
		return err
	}
	parsed, err := ParseAction(string)
	if err != nil {
		return err
	}
	*a = parsed
	return nil
}

type CooldownSchema struct {
	TotalSeconds     uint32    `json:"total_seconds"`
	RemainingSeconds int32     `json:"remaining_seconds"`
	Expiration       time.Time `json:"expiration"`
	Reason           Action    `json:"reason"`
}
