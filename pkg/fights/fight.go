package fights

import (
	"errors"
	"log/slog"
	"math"

	"github.com/NChitty/archaeologist/pkg/characters"
	"github.com/NChitty/archaeologist/pkg/items"
	"github.com/NChitty/archaeologist/pkg/items/effects"
	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
)

type FightResult struct {
	Win             bool
	Turns           int
	CharacterTurns  int
	CharacterHpLoss int
	RestoreTurn     int
	CharacterDmg    int
	MonsterDmg      int
}

func DEFAULT() FightResult {
	return FightResult{
		Win:             false,
		Turns:           100,
		CharacterTurns:  100,
		CharacterHpLoss: 1000,
		RestoreTurn:     1000,
		CharacterDmg:    0,
		MonsterDmg:      0,
	}
}

type FightService struct {
	Character *characters.Character
	Monster   *artifactsmmo.MonsterSchema

	effectAccumulator *effects.EffectAccumulator
	logger            *slog.Logger
	itemService       *items.ItemService
}

func NewFightService(
	effectAccumulator *effects.EffectAccumulator,
	logger *slog.Logger,
	itemService *items.ItemService,
) *FightService {
	return &FightService{
		effectAccumulator: effectAccumulator,
		logger:            logger,
		itemService:       itemService,
	}
}

func (service *FightService) CalculateFightResult() (*FightResult, error) {
	if service.Character == nil {
		return nil, errors.New("Character is nil")
	}
	if service.Monster == nil {
		return nil, errors.New("Monster is nil")
	}
	equipment := service.itemService.GetCharacterEquipment(service.Character)
	service.effectAccumulator.Reset()
	service.effectAccumulator.Accumulate(equipment)
	characterDmg := service.calculateCharacterDamage()
	characterTurns, err := calculateTurns(service.Monster.Hp, characterDmg)
	result := DEFAULT()
	if err != nil {
		return &result, err
	}
	monsterResult := service.calculateMonsterResult(characterTurns, characterDmg)
	numberTurns := characterTurns*2 - 1
	if monsterResult.monsterTurns < characterTurns {
		numberTurns = monsterResult.monsterTurns * 2
	}
	return &FightResult{
		monsterResult.monsterTurns >= characterTurns,
		numberTurns,
		characterTurns,
		monsterResult.characterHpLoss,
		monsterResult.restoreTurns,
		characterDmg,
		monsterResult.monsterDmg,
	}, nil
}

func (service *FightService) calculateCharacterDamage() int {
	attackAir := float64(service.effectAccumulator.GetEffects()[effects.AttackAir])
	attackFire := float64(service.effectAccumulator.GetEffects()[effects.AttackFire])
	attackEarth := float64(service.effectAccumulator.GetEffects()[effects.AttackEarth])
	attackWater := float64(service.effectAccumulator.GetEffects()[effects.AttackWater])
	dmgAir := float64(service.effectAccumulator.GetEffects()[effects.DmgAir])
	dmgFire := float64(service.effectAccumulator.GetEffects()[effects.DmgFire])
	dmgEarth := float64(service.effectAccumulator.GetEffects()[effects.DmgEarth])
	dmgWater := float64(service.effectAccumulator.GetEffects()[effects.DmgWater])
	resAir := float64(service.Monster.ResAir)
	resFire := float64(service.Monster.ResFire)
	resEarth := float64(service.Monster.ResEarth)
	resWater := float64(service.Monster.ResWater)
	dmg := int(math.Round(attackAir*(1+dmgAir/100)*(1-resAir/100))) +
		int(math.Round(attackFire*(1+dmgFire/100)*(1-resFire/100))) +
		int(math.Round(attackEarth*(1+dmgEarth/100)*(1-resEarth/100))) +
		int(math.Round(attackWater*(1+dmgWater/100)*(1-resWater/100)))
	return dmg
}

