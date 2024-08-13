package actions

import "github.com/NChitty/artifactsmmo/pkg/schemas"

type BankItemDepositRequestSchema schemas.SimpleItemSchema

type BankItemWithdrawRequestSchema schemas.SimpleItemSchema

type BankGoldDepositRequestSchema schemas.SimpleItemSchema

type BankGoldWithdrawRequestSchema schemas.SimpleItemSchema

type BankItemSchema struct {
	Cooldown  schemas.CooldownSchema     `json:"cooldown"`
	Item      schemas.ItemSchema         `json:"item"`
	Bank      []schemas.SimpleItemSchema `json:"bank"`
	Character schemas.CharacterSchema    `json:"character"`
}

type GoldSchema struct {
	Quantity uint32 `json:"quantity"`
}

type GoldTransactinSchema struct {
	Cooldown  schemas.CooldownSchema  `json:"cooldown"`
	Bank      GoldSchema              `json:"bank"`
	Character schemas.CharacterSchema `json:"character"`
}
