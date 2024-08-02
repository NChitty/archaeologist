package responses

import (
	"encoding/json"
)

type ResponseContainer struct {
  Data json.RawMessage `json:"data"`
}
