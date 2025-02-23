package characters

import (
	"context"
	"errors"
	"log/slog"
	"time"

	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
)

type CharacterWrapper struct {
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
}

func FromSchema(schema artifactsmmo.CharacterSchema) *CharacterWrapper {

	return &CharacterWrapper{
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

type CharacterService struct {
	client *artifactsmmo.ClientWithResponses
	logger *slog.Logger
}

func NewCharacterService(
	logger *slog.Logger,
	client *artifactsmmo.ClientWithResponses,
) *CharacterService {
	character := &CharacterService{
		client: client,
		logger: logger,
	}
	return character
}

func (service *CharacterService) UpdateCharacter(character *CharacterWrapper) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	characterResp, err := service.client.GetCharacterCharactersNameGetWithResponse(
		ctx,
		character.Name,
	)
	if err != nil {
		service.logger.Error("Could not retrieve character", "error", err)
		return err
	}
	if characterResp.StatusCode() == 404 {
		service.logger.Error("Could not retrieve character", "name", character.Name)
		return errors.New("Could not retrieve character with name: " + character.Name)
	}

	*character = *FromSchema(characterResp.JSON200.Data)

	return nil
}

func (service *CharacterService) GetInventory(character *CharacterWrapper) (map[string]artifactsmmo.InventorySlot, error) {
	err := service.UpdateCharacter(character)
	if err != nil {
		return nil, err
	}
	return character.Inventory, nil
}

func (service *CharacterService) WaitCooldown(character *CharacterWrapper) error {
	err := service.UpdateCharacter(character)
	if err != nil {
		return err
	}
	if time.Now().Before(*character.CooldownExpiration) {
		service.logger.Debug(
			"Waiting for cooldown",
			"expiration", character.CooldownExpiration,
			"timeRemaining", character.CooldownExpiration.Sub(time.Now().UTC()),
		)
		time.Sleep(character.CooldownExpiration.Sub(time.Now()))
	}
	return nil
}
