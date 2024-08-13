package mappers

import (
	"github.com/NChitty/artifactsmmo/pkg/models"
	"github.com/NChitty/artifactsmmo/pkg/schemas"
)

func ToInventorySlotModel(schema schemas.InventorySlotSchema) models.InventorySlot {
	return models.InventorySlot{
		Id:       schema.Id,
		ItemCode: schema.ItemCode,
		Quantity: schema.Quantity,
	}
}

func ToCharacterModel(schema schemas.CharacterSchema, model *models.Character) {
	model.Name = schema.Name
	model.Skin = schema.Skin
	model.Level = schema.Level
	model.Experience = schema.Experience
	model.ExperienceNeeded = schema.ExperienceNeeded
	model.Gold = schema.Gold
	model.Mining = models.Skill{
		Level:            schema.MiningLevel,
		Experience:       schema.MiningExperience,
		ExperienceNeeded: schema.MiningExperienceNeeded,
	}
	model.Woodcutting = models.Skill{
		Level:            schema.WoodcuttingLevel,
		Experience:       schema.WoodcuttingExperience,
		ExperienceNeeded: schema.WoodcuttingExperienceNeeded,
	}
	model.Fishing = models.Skill{
		Level:            schema.FishingLevel,
		Experience:       schema.FishingExperience,
		ExperienceNeeded: schema.FishingExperienceNeeded,
	}
	model.WeaponCrafting = models.Skill{
		Level:            schema.WeaponCraftingLevel,
		Experience:       schema.WeaponCraftingExperience,
		ExperienceNeeded: schema.WeaponCraftingExperienceNeeded,
	}
	model.GearCrafting = models.Skill{
		Level:            schema.GearCraftingLevel,
		Experience:       schema.GearCraftingExperience,
		ExperienceNeeded: schema.GearCraftingExperienceNeeded,
	}
	model.JewelryCrafting = models.Skill{
		Level:            schema.JewelryCraftingLevel,
		Experience:       schema.JewelryCraftingExperience,
		ExperienceNeeded: schema.JewelryCraftingExperienceNeeded,
	}
	model.Cooking = models.Skill{
		Level:            schema.CookingLevel,
		Experience:       schema.CookingExperience,
		ExperienceNeeded: schema.CookingExperienceNeeded,
	}
	model.HealthPoints = schema.HealthPoints
	model.Haste = schema.Haste
	model.Air = models.CombatStats{
		Attack:     schema.AirAttack,
		Damage:     schema.AirDamage,
		Resistance: schema.AirResistance,
	}
	model.Earth = models.CombatStats{
		Attack:     schema.EarthAttack,
		Damage:     schema.EarthDamage,
		Resistance: schema.EarthResistance}
	model.Fire = models.CombatStats{
		Attack:     schema.FireAttack,
		Damage:     schema.FireDamage,
		Resistance: schema.FireResistance}
	model.Water = models.CombatStats{
		Attack:     schema.WaterAttack,
		Damage:     schema.WaterDamage,
		Resistance: schema.WaterResistance,
	}

	model.Position = models.Position{X: schema.X, Y: schema.Y}

	var inventorySlots []models.InventorySlot
	for _, slotSchema := range schema.Inventory {
		slotModel := ToInventorySlotModel(slotSchema)
		inventorySlots = append(inventorySlots, slotModel)
	}

	model.Inventory = models.Inventory{
		WeaponSlot:              schema.WeaponSlot,
		ShieldSlot:              schema.ShieldSlot,
		HelmetSlot:              schema.HelmetSlot,
		BodyArmorSlot:           schema.BodyArmorSlot,
		LegArmorSlot:            schema.LegArmorSlot,
		BootsSlot:               schema.BootsSlot,
		RingSlot1:               schema.RingSlot1,
		RingSlot2:               schema.RingSlot2,
		AmuletSlot:              schema.AmuletSlot,
		ArtifactSlot1:           schema.ArtifactSlot1,
		ArtifactSlot2:           schema.ArtifactSlot2,
		ArtifactSlot3:           schema.ArtifactSlot3,
		ConsumableSlot1:         schema.ConsumableSlot1,
		ConsumableSlot1Quantity: schema.ConsumableSlot1Quantity,
		ConsumableSlot2:         schema.ConsumableSlot2,
		ConsumableSlot2Quantity: schema.ConsumableSlot2Quantity,
		InventorySize:           schema.InventorySize,
		Inventory:               inventorySlots,
	}
	model.Task = schema.Task
	model.TaskType = schema.TaskType
	model.TaskProgress = schema.TaskProgress
	model.TaskTotal = schema.TaskTotal
}
