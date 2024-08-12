package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/NChitty/artifactsmmo/pkg/schemas"
)

type CharacterClient interface {
	GetAllCharacters(
		ctx context.Context,
		reqEditors ...RequestEditorFn,
	) (*schemas.PagedResponseContainer[schemas.CharacterSchema], error)
}

func (c *Client) GetAllCharacters(
	ctx context.Context,
	reqEditors ...RequestEditorFn,
) (*schemas.PagedResponseContainer[schemas.CharacterSchema], error) {
	req, err := NewGetAllCharactersRequest(c.Server)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}

	res, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var responseContainer *schemas.PagedResponseContainer[schemas.CharacterSchema]
	if err := json.Unmarshal(body, &responseContainer); err != nil {
		return nil, err
	}

	return responseContainer, nil
}

func NewGetAllCharactersRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/characters/")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}
