package schemas

type ResourceSchema struct {
	Name  string           `json:"name"`
	Code  string           `json:"code"`
	Skill Skill            `json:"skill"`
	Level int32            `json:"level"`
	Drops []DropRateSchema `json:"drops"`
}
