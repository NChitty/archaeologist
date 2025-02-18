package items_test

import (
	"bytes"
	_ "embed"
	"errors"
	"io"
	"log/slog"
	"net/http"
	"testing"

	"github.com/NChitty/archaeologist/pkg/items"
	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
	"github.com/stretchr/testify/assert"
)

var mockItemService *items.ItemService

//go:embed get_item_dne.json
var dne string

//go:embed get_item_copper_ore.json
var copperOre string

//go:embed get_item_copper.json
var copper string

type mockResponse struct {
	status     string
	statusCode int
	response   string
}

type mockHttpClient struct {
	responses map[string]mockResponse
}

func (client mockHttpClient) Do(req *http.Request) (*http.Response, error) {
	path := req.URL.Path
	mockResponse := client.responses[path]
	return &http.Response{
		Status:     mockResponse.status,
		StatusCode: mockResponse.statusCode,
		Body:       io.NopCloser(bytes.NewBufferString(mockResponse.response)),
		Header:     map[string][]string{"Content-Type": []string{"application/json"}},
	}, nil
}

func TestGetItem(t *testing.T) {
	mockItemService, err := items.NewItemService(
		slog.Default(),
		"none",
		artifactsmmo.WithHTTPClient(mockHttpClient{
			responses: map[string]mockResponse{
				"/items/dne": mockResponse{
					status:     "404 Not Found",
					statusCode: 404,
					response:   dne,
				},
				"/items/copper_ore": mockResponse{
					status:     "200 OK",
					statusCode: 200,
					response:   copperOre,
				},
				"/items/copper": mockResponse{
					status:     "200 OK",
					statusCode: 200,
					response:   copper,
				},
			},
		}),
	)
	assert.NoError(t, err)

	type result struct {
		okResult  *artifactsmmo.ItemSchema
		errResult error
	}

	type testCase struct {
		item     string
		expected result
	}
	testCases := []testCase{
		testCase{
			item:     "dne",
			expected: result{nil, errors.New("Item not found.")},
		},
		testCase{
			item: "copper_ore",
			expected: result{&artifactsmmo.ItemSchema{
				Name:        "Copper Ore",
				Code:        "copper_ore",
				Level:       1,
				Type:        "resource",
				Subtype:     "mining",
				Description: "",
				Effects:     &[]artifactsmmo.ItemEffectSchema{},
				Craft:       nil,
				Tradeable:   true,
			}, nil},
		},
	}
	t.Run("Get Item", func(t *testing.T) {
		for _, testCase := range testCases {
			item, err := mockItemService.GetItem(testCase.item)
			assert.Equal(t, testCase.expected.okResult, item)
			assert.Equal(t, testCase.expected.errResult, err)
		}
	})
}
