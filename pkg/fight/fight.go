package fight

import (
	"math"

	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
)

func CalculateFightResult(character *artifactsmmo.CharacterSchema, monster *artifactsmmo.MonsterSchema) {
	characterDmg := calculateCharacterDamage(character, monster)
	_ = monster.Hp / characterDmg
}

func calculateCharacterDamage(character *artifactsmmo.CharacterSchema, monster *artifactsmmo.MonsterSchema) int {
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
	dmg := int(math.Round(attackAir*(1+dmgAir/100)*(1-resAir/100))) +
		int(math.Round(attackFire*(1+dmgFire/100)*(1-resFire/100))) +
		int(math.Round(attackEarth*(1+dmgEarth/100)*(1-resEarth/100))) +
		int(math.Round(attackWater*(1+dmgWater/100)*(1-resWater/100)))
	return dmg
}
