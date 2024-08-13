package schemas

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Slot int

const (
	Weapon = iota
	Shield
	Helmet
	BodyArmor
	LegArmor
	Boots
	Ring1
	Ring2
	Amulet
	Artifact1
	Artifact2
	Artifact3
	Consumable1
	Consumable2
)

var slotName = map[Slot]string{
	Weapon:      "weapon",
	Shield:      "shield",
	Helmet:      "helmt",
	BodyArmor:   "body_armor",
	LegArmor:    "leg_armor",
	Boots:       "boots",
	Ring1:       "ring1",
	Ring2:       "ring2",
	Amulet:      "amulet",
	Artifact1:   "artifact1",
	Artifact2:   "artifact2",
	Artifact3:   "artifact3",
	Consumable1: "consumable1",
	Consumable2: "consumable2",
}

var slotValue = map[string]Slot{
	"weapon":      Weapon,
	"shield":      Shield,
	"helmt":       Helmet,
	"body_armor":  BodyArmor,
	"leg_armor":   LegArmor,
	"boots":       Boots,
	"ring1":       Ring1,
	"ring2":       Ring2,
	"amulet":      Amulet,
	"artifact1":   Artifact1,
	"artifact2":   Artifact2,
	"artifact3":   Artifact3,
	"consumable1": Consumable1,
	"consumable2": Consumable2,
}

func (s Slot) String() string {
	return slotName[s]
}

func ParseSlot(string string) (Slot, error) {
	string = strings.TrimSpace(strings.ToLower(string))
	value, ok := slotValue[string]
	if !ok {
		return Slot(0), fmt.Errorf("%q is not a valid Slot", string)
	}
	return Slot(value), nil
}

func (s *Slot) UnmarshalJSON(data []byte) error {
	var slots string
	if err := json.Unmarshal(data, &slots); err != nil {
		return err
	}
	parsed, err := ParseSlot(slots)
	if err != nil {
		return err
	}
	*s = parsed
	return nil
}

type ItemEffectSchema struct {
	Name  string `json:"name"`
	Value int32  `json:"value"`
}

type SimpleItemSchema struct {
	Code     string `json:"code"`
	Quantity uint32 `json:"quantity"`
}

type DropSchema SimpleItemSchema

type DropRateSchema struct {
	Code            string `json:"code"`
	Rate            int32  `json:"rate"`
	MinimumQuantity int32  `json:"min_quantity"`
	MaximumQuantity int32  `json:"max_quantity"`
}

type CraftSchema struct {
	Skill    Skill               `json:"skill"`
	Level    *int32              `json:"level,omitempty"`
	Items    []*SimpleItemSchema `json:"items,omitempty"`
	Quantity *int32              `json:"quantity,omitempty"`
}

type ItemSchema struct {
	Name        string              `json:"name"`
	Code        string              `json:"code"`
	Level       uint32              `json:"level"`
	Type        string              `json:"type"`
	Subtype     string              `json:"subtype"`
	Description string              `json:"description"`
	Effects     []*ItemEffectSchema `json:"effects"`
	Craft       *CraftSchema        `json:"craft"`
}
