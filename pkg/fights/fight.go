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
	RestoreTurns    int
	CharacterDmg    int
	MonsterDmg      int
}

func DEFAULT() FightResult {
	return FightResult{
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

func (service *FightService) CalculateFightResult(
	character *characters.Character,
	monster *artifactsmmo.MonsterSchema,
) (*FightResult, error) {
	equipment := service.itemService.GetCharacterEquipment(character)
	service.effectAccumulator.Reset()
	service.effectAccumulator.Accumulate(equipment)
	characterDmg := service.calculateCharacterDamage(monster)
	characterTurns, err := calculateTurns(monster.Hp, characterDmg)
	result := DEFAULT()
	if err != nil {
		return &result, err
	}
	monsterResult := service.calculateMonsterResult(monster, character.Character.MaxHp, characterTurns, characterDmg)
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

func (service *FightService) calculateCharacterDamage(monster *artifactsmmo.MonsterSchema) int {
	attackAir := float64(service.effectAccumulator.GetEffect(effects.AttackAir))
	attackFire := float64(service.effectAccumulator.GetEffect(effects.AttackFire))
	attackEarth := float64(service.effectAccumulator.GetEffect(effects.AttackEarth))
	attackWater := float64(service.effectAccumulator.GetEffect(effects.AttackWater))
	dmgAir := float64(service.effectAccumulator.GetEffect(effects.DmgAir))
	dmgFire := float64(service.effectAccumulator.GetEffect(effects.DmgFire))
	dmgEarth := float64(service.effectAccumulator.GetEffect(effects.DmgEarth))
	dmgWater := float64(service.effectAccumulator.GetEffect(effects.DmgWater))
	resAir := float64(monster.ResAir)
	resFire := float64(monster.ResFire)
	resEarth := float64(monster.ResEarth)
	resWater := float64(monster.ResWater)
	unblockedDmgAir := math.Round(attackAir * (1 + dmgAir/100))
	unblockedDmgFire := math.Round(attackFire * (1 + dmgFire/100))
	unblockedDmgEarth := math.Round(attackEarth * (1 + dmgEarth/100))
	unblockedDmgWater := math.Round(attackWater * (1 + dmgWater/100))
	dmg := int(math.Round(unblockedDmgAir*(1-resAir/100))) +
		int(math.Round(unblockedDmgFire*(1-resFire/100))) +
		int(math.Round(unblockedDmgEarth*(1-resEarth/100))) +
		int(math.Round(unblockedDmgWater*(1-resWater/100)))
	return dmg
}

func (service *FightService) calculateMonsterDamage(monster *artifactsmmo.MonsterSchema) int {
	airDmg := int(math.Round(
		float64(monster.AttackAir) *
			(1 - float64(service.effectAccumulator.GetEffect(effects.ResAir)+service.effectAccumulator.GetEffect(effects.BoostResAir))/float64(100))))
	fireDmg := int(math.Round(
		float64(monster.AttackFire) *
			(1 - float64(service.effectAccumulator.GetEffect(effects.ResFire)+service.effectAccumulator.GetEffect(effects.BoostResFire))/float64(100))))
	earthDmg := int(math.Round(
		float64(monster.AttackEarth) *
			(1 - float64(service.effectAccumulator.GetEffect(effects.ResEarth)+service.effectAccumulator.GetEffect(effects.BoostResEarth))/float64(100))))
	waterDmg := int(math.Round(
		float64(monster.AttackWater) *
			(1 - float64(service.effectAccumulator.GetEffect(effects.ResWater)+service.effectAccumulator.GetEffect(effects.BoostResWater))/float64(100))))
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

type monsterResult struct {
	monsterTurns    int
	restoreTurns    int
	characterHpLoss int
	monsterDmg      int
}

func (service *FightService) calculateMonsterResult(
  monster *artifactsmmo.MonsterSchema,
	maxCharacterHp int,
	maxCharacterTurn int,
	characterDmg int,
) monsterResult {
	monsterDmg := service.calculateMonsterDamage(monster)
	characterMaxHpWithBoost := maxCharacterHp + service.effectAccumulator.GetEffect(effects.BoostHp)
	halfCharacterMaxHpWithBoost := characterMaxHpWithBoost / 2
	if !service.effectAccumulator.CanRestore() {
		monsterTurn, err := calculateTurns(characterMaxHpWithBoost, service.calculateMonsterDamage(monster))

		if err != nil {
			return monsterResult{maxCharacterTurn - 1, 0, 0, 0}
		}

		monsterTotalDmg := monsterDmg * monsterTurn
		if monsterTurn > maxCharacterTurn {
			monsterTotalDmg = monsterDmg * (maxCharacterTurn - 1)
		}
		return monsterResult{monsterTurn, 0,
			max(0, monsterTotalDmg-(characterMaxHpWithBoost-maxCharacterHp)), monsterDmg}
	}
	halfMonsterTurn, err := calculateTurns(halfCharacterMaxHpWithBoost, service.calculateMonsterDamage(monster))
	if err != nil {
		return monsterResult{maxCharacterTurn - 1, 0, 0, 0}
	}
	if halfMonsterTurn >= maxCharacterTurn {
		return monsterResult{halfMonsterTurn * 2, 0,
			max(0, (maxCharacterTurn-1)*monsterDmg-(characterMaxHpWithBoost-maxCharacterHp)),
			monsterDmg}
	}
	monsterTurn := halfMonsterTurn
	characterHp := characterMaxHpWithBoost
	monsterHp := monster.Hp
	restoreTurns := 0
	for i := 1; characterHp >= 0 && monsterHp >= 0; i += 2 {
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
		max(0, maxCharacterHp-characterHp),
		monsterDmg}
}
