package fights

import (
	"errors"
	"log/slog"
	"math"

	"github.com/NChitty/archaeologist/pkg/actors"
	"github.com/NChitty/archaeologist/pkg/items/effects"
	"github.com/NChitty/archaeologist/pkg/models/character"
	"github.com/NChitty/archaeologist/pkg/models/item"
	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
)

type ItemAccessor interface {
	GetCharacterEquipment(character *character.Character) *[]character.Equipment
}

func defaultFightResult() actors.FightResult {
	return actors.FightResult{
		Win:             false,
		Turns:           100,
		CharacterTurns:  100,
		CharacterHpLoss: 1000,
		RestoreTurns:    1000,
		CharacterDmg:    0,
		MonsterDmg:      0,
	}
}

type FightService struct {
	effectAccumulator *effects.EffectAccumulator
	logger            *slog.Logger
	itemService       ItemAccessor
}

func NewFightService(
	effectAccumulator *effects.EffectAccumulator,
	logger *slog.Logger,
	itemService ItemAccessor,
) *FightService {
	return &FightService{
		effectAccumulator: effectAccumulator,
		logger:            logger,
		itemService:       itemService,
	}
}

// Calculates the fight result without random effects such as critical strikes (and thus lifesteal)
// and blocking.
func (service *FightService) CalculateFightResult(
	character *character.Character,
	monster *artifactsmmo.MonsterSchema,
) (*actors.FightResult, error) {
	equipment := service.itemService.GetCharacterEquipment(character)
	service.effectAccumulator.Reset()
	service.effectAccumulator.Accumulate(equipment)
	characterDmg := service.calculateCharacterDamage(monster)
	characterMaxTurns, err := calculateTurns(monster.Hp, characterDmg)
	result := defaultFightResult()
	if err != nil {
		return &result, err
	}

	if service.effectAccumulator.IsSimpleEffects() &&
		!service.effectAccumulator.CanRestore() &&
		len(*monster.Effects) == 0 {

		monsterResult := service.calculateSimpleMonsterResult(monster, character.MaxHp, characterMaxTurns)
		numberTurns := characterMaxTurns*2 - 1
		if monsterResult.monsterTurns < characterMaxTurns {
			numberTurns = monsterResult.monsterTurns * 2
		}
		return &actors.FightResult{
			Win:             monsterResult.monsterTurns >= monsterResult.characterTurns,
			Turns:           numberTurns,
			CharacterTurns:  characterMaxTurns,
			CharacterHpLoss: monsterResult.characterHpLoss,
			RestoreTurns:    monsterResult.restoreTurns,
			CharacterDmg:    characterDmg,
			MonsterDmg:      monsterResult.monsterDmg,
		}, nil
	}

	monsterResult := service.calculateMonsterResult(monster, character.MaxHp, characterDmg)
	numberTurns := monsterResult.characterTurns + monsterResult.monsterTurns
	return &actors.FightResult{
		Win:             monsterResult.monsterTurns < monsterResult.characterTurns,
		Turns:           numberTurns,
		CharacterTurns:  monsterResult.characterTurns,
		CharacterHpLoss: monsterResult.characterHpLoss,
		RestoreTurns:    monsterResult.restoreTurns,
		CharacterDmg:    characterDmg,
		MonsterDmg:      monsterResult.monsterDmg,
	}, nil
}

