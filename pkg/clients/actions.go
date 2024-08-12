package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/NChitty/artifactsmmo/pkg/schemas"
)

type ActionClient interface {
	PostCharacterMoveAction(
		ctx context.Context,
		request schemas.PositionSchema,
		reqEditor ...RequestEditorFn,
	) (*schemas.CharacterMovementDataSchema, error)
}

func (c *Client) PostCharacterMoveAction(
	ctx context.Context,
	character string,
	request schemas.PositionSchema,
	reqEditors ...RequestEditorFn,
) (*schemas.CharacterMovementDataSchema, error) {
	req, err := NewPostMyCharacterActionMoveRequest(c.Server, character, request)
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

	var characterMovementDataResponse *schemas.CharacterMovementDataSchema
	if err := json.Unmarshal(responseContainer.Data, &characterMovementDataResponse); err != nil {
		return nil, err
	}

	return characterMovementDataResponse, nil
}

func NewPostMyCharacterActionMoveRequest(
	server string,
	character string,
	request schemas.PositionSchema,
) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/my/%s/action/move", character)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	payload, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}

	return req, nil
}
