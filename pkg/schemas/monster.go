package schemas

type MonsterSchema struct {
	Name            string           `json:"name"`
	Code            string           `json:"code"`
	Level           int32            `json:"level"`
	HealthPoints    int32            `json:"hp"`
	FireAttack      int32            `json:"attack_fire"`
	EarthAttack     int32            `json:"attack_earth"`
	WaterAttack     int32            `json:"attack_water"`
	AirAttack       int32            `json:"attack_air"`
	FireResistence  int32            `json:"res_fire"`
	EarthResistence int32            `json:"res_earth"`
	WaterResistence int32            `json:"res_water"`
	AirResistence   int32            `json:"res_air"`
	MinimumGold     int32            `json:"min_gold"`
	MaximumGold     int32            `json:"max_gold"`
	Drops           []DropRateSchema `json:"drops"`
}
