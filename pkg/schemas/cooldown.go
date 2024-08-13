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
	Fight
	Crafting
	Gathering
	BuyGlobalExchange
	SellGlobalExchange
	DeleteItem
	DepositBank
	WithdrawBank
	Equip
	Unequip
	Task
	Recycling
)

var actionName = map[Action]string{
	Movement:           "movement",
	Fight:              "fight",
	Crafting:           "crafting",
	Gathering:          "gathering",
	BuyGlobalExchange:  "buy_ge",
	SellGlobalExchange: "sell_ge",
	DeleteItem:         "delete_item",
	DepositBank:        "deposit_bank",
	WithdrawBank:       "withdraw_bank",
	Equip:              "equip",
	Unequip:            "unequip",
	Task:               "task",
	Recycling:          "recycling",
}

var actionValue = map[string]Action{
	"movement": Movement,
	"fight": Fight,
	"crafting": Crafting,
	"gathering": Gathering,
	"buy_ge": BuyGlobalExchange,
	"sell_ge": SellGlobalExchange,
	"delete_item": DeleteItem,
	"deposit_bank": DepositBank,
	"withdraw_bank": WithdrawBank,
	"equip": Equip,
	"unequip": Unequip,
	"task": Task,
	"recycling": Recycling,
}

func (a Action) String() string {
	return actionName[a]
}

func ParseAction(string string) (Action, error) {
	string = strings.TrimSpace(strings.ToLower(string))
	value, ok := actionValue[string]
	if !ok {
		return Action(0), fmt.Errorf("%q is not a valid Action", string)
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