func (service *FightService) calculateCharacterDamage(monster *artifactsmmo.MonsterSchema) int {
	attackAir := float64(service.effectAccumulator.GetEffect(item.AttackAirEffect))
	attackFire := float64(service.effectAccumulator.GetEffect(item.AttackFireEffect))
	attackEarth := float64(service.effectAccumulator.GetEffect(item.AttackEarthEffect))
	attackWater := float64(service.effectAccumulator.GetEffect(item.AttackWaterEffect))
	dmg := float64(service.effectAccumulator.GetEffect(item.DmgEffect))
	dmgAir := float64(service.effectAccumulator.GetEffect(item.DmgAirEffect))
	dmgFire := float64(service.effectAccumulator.GetEffect(item.DmgFireEffect))
	dmgEarth := float64(service.effectAccumulator.GetEffect(item.DmgEarthEffect))
	dmgWater := float64(service.effectAccumulator.GetEffect(item.DmgWaterEffect))
	resAir := float64(monster.ResAir)
	resFire := float64(monster.ResFire)
	resEarth := float64(monster.ResEarth)
	resWater := float64(monster.ResWater)
	unblockedDmgAir := math.Round(attackAir * (1 + (dmgAir+dmg)/100))
	unblockedDmgFire := math.Round(attackFire * (1 + (dmgFire+dmg)/100))
	unblockedDmgEarth := math.Round(attackEarth * (1 + (dmgEarth+dmg)/100))
	unblockedDmgWater := math.Round(attackWater * (1 + (dmgWater+dmg)/100))
	finalDmg := int(math.Round(unblockedDmgAir*(1-resAir/100))) +
		int(math.Round(unblockedDmgFire*(1-resFire/100))) +
		int(math.Round(unblockedDmgEarth*(1-resEarth/100))) +
		int(math.Round(unblockedDmgWater*(1-resWater/100)))
	return finalDmg
}

func (service *FightService) calculateMonsterDamage(monster *artifactsmmo.MonsterSchema) int {
	airDmg := int(math.Round(
		float64(monster.AttackAir) *
			(1 - float64(service.effectAccumulator.GetEffect(item.ResAirEffect)+service.effectAccumulator.GetEffect(item.BoostResAirEffect))/float64(100))))
	fireDmg := int(math.Round(
		float64(monster.AttackFire) *
			(1 - float64(service.effectAccumulator.GetEffect(item.ResFireEffect)+service.effectAccumulator.GetEffect(item.BoostResFireEffect))/float64(100))))
	earthDmg := int(math.Round(
		float64(monster.AttackEarth) *
			(1 - float64(service.effectAccumulator.GetEffect(item.ResEarthEffect)+service.effectAccumulator.GetEffect(item.BoostResEarthEffect))/float64(100))))
	waterDmg := int(math.Round(
		float64(monster.AttackWater) *
			(1 - float64(service.effectAccumulator.GetEffect(item.ResWaterEffect)+service.effectAccumulator.GetEffect(item.BoostResWaterEffect))/float64(100))))
	return airDmg + fireDmg + earthDmg + waterDmg
}

func calculateTurns(hp int, dmg int) (int, error) {
	if dmg == 0 {
		return 100, errors.New("Dmg is zero")
	}
	if hp%dmg == 0 {
		return hp / dmg, nil
	}
	return hp/dmg + 1, nil
}

func (service *FightService) calculateSimpleMonsterResult(
	monster *artifactsmmo.MonsterSchema,
	maxCharacterHp int,
	maxCharacterTurn int,
) monsterResult {
	monsterDmg := service.calculateMonsterDamage(monster)
	characterMaxHpWithBoost := maxCharacterHp + service.effectAccumulator.GetEffect(item.BoostHpEffect)
	monsterTurn, err := calculateTurns(characterMaxHpWithBoost, service.calculateMonsterDamage(monster))

	if err != nil {
		return monsterResult{maxCharacterTurn, maxCharacterTurn - 1, 0, 0, 0}
	}

	monsterTotalDmg := monsterDmg * monsterTurn
	if monsterTurn > maxCharacterTurn {
		monsterTotalDmg = monsterDmg * (maxCharacterTurn - 1)
	}
	return monsterResult{maxCharacterTurn, monsterTurn, 0,
		max(0, monsterTotalDmg-(characterMaxHpWithBoost-maxCharacterHp)), monsterDmg}
}

type monsterResult struct {
	characterTurns  int
	monsterTurns    int
	restoreTurns    int
	characterHpLoss int
	monsterDmg      int
}

