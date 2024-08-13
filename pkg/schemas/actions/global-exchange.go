package actions

import "github.com/NChitty/archaeologist/pkg/schemas"

type GlobalExchangeTransactionRequestSchema struct {
	Code       string `json:"code"`
	Quantity   int32  `json:"quantity"`
	Price      int32  `json:"price"`
}

type GlobalExchangeTransactionSchema struct {
	Code       string `json:"code"`
	Quantity   int32  `json:"quantity"`
	Price      int32  `json:"price"`
	TotalPrice int32  `json:"total_price"`
}

type GlobalExchangeTransactionListSchema struct {
  Cooldown schemas.CooldownSchema `json:"cooldown"`
  Transaction GlobalExchangeTransactionSchema `json:"transaction"`
  Character schemas.CharacterSchema `json:"character"`
}
