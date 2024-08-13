package actions

import "github.com/NChitty/archaeologist/pkg/schemas"

type EquipItemRequest struct {
	Code string       `json:"code"`
	Slot schemas.Slot `json:"slot"`
}

type UnequipItemRequest struct {
	Slot schemas.Slot `json:"slot"`
}

type DeleteItemRequest schemas.SimpleItemSchema
type RecycleItemRequest schemas.SimpleItemSchema

type RecyclingItemsSchema struct {
	Items []schemas.DropSchema `json:"items"`
}

type RecyclingDataSchema struct {
	Cooldown  schemas.CooldownSchema  `json:"cooldown"`
	Details   RecyclingItemsSchema    `json:"details"`
	Character schemas.CharacterSchema `json:"character"`
}

type EquipItemResponseSchema struct {
	Cooldown  schemas.CooldownSchema  `json:"cooldown"`
	Slot      schemas.Slot            `json:"slot"`
	Item      schemas.ItemSchema      `json:"item"`
	Character schemas.CharacterSchema `json:"character"`
}

type DeleteItemSchema struct {
	Cooldown  schemas.CooldownSchema   `json:"cooldown"`
	Item      schemas.SimpleItemSchema `json:"item"`
	Character schemas.CharacterSchema  `json:"character"`
}
