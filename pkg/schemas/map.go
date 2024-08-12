package schemas

type MapContentSchema struct {
  Type string `json:"type"`
  Code string `json:"content"`
}

type MapSchema struct {
	Name    string             `json:"name"`
	Skin    string             `json:"skin"`
	X       int32              `json:"x"`
	Y       int32              `json:"y"`
	Content MapContentSchema `json:"content"`
}