func (service *FightService) calculateMonsterResult(
	monster *artifactsmmo.MonsterSchema,
	maxCharacterHp int,
	characterDmg int,
) monsterResult {
	monsterDmg := service.calculateMonsterDamage(monster)
	characterMaxHpWithBoost := maxCharacterHp + service.effectAccumulator.GetEffect(item.BoostHpEffect)
	halfCharacterMaxHpWithBoost := characterMaxHpWithBoost / 2

	type statusEffects struct {
		poison int
		burn   float64
	}
	characterStatusEffects := statusEffects{
		poison: 0,
		burn:   0,
	}
	monsterStatusEffects := statusEffects{
		poison: 0,
		burn:   0,
	}
	var monsterBurnDeduction float64
	var characterBurnDeduction float64
	monsterTurn := 1
	characterTurn := 1
	characterHp := characterMaxHpWithBoost
	monsterHp := monster.Hp
	restoreTurns := 0
	cureTurns := 0
	for characterHp >= 0 || monsterHp >= 0 {
		if characterTurn == 1 {
			monsterStatusEffects.poison = service.effectAccumulator.GetEffect(item.PoisonEffect)
			monsterStatusEffects.burn = float64(service.effectAccumulator.GetEffect(item.BurnEffect)) / 100 *
				(float64(service.effectAccumulator.GetEffect(item.AttackAirEffect)) +
					float64(service.effectAccumulator.GetEffect(item.AttackEarthEffect)) +
					float64(service.effectAccumulator.GetEffect(item.AttackFireEffect)) +
					float64(service.effectAccumulator.GetEffect(item.AttackWaterEffect)))
			characterBurnDeduction = float64(monsterStatusEffects.burn) / 10
		}
		characterTurn++
		if characterHp < halfCharacterMaxHpWithBoost {
			restoreValue := service.effectAccumulator.GetRestoreEffect(restoreTurns)
			if restoreValue > 0 {
				restoreTurns++
			}
			characterHp += restoreValue
		}

		if characterStatusEffects.poison > 0 {
			cureValue := service.effectAccumulator.GetCureEffect(cureTurns)
			if cureValue > 0 {
				cureTurns++
			}
			characterStatusEffects.poison = characterStatusEffects.poison - cureValue
		}

		characterHp -= max(int(math.Round(characterStatusEffects.burn)), 0)
		characterStatusEffects.burn = float64(characterStatusEffects.burn) - monsterBurnDeduction

		characterHp -= max(characterStatusEffects.poison, 0)

		monsterHp -= characterDmg

		if characterHp <= 0 {
			break
		}

		if characterTurn%3 == 0 && service.effectAccumulator.GetEffect(item.HealingEffect) != 0 {
			characterHp = min(characterHp+int(math.Round(float64(service.effectAccumulator.GetEffect(item.HealingEffect))/100*float64(characterMaxHpWithBoost))), characterMaxHpWithBoost)
		}

		if monsterHp <= 0 {
			break
		}
		if monsterTurn == 1 {
			characterStatusEffects.poison = effects.GetEffect(monster.Effects, item.PoisonEffect)
			characterStatusEffects.burn = float64(effects.GetEffect(monster.Effects, item.BurnEffect)) / 100 *
				(float64(monster.AttackAir) +
					float64(monster.AttackEarth) +
					float64(monster.AttackFire) +
					float64(monster.AttackWater))
			monsterBurnDeduction = float64(characterStatusEffects.burn) / 10
		}
		monsterTurn++
		if monsterTurn == effects.GetEffect(monster.Effects, item.ReconstitutionEffect) {
			monsterHp = monster.Hp
		}
		monsterHp -= max(int(math.Round(monsterStatusEffects.burn)), 0)
		monsterStatusEffects.burn = float64(monsterStatusEffects.burn) - characterBurnDeduction
		monsterHp -= max(monsterStatusEffects.poison, 0)
		characterHp -= monsterDmg
		if characterHp <= 0 {
			break
		}
		if monsterHp <= 0 {
			break
		}
	}
	return monsterResult{
		characterTurn,
		monsterTurn,
		restoreTurns,
		max(0, maxCharacterHp-characterHp),
		monsterDmg}
}
