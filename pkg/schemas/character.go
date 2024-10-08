package schemas

import "time"

type InventorySlotSchema struct {
	Id       int32  `json:"slot"`
	ItemCode string `json:"code"`
	Quantity int32  `json:"quantity"`
}

type CharacterSchema struct {
	Name                            string                `json:"name"`
	Skin                            Skin                  `json:"skin"`
	Level                           int32                 `json:"level"`
	Experience                      int32                 `json:"xp"`
	ExperienceNeeded                int32                 `json:"max_xp"`
	Gold                            int64                 `json:"gold"`
	MiningLevel                     int32                 `json:"mining_level"`
	MiningExperience                int32                 `json:"mining_xp"`
	MiningExperienceNeeded          int32                 `json:"mining_max_xp"`
	WoodcuttingLevel                int32                 `json:"woodcutting_level"`
	WoodcuttingExperience           int32                 `json:"woodcutting_xp"`
	WoodcuttingExperienceNeeded     int32                 `json:"woodcutting_max_xp"`
	FishingLevel                    int32                 `json:"fishing_level"`
	FishingExperience               int32                 `json:"fishing_xp"`
	FishingExperienceNeeded         int32                 `json:"fishing_max_xp"`
	WeaponCraftingLevel             int32                 `json:"weaponcrafting_level"`
	WeaponCraftingExperience        int32                 `json:"weaponcrafting_xp"`
	WeaponCraftingExperienceNeeded  int32                 `json:"weaponcrafting_max_xp"`
	GearCraftingLevel               int32                 `json:"gearcrafting_level"`
	GearCraftingExperience          int32                 `json:"gearcrafting_xp"`
	GearCraftingExperienceNeeded    int32                 `json:"gearcrafting_max_xp"`
	JewelryCraftingLevel            int32                 `json:"jewelrycrafting_level"`
	JewelryCraftingExperience       int32                 `json:"jewelrycrafting_xp"`
	JewelryCraftingExperienceNeeded int32                 `json:"jewelrycrafting_max_xp"`
	CookingLevel                    int32                 `json:"cooking_level"`
	CookingExperience               int32                 `json:"cooking_xp"`
	CookingExperienceNeeded         int32                 `json:"cooking_max_xp"`
	HealthPoints                    int32                 `json:"hp"`
	Haste                           int32                 `json:"haste"`
	FireAttack                      int32                 `json:"attack_fire"`
	EarthAttack                     int32                 `json:"attack_earth"`
	WaterAttack                     int32                 `json:"attack_water"`
	AirAttack                       int32                 `json:"attack_air"`
	FireDamage                      int32                 `json:"dmg_fire"`
	EarthDamage                     int32                 `json:"dmg_earth"`
	WaterDamage                     int32                 `json:"dmg_water"`
	AirDamage                       int32                 `json:"dmg_air"`
	FireResistance                  int32                 `json:"res_fire"`
	EarthResistance                 int32                 `json:"res_earth"`
	WaterResistance                 int32                 `json:"res_water"`
	AirResistance                   int32                 `json:"res_air"`
	X                               int32                 `json:"x"`
	Y                               int32                 `json:"y"`
	Cooldown                        int32                 `json:"cooldown"`
	CooldownExpiration              time.Time             `json:"cooldown_expiration"`
	WeaponSlot                      string                `json:"weapon_slot"`
	ShieldSlot                      string                `json:"shield_slot"`
	HelmetSlot                      string                `json:"helmet_slot"`
	BodyArmorSlot                   string                `json:"body_armor_slot"`
	LegArmorSlot                    string                `json:"leg_armor_slot"`
	BootsSlot                       string                `json:"boots_slot"`
	RingSlot1                       string                `json:"ring1_slot"`
	RingSlot2                       string                `json:"ring2_slot"`
	AmuletSlot                      string                `json:"amulet_slot"`
	ArtifactSlot1                   string                `json:"artifact1_slot"`
	ArtifactSlot2                   string                `json:"artifact2_slot"`
	ArtifactSlot3                   string                `json:"artifact3_slot"`
	ConsumableSlot1                 string                `json:"consumable1_slot"`
	ConsumableSlot1Quantity         int32                 `json:"consumable1_slot_quantity"`
	ConsumableSlot2                 string                `json:"consumable2_slot"`
	ConsumableSlot2Quantity         int32                 `json:"consumable2_slot_quantity"`
	Task                            string                `json:"task"`
	TaskType                        string                `json:"task_type"`
	TaskProgress                    int32                 `json:"task_progress"`
	TaskTotal                       int32                 `json:"task_total"`
	InventorySize                   int32                 `json:"inventory_max_items"`
	Inventory                       []InventorySlotSchema `json:"inventory"`
}

type LogSchema struct {
  Character string
  Account string
  Type string
  Description string
  Content string
  Cooldown int32
  Expiration time.Time
  CreatedAt time.Time
}
