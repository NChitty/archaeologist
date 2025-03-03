package item

type Effect int

const (
	HpEffect Effect = iota
	BoostHpEffect

	AttackAirEffect
	AttackEarthEffect
	AttackFireEffect
	AttackWaterEffect

	DmgEffect
	DmgAirEffect
	DmgEarthEffect
	DmgFireEffect
	DmgWaterEffect

	ResAirEffect
	ResEarthEffect
	ResFireEffect
	ResWaterEffect

	BoostDmgAirEffect
	BoostDmgEarthEffect
	BoostDmgFireEffect
	BoostDmgWaterEffect

	BoostResAirEffect
	BoostResEarthEffect
	BoostResFireEffect
	BoostResWaterEffect

	RestoreEffect

	BurnEffect
	HealEffect
	HealingEffect
	LifestealEffect
	PoisonEffect
	AntipoisonEffect
	CriticalStrikeEffect
	ReconstitutionEffect
)

var effectName = map[Effect]string{
	HpEffect:      "hp",
	BoostHpEffect: "boost_hp",

	AttackAirEffect:   "attack_air",
	AttackEarthEffect: "attack_earth",
	AttackFireEffect:  "attack_fire",
	AttackWaterEffect: "attack_water",

	DmgEffect:      "dmg",
	DmgAirEffect:   "dmg_air",
	DmgEarthEffect: "dmg_earth",
	DmgFireEffect:  "dmg_fire",
	DmgWaterEffect: "dmg_water",

	ResAirEffect:   "res_air",
	ResEarthEffect: "res_earth",
	ResFireEffect:  "res_fire",
	ResWaterEffect: "res_water",

	BoostDmgAirEffect:   "boost_dmg_air",
	BoostDmgEarthEffect: "boost_dmg_earth",
	BoostDmgFireEffect:  "boost_dmg_fire",
	BoostDmgWaterEffect: "boost_dmg_water",

	BoostResAirEffect:   "boost_res_air",
	BoostResEarthEffect: "boost_res_earth",
	BoostResFireEffect:  "boost_res_fire",
	BoostResWaterEffect: "boost_res_water",

	RestoreEffect: "restore",

	AntipoisonEffect:     "antipoison",
	BurnEffect:           "burn",
	CriticalStrikeEffect: "critical_strike",
	HealEffect:           "heal",
	HealingEffect:        "healing",
	LifestealEffect:      "lifesteal",
	PoisonEffect:         "poison",
	ReconstitutionEffect: "reconstitution",
}

var effect = map[string]Effect{
	"hp":       HpEffect,
	"boost_hp": BoostHpEffect,

	"attack_air":   AttackAirEffect,
	"attack_earth": AttackEarthEffect,
	"attack_fire":  AttackFireEffect,
	"attack_water": AttackWaterEffect,

	"dmg":       DmgEffect,
	"dmg_air":   DmgAirEffect,
	"dmg_earth": DmgEarthEffect,
	"dmg_fire":  DmgFireEffect,
	"dmg_water": DmgWaterEffect,

	"res_air":   ResAirEffect,
	"res_earth": ResEarthEffect,
	"res_fire":  ResFireEffect,
	"res_water": ResWaterEffect,

	"boost_dmg_air":   BoostDmgAirEffect,
	"boost_dmg_earth": BoostDmgEarthEffect,
	"boost_dmg_fire":  BoostDmgFireEffect,
	"boost_dmg_water": BoostDmgWaterEffect,

	"boost_res_air":   BoostResAirEffect,
	"boost_res_earth": BoostResEarthEffect,
	"boost_res_fire":  BoostResFireEffect,
	"boost_res_water": BoostResWaterEffect,

	"restore": RestoreEffect,

	"antipoison":      AntipoisonEffect,
	"burn":            BurnEffect,
	"critical_strike": CriticalStrikeEffect,
	"heal":            HealEffect,
	"healing":         HealingEffect,
	"lifesteal":       LifestealEffect,
	"poison":          PoisonEffect,
	"reconstitution":  ReconstitutionEffect,
}

func (ie Effect) GetEffectCode() string {
	return effectName[ie]
}

func GetEffect(effectCode string) Effect {
	return effect[effectCode]
}

func Effects() map[Effect]string {
	return effectName
}
