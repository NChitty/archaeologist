package character

import (
	"time"

	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
)

type Character struct {
	Account              string
	AlchemyLevel         int
	AlchemyMaxXp         int
	AlchemyXp            int
	AmuletSlot           string
	Artifact1Slot        string
	Artifact2Slot        string
	Artifact3Slot        string
	AttackAir            int
	AttackEarth          int
	AttackFire           int
	AttackWater          int
	BagSlot              string
	BodyArmorSlot        string
	BootsSlot            string
	CookingLevel         int
	CookingMaxXp         int
	CookingXp            int
	Cooldown             int
	CooldownExpiration   *time.Time
	CriticalStrike       int
	Dmg                  int
	DmgAir               int
	DmgEarth             int
	DmgFire              int
	DmgWater             int
	FishingLevel         int
	FishingMaxXp         int
	FishingXp            int
	GearcraftingLevel    int
	GearcraftingMaxXp    int
	GearcraftingXp       int
	Gold                 int
	Haste                int
	HelmetSlot           string
	Hp                   int
	InventoryMaxItems    int
	JewelrycraftingLevel int
	JewelrycraftingMaxXp int
	JewelrycraftingXp    int
	LegArmorSlot         string
	Level                int
	MaxHp                int
	MaxXp                int
	MiningLevel          int
	MiningMaxXp          int
	MiningXp             int
	Name                 string
	Prospecting          int
	ResAir               int
	ResEarth             int
	ResFire              int
	ResWater             int
	Ring1Slot            string
	Ring2Slot            string
	RuneSlot             string
	ShieldSlot           string
	Skin                 artifactsmmo.CharacterSkin
	Speed                int
	Task                 string
	TaskProgress         int
	TaskTotal            int
	TaskType             string
	Utility1Slot         string
	Utility1SlotQuantity int
	Utility2Slot         string
	Utility2SlotQuantity int
	WeaponSlot           string
	WeaponcraftingLevel  int
	WeaponcraftingMaxXp  int
	WeaponcraftingXp     int
	Wisdom               int
	WoodcuttingLevel     int
	WoodcuttingMaxXp     int
	WoodcuttingXp        int
	X                    int
	Xp                   int
	Y                    int

	Inventory map[string]artifactsmmo.InventorySlot

  Token string
}

func FromSchema(schema artifactsmmo.CharacterSchema) *Character {

	return &Character{
		Account:              schema.Account,
		AlchemyLevel:         schema.AlchemyLevel,
		AlchemyMaxXp:         schema.AlchemyMaxXp,
		AlchemyXp:            schema.AlchemyXp,
		AmuletSlot:           schema.AmuletSlot,
		Artifact1Slot:        schema.Artifact1Slot,
		Artifact2Slot:        schema.Artifact2Slot,
		Artifact3Slot:        schema.Artifact3Slot,
		AttackAir:            schema.AttackAir,
		AttackEarth:          schema.AttackEarth,
		AttackFire:           schema.AttackFire,
		AttackWater:          schema.AttackWater,
		BagSlot:              schema.BagSlot,
		BodyArmorSlot:        schema.BodyArmorSlot,
		BootsSlot:            schema.BootsSlot,
		CookingLevel:         schema.CookingLevel,
		CookingMaxXp:         schema.CookingMaxXp,
		CookingXp:            schema.CookingXp,
		Cooldown:             schema.Cooldown,
		CooldownExpiration:   schema.CooldownExpiration,
		CriticalStrike:       schema.CriticalStrike,
		Dmg:                  schema.Dmg,
		DmgAir:               schema.DmgAir,
		DmgEarth:             schema.DmgEarth,
		DmgFire:              schema.DmgFire,
		DmgWater:             schema.DmgWater,
		FishingLevel:         schema.FishingLevel,
		FishingMaxXp:         schema.FishingMaxXp,
		FishingXp:            schema.FishingXp,
		GearcraftingLevel:    schema.GearcraftingLevel,
		GearcraftingMaxXp:    schema.GearcraftingMaxXp,
		GearcraftingXp:       schema.GearcraftingXp,
		Gold:                 schema.Gold,
		Haste:                schema.Haste,
		HelmetSlot:           schema.HelmetSlot,
		Hp:                   schema.Hp,
		InventoryMaxItems:    schema.InventoryMaxItems,
		JewelrycraftingLevel: schema.JewelrycraftingLevel,
		JewelrycraftingMaxXp: schema.JewelrycraftingMaxXp,
		JewelrycraftingXp:    schema.JewelrycraftingXp,
		LegArmorSlot:         schema.LegArmorSlot,
		Level:                schema.Level,
		MaxHp:                schema.MaxHp,
		MaxXp:                schema.MaxXp,
		MiningLevel:          schema.MiningLevel,
		MiningMaxXp:          schema.MiningMaxXp,
		MiningXp:             schema.MiningXp,
		Name:                 schema.Name,
		Prospecting:          schema.Prospecting,
		ResAir:               schema.ResAir,
		ResEarth:             schema.ResEarth,
		ResFire:              schema.ResFire,
		ResWater:             schema.ResWater,
		Ring1Slot:            schema.Ring1Slot,
		Ring2Slot:            schema.Ring2Slot,
		RuneSlot:             schema.RuneSlot,
		ShieldSlot:           schema.ShieldSlot,
		Skin:                 schema.Skin,
		Speed:                schema.Speed,
		Task:                 schema.Task,
		TaskProgress:         schema.TaskProgress,
		TaskTotal:            schema.TaskTotal,
		TaskType:             schema.TaskType,
		Utility1Slot:         schema.Utility1Slot,
		Utility1SlotQuantity: schema.Utility1SlotQuantity,
		Utility2Slot:         schema.Utility2Slot,
		Utility2SlotQuantity: schema.Utility2SlotQuantity,
		WeaponSlot:           schema.WeaponSlot,
		WeaponcraftingLevel:  schema.WeaponcraftingLevel,
		WeaponcraftingMaxXp:  schema.WeaponcraftingMaxXp,
		WeaponcraftingXp:     schema.WeaponcraftingXp,
		Wisdom:               schema.Wisdom,
		WoodcuttingLevel:     schema.WoodcuttingLevel,
		WoodcuttingMaxXp:     schema.WoodcuttingMaxXp,
		WoodcuttingXp:        schema.WoodcuttingXp,
		X:                    schema.X,
		Xp:                   schema.Xp,
		Y:                    schema.Y,

		Inventory: mapInventory(schema.Inventory),
	}
}

