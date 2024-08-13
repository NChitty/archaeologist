package actions

import "github.com/NChitty/archaeologist/pkg/schemas"

type BlockedHitsSchema struct {
	Fire  uint32 `json:"fire"`
	Earth uint32 `json:"earth"`
	Water uint32 `json:"water"`
	Air   uint32 `json:"air"`
	Total uint32 `json:"total"`
}

type FightSchema struct {
	Experience         int32                `json:"xp"`
	Gold               int32                `json:"gold"`
	Drops              []schemas.DropSchema `json:"drops"`
	Turns              uint32               `json:"turns"`
	MonsterBlockedHits BlockedHitsSchema    `json:"monster_blocked_hits"`
	PlayerBlockedHits  BlockedHitsSchema    `json:"player_blocked_hits"`
	Logs               []string             `json:"logs"`
	Result             schemas.FightResult  `json:"result"`
}

type CharacterFightDataSchema struct {
	Cooldown  schemas.CooldownSchema  `json:"cooldown"`
	Fight     FightSchema             `json:"fight"`
	Character schemas.CharacterSchema `json:"character"`
}
