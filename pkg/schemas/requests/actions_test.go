package requests_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/NChitty/artifactsmmo/pkg/schemas/requests"
)

func TestPositionMarshalling(t *testing.T) {
  pos := requests.PositionSchema{X: 0, Y: 1}
  expected := []byte(`{"x":0,"y":1}`)
  actual, err := json.Marshal(pos)
  if !bytes.Equal(expected, actual) || err != nil {
    t.Fatalf(`json.Marshal(pos) = %q, %v`, actual, err)
  }
}