func (character *Character) FromSchema(schema artifactsmmo.CharacterSchema) {
	character.Account = schema.Account
	character.AlchemyLevel = schema.AlchemyLevel
	character.AlchemyMaxXp = schema.AlchemyMaxXp
	character.AlchemyXp = schema.AlchemyXp
	character.AmuletSlot = schema.AmuletSlot
	character.Artifact1Slot = schema.Artifact1Slot
	character.Artifact2Slot = schema.Artifact2Slot
	character.Artifact3Slot = schema.Artifact3Slot
	character.AttackAir = schema.AttackAir
	character.AttackEarth = schema.AttackEarth
	character.AttackFire = schema.AttackFire
	character.AttackWater = schema.AttackWater
	character.BagSlot = schema.BagSlot
	character.BodyArmorSlot = schema.BodyArmorSlot
	character.BootsSlot = schema.BootsSlot
	character.CookingLevel = schema.CookingLevel
	character.CookingMaxXp = schema.CookingMaxXp
	character.CookingXp = schema.CookingXp
	character.Cooldown = schema.Cooldown
	character.CooldownExpiration = schema.CooldownExpiration
	character.CriticalStrike = schema.CriticalStrike
	character.Dmg = schema.Dmg
	character.DmgAir = schema.DmgAir
	character.DmgEarth = schema.DmgEarth
	character.DmgFire = schema.DmgFire
	character.DmgWater = schema.DmgWater
	character.FishingLevel = schema.FishingLevel
	character.FishingMaxXp = schema.FishingMaxXp
	character.FishingXp = schema.FishingXp
	character.GearcraftingLevel = schema.GearcraftingLevel
	character.GearcraftingMaxXp = schema.GearcraftingMaxXp
	character.GearcraftingXp = schema.GearcraftingXp
	character.Gold = schema.Gold
	character.Haste = schema.Haste
	character.HelmetSlot = schema.HelmetSlot
	character.Hp = schema.Hp
	character.InventoryMaxItems = schema.InventoryMaxItems
	character.JewelrycraftingLevel = schema.JewelrycraftingLevel
	character.JewelrycraftingMaxXp = schema.JewelrycraftingMaxXp
	character.JewelrycraftingXp = schema.JewelrycraftingXp
	character.LegArmorSlot = schema.LegArmorSlot
	character.Level = schema.Level
	character.MaxHp = schema.MaxHp
	character.MaxXp = schema.MaxXp
	character.MiningLevel = schema.MiningLevel
	character.MiningMaxXp = schema.MiningMaxXp
	character.MiningXp = schema.MiningXp
	character.Name = schema.Name
	character.Prospecting = schema.Prospecting
	character.ResAir = schema.ResAir
	character.ResEarth = schema.ResEarth
	character.ResFire = schema.ResFire
	character.ResWater = schema.ResWater
	character.Ring1Slot = schema.Ring1Slot
	character.Ring2Slot = schema.Ring2Slot
	character.RuneSlot = schema.RuneSlot
	character.ShieldSlot = schema.ShieldSlot
	character.Skin = schema.Skin
	character.Speed = schema.Speed
	character.Task = schema.Task
	character.TaskProgress = schema.TaskProgress
	character.TaskTotal = schema.TaskTotal
	character.TaskType = schema.TaskType
	character.Utility1Slot = schema.Utility1Slot
	character.Utility1SlotQuantity = schema.Utility1SlotQuantity
	character.Utility2Slot = schema.Utility2Slot
	character.Utility2SlotQuantity = schema.Utility2SlotQuantity
	character.WeaponSlot = schema.WeaponSlot
	character.WeaponcraftingLevel = schema.WeaponcraftingLevel
	character.WeaponcraftingMaxXp = schema.WeaponcraftingMaxXp
	character.WeaponcraftingXp = schema.WeaponcraftingXp
	character.Wisdom = schema.Wisdom
	character.WoodcuttingLevel = schema.WoodcuttingLevel
	character.WoodcuttingMaxXp = schema.WoodcuttingMaxXp
	character.WoodcuttingXp = schema.WoodcuttingXp
	character.X = schema.X
	character.Xp = schema.Xp
	character.Y = schema.Y

	character.Inventory = mapInventory(schema.Inventory)
}

func mapInventory(inventory *[]artifactsmmo.InventorySlot) map[string]artifactsmmo.InventorySlot {
	inventoryMap := make(map[string]artifactsmmo.InventorySlot)
  if inventory == nil {
    return inventoryMap
  }
	for _, slot := range *inventory {
		inventoryMap[slot.Code] = slot
	}
	return inventoryMap
}

type Equipment struct {
	Item     artifactsmmo.ItemSchema
	Quantity int
}
