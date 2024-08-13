package models

import "github.com/NChitty/archaeologist/pkg/schemas"

type Position struct {
	X int32
	Y int32
}

type Skill struct {
	Level            int32
	Experience       int32
	ExperienceNeeded int32
}

type CombatStats struct {
	Attack     int32
	Damage     int32
	Resistance int32
}

type InventorySlot struct {
	Id       int32
	ItemCode string
	Quantity int32
}

type Inventory struct {
	WeaponSlot              string
	ShieldSlot              string
	HelmetSlot              string
	BodyArmorSlot           string
	LegArmorSlot            string
	BootsSlot               string
	RingSlot1               string
	RingSlot2               string
	AmuletSlot              string
	ArtifactSlot1           string
	ArtifactSlot2           string
	ArtifactSlot3           string
	ConsumableSlot1         string
	ConsumableSlot1Quantity int32
	ConsumableSlot2         string
	ConsumableSlot2Quantity int32
	InventorySize           int32
	Inventory               []InventorySlot
}

type Character struct {
	Name             string
	Skin             schemas.Skin
	Level            int32
	Experience       int32
	ExperienceNeeded int32
	Gold             int64
	Mining           Skill
	Woodcutting      Skill
	Fishing          Skill
	WeaponCrafting   Skill
	GearCrafting     Skill
	JewelryCrafting  Skill
	Cooking          Skill
	HealthPoints     int32
	Haste            int32
	Fire             CombatStats
	Earth            CombatStats
	Water            CombatStats
	Air              CombatStats
	Position         Position
	Inventory        Inventory
	Task             string
	TaskType         string
	TaskProgress     int32
	TaskTotal        int32
}
