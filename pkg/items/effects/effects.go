package effects

type ItemEffect int

const (
	Hp ItemEffect = iota
	BoostHp

	AttackAir
	AttackEarth
	AttackFire
	AttackWater

	DmgAir
	DmgEarth
	DmgFire
	DmgWater

	ResAir
	ResEarth
	ResFire
	ResWater

	BoostDmgAir
	BoostDmgEarth
	BoostDmgFire
	BoostDmgWater

	BoostResAir
	BoostResEarth
	BoostResFire
	BoostResWater

	Restore
)

var effectName = map[ItemEffect]string{
	Hp:      "hp",
	BoostHp: "boost_hp",

	AttackAir:   "attack_air",
	AttackEarth: "attack_earth",
	AttackFire:  "attack_fire",
	AttackWater: "attack_water",

	DmgAir:   "dmg_air",
	DmgEarth: "dmg_earth",
	DmgFire:  "dmg_fire",
	DmgWater: "dmg_water",

	ResAir:   "res_air",
	ResEarth: "res_earth",
	ResFire:  "res_fire",
	ResWater: "res_water",

	BoostDmgAir:   "boost_dmg_air",
	BoostDmgEarth: "boost_dmg_earth",
	BoostDmgFire:  "boost_dmg_fire",
	BoostDmgWater: "boost_dmg_water",

	BoostResAir:   "boost_res_air",
	BoostResEarth: "boost_res_earth",
	BoostResFire:  "boost_res_fire",
	BoostResWater: "boost_res_water",

	Restore: "restore",
}

var effect = map[string]ItemEffect{
	"hp":       Hp,
	"boost_hp": BoostHp,

	"attack_air":   AttackAir,
	"attack_earth": AttackEarth,
	"attack_fire":  AttackFire,
	"attack_water": AttackWater,

	"dmg_air":   DmgAir,
	"dmg_earth": DmgEarth,
	"dmg_fire":  DmgFire,
	"dmg_water": DmgWater,

	"res_air":   ResAir,
	"res_earth": ResEarth,
	"res_fire":  ResFire,
	"res_water": ResWater,

	"boost_dmg_air":   BoostDmgAir,
	"boost_dmg_earth": BoostDmgEarth,
	"boost_dmg_fire":  BoostDmgFire,
	"boost_dmg_water": BoostDmgWater,

	"boost_res_air":   BoostResAir,
	"boost_res_earth": BoostResEarth,
	"boost_res_fire":  BoostResFire,
	"boost_res_water": BoostResWater,

	"restore": Restore,
}

func (ie ItemEffect) GetEffectName() string {
	return effectName[ie]
}

func GetEffect(effectName string) ItemEffect {
	return effect[effectName]
}

func Effects() map[ItemEffect]string {
  return effectName
}
