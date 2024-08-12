package schemas

import "encoding/json"

type ResponseContainer struct {
	Data json.RawMessage `json:"data"`
}

type PageableSchema interface {
	CharacterSchema | MapSchema
}

type PagedResponseContainer[T PageableSchema] struct {
	Data  []T    `json:"data"`
	Total int32  `json:"total"`
	Page  int32  `json:"page"`
	Size  int32  `json:"size"`
	Pages *int32 `json:"pages,omitempty"`
}