func (service *FightService) calculateMonsterDamage() int {
	airDmg := int(math.Round(
		float64(service.Monster.AttackAir) *
			(1 - float64(service.effectAccumulator.GetEffect(effects.ResAir)+service.effectAccumulator.GetEffect(effects.BoostResAir))/float64(100))))
	fireDmg := int(math.Round(
		float64(service.Monster.AttackFire) *
			(1 - float64(service.effectAccumulator.GetEffect(effects.ResFire)+service.effectAccumulator.GetEffect(effects.BoostResFire))/float64(100))))
	earthDmg := int(math.Round(
		float64(service.Monster.AttackEarth) *
			(1 - float64(service.effectAccumulator.GetEffect(effects.ResEarth)+service.effectAccumulator.GetEffect(effects.BoostResEarth))/float64(100))))
	waterDmg := int(math.Round(
		float64(service.Monster.AttackWater) *
			(1 - float64(service.effectAccumulator.GetEffect(effects.ResWater)+service.effectAccumulator.GetEffect(effects.BoostResWater))/float64(100))))
	return airDmg + fireDmg + earthDmg + waterDmg
}

func calculateTurns(hp int, dmg int) (int, error) {
	if dmg == 0 {
		return 100, errors.New("Dmg is zero")
	}
	return int(math.Ceil(float64(hp) / float64(dmg))), nil
}

type monsterResult struct {
	monsterTurns    int
	restoreTurns    int
	characterHpLoss int
	monsterDmg      int
}

func (service *FightService) calculateMonsterResult(
	maxCharacterTurn int,
	characterDmg int,
) monsterResult {
	monsterDmg := service.calculateMonsterDamage()
	characterMaxHp := service.Character.Character.Hp + service.effectAccumulator.GetEffects()[effects.Hp]
	characterMaxHpWithBoost := characterMaxHp + service.effectAccumulator.GetEffects()[effects.BoostHp]
	halfCharacterMaxHpWithBoost := characterMaxHpWithBoost / 2
	if !service.effectAccumulator.CanRestore() {
		monsterTurn, err := calculateTurns(characterMaxHpWithBoost, service.calculateMonsterDamage())

		if err != nil {
			return monsterResult{maxCharacterTurn - 1, 0, 0, 0}
		}

		monsterTotalDmg := monsterDmg * monsterTurn
		if monsterTurn > maxCharacterTurn {
			monsterTotalDmg = monsterDmg * (maxCharacterTurn - 1)
		}
		return monsterResult{monsterTurn, 0,
			max(0, monsterTotalDmg-(characterMaxHpWithBoost-characterMaxHp)), monsterDmg}
	}
	halfMonsterTurn, err := calculateTurns(halfCharacterMaxHpWithBoost, service.calculateMonsterDamage())
	if err != nil {
		return monsterResult{maxCharacterTurn - 1, 0, 0, 0}
	}
	if halfMonsterTurn >= maxCharacterTurn {
		return monsterResult{halfMonsterTurn * 2, 0,
			max(0, (maxCharacterTurn-1)*monsterDmg-(characterMaxHpWithBoost-characterMaxHp)),
			monsterDmg}
	}
	monsterTurn := halfMonsterTurn
	characterHp := characterMaxHpWithBoost - halfMonsterTurn*monsterDmg
	monsterHp := service.Monster.Hp - halfMonsterTurn*characterDmg
	restoreTurns := 0
	for characterHp >= 0 && monsterHp >= 0 {
		if characterHp < halfCharacterMaxHpWithBoost {
			restoreValue := service.effectAccumulator.GetRestoreEffect(restoreTurns)
			if restoreValue > 0 {
				restoreTurns++
			}
			characterHp += restoreValue
		}
		monsterHp -= characterDmg
		if monsterHp > 0 {
			monsterTurn++
			characterHp -= monsterDmg
		}
	}
	return monsterResult{
		monsterTurn,
		restoreTurns,
		max(0, characterMaxHp-characterHp),
		monsterDmg}
}
