package actions

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/NChitty/artifactsmmo/pkg/schemas"
)

type BlockedHitsSchema struct {
	Fire  uint32 `json:"fire"`
	Earth uint32 `json:"earth"`
	Water uint32 `json:"water"`
	Air   uint32 `json:"air"`
	Total uint32 `json:"total"`
}

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

type FightSchema struct {
	Experience         int32             `json:"xp"`
	Gold               int32             `json:"gold"`
	Drops              []schemas.DropSchema      `json:"drops"`
	Turns              uint32            `json:"turns"`
	MonsterBlockedHits BlockedHitsSchema `json:"monster_blocked_hits"`
	PlayerBlockedHits  BlockedHitsSchema `json:"player_blocked_hits"`
	Logs               []string          `json:"logs"`
	Result             FightResult       `json:"result"`
}
