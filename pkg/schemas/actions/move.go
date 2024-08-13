package actions

import "github.com/NChitty/archaeologist/pkg/schemas"

type PositionSchema struct {
	X int32 `json:"x"`
	Y int32 `json:"y"`
}

type DestinationSchema struct {
	Name    string `json:"name"`
	X       int32  `json:"x"`
	Y       int32  `json:"y"`
	Content string `json:"content"`
}

type CharacterMovementDataSchema struct {
	Cooldown    schemas.CooldownSchema    `json:"cooldown"`
	Destination DestinationSchema `json:"destination"`
	Character   schemas.CharacterSchema   `json:"CharacterSchema"`
}
