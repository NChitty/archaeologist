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

type ServerClient interface {
	GetStatus(ctx context.Context, reqEditors ...RequestEditorFn) (*schemas.StatusSchema, error)
}

func (c *Client) GetStatus(
	ctx context.Context,
	reqEditors ...RequestEditorFn,
) (*schemas.StatusSchema, error) {
	req, err := NewGetStatusRequest(c.Server)
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

	var responseContainer *schemas.ResponseContainer
	if err := json.Unmarshal(body, &responseContainer); err != nil {
		return nil, err
	}

	var statusResponse *schemas.StatusSchema
	if err := json.Unmarshal(responseContainer.Data, &statusResponse); err != nil {
		return nil, err
	}

	return statusResponse, nil
}

func NewGetStatusRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/")
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
