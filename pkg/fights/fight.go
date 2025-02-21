package fights

import (
	"errors"
	"math"

	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
)

type FightResult struct {
	Win             bool
	Turns           uint16
	CharacterTurns  uint16
	CharacterHpLoss uint16
	RestoreTurn     uint16
	CharacterDmg    uint16
	MonsterDmg      uint16
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

func CalculateFightResult(
	character *artifactsmmo.CharacterSchema,
	monster *artifactsmmo.MonsterSchema,
) *FightResult {
	characterDmg := calculateCharacterDamage(character, monster)
	_, err := calculateTurns(monster.Hp, characterDmg)
	result := DEFAULT()
	if err != nil {
		return &result
	}
	return &FightResult{}
}

func calculateCharacterDamage(
	character *artifactsmmo.CharacterSchema,
	monster *artifactsmmo.MonsterSchema,
) uint16 {
	attackAir := float64(character.AttackAir)
	attackFire := float64(character.AttackFire)
	attackEarth := float64(character.AttackEarth)
	attackWater := float64(character.AttackWater)
	dmgAir := float64(character.DmgAir)
	dmgFire := float64(character.DmgFire)
	dmgEarth := float64(character.DmgEarth)
	dmgWater := float64(character.DmgWater)
	resAir := float64(monster.ResAir)
	resFire := float64(monster.ResFire)
	resEarth := float64(monster.ResEarth)
	resWater := float64(monster.ResWater)
	dmg := uint16(math.Round(attackAir*(1+dmgAir/100)*(1-resAir/100))) +
		uint16(math.Round(attackFire*(1+dmgFire/100)*(1-resFire/100))) +
		uint16(math.Round(attackEarth*(1+dmgEarth/100)*(1-resEarth/100))) +
		uint16(math.Round(attackWater*(1+dmgWater/100)*(1-resWater/100)))
	return dmg
}

func calculateTurns(hp int, dmg uint16) (uint16, error) {
	if dmg == 0 {
		return 100, errors.New("Dmg is zero")
	}
	return uint16(math.Ceil(float64(hp) / float64(dmg))), nil
}

type monsterResult struct {
	monsterTurn     uint16
	restoreTurn     uint16
	characterHpLoss uint16
	characterDmg    uint16
}

func calculateMonsterResult(
	characterHp int,
	monster *artifactsmmo.MonsterSchema,
	maxCharacterTurn uint16,
	characterDmg uint16,
) monsterResult {
	monsterDmg := calculateMonsterDamage(monster, effectsCumulator)
	characterMaxHp := characterHp + effectsCumulator.getHp()
	characterMaxHpWithBoost := characterMaxHp + effectsCumulator.getBoostHp()
	halfCharacterMaxHpWithBoost := characterMaxHpWithBoost / 2
	if !effectsCumulator.isRestore() {
		monsterTurn := calculateTurns(characterMaxHpWithBoost, calculMonsterDamage(monster, effectsCumulator))

		monsterTotalDmg := monsterDmg * math.Max(maxCharacterTurn-1, monsterTurn)
		return monsterResult{monsterTurn, 0,
			uint16(math.Max(0, monsterTotalDmg-(characterMaxHpWithBoost-characterMaxHp))), monsterDmg}
	}
	halfMonsterTurn := calculTurns(halfCharacterMaxHpWithBoost, calculMonsterDamage(monster, effectsCumulator))
	if halfMonsterTurn >= maxCharacterTurn {
		return monsterResult{halfMonsterTurn * 2, 0,
			uint16(math.Max(0, (maxCharacterTurn-1)*monsterDmg-(characterMaxHpWithBoost-characterMaxHp))),
			monsterDmg}
	}
	monsterTurn := halfMonsterTurn
	characterHp = characterMaxHpWithBoost - halfMonsterTurn*monsterDmg
	monsterHp := monster.Hp - halfMonsterTurn*characterDmg
	restoreTurn := 1
	for characterHp >= 0 && monsterHp >= 0 {
		if characterHp < halfCharacterMaxHpWithBoost {
			restoreValue := effectsCumulator.getRestoreEffectValue(restoreTurn)
			if restoreValue > 0 {
				restoreTurn++
			}
			characterHp += restoreValue
		}
		monsterTurn++
		monsterHp -= characterDmg
		if monsterHp > 0 {
			characterHp -= monsterDmg
		}
	}
	return monsterResult{monsterTurn, uint16(restoreTurn - 1), uint16(math.Max(0, (characterMaxHp - characterHp))),
		monsterDmg}
}
