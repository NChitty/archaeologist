package clients

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/NChitty/archaeologist/pkg/schemas"
)

type MapClient interface {
	GetAllMaps(
		ctx context.Context,
		reqEditors ...RequestEditorFn,
	) (*schemas.PagedResponseContainer[schemas.MapSchema], error)
}

func (c *Client) GetAllMaps(
	ctx context.Context,
	reqEditors ...RequestEditorFn,
) (*schemas.PagedResponseContainer[schemas.MapSchema], error) {
	req, err := NewGetAllMapsRequest(c.Server)
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

	var responseContainer *schemas.PagedResponseContainer[schemas.MapSchema]
	if err := json.Unmarshal(body, &responseContainer); err != nil {
		return nil, err
	}

	return responseContainer, nil
}

func NewGetAllMapsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/maps/")
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
