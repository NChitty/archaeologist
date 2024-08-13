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
	"github.com/NChitty/artifactsmmo/pkg/schemas/actions"
)

type ActionClient interface {
	PostCharacterMove(
		ctx context.Context,
		character string,
		request actions.PositionSchema,
		reqEditor ...RequestEditorFn,
	) (*actions.CharacterMovementDataSchema, error)

	PostCharacterEquipItem(
		ctx context.Context,
		character string,
		request actions.EquipItemRequest,
		reqEditor ...RequestEditorFn,
	) (*actions.EquipItemResponseSchema, error)

	PostCharacterUnequipItem(
		ctx context.Context,
		character string,
		request actions.UnequipItemRequest,
		reqEditor ...RequestEditorFn,
	) (*actions.EquipItemResponseSchema, error)

	PostCharacterFight(
		ctx context.Context,
		character string,
		reqEditor ...RequestEditorFn,
	) (*actions.CharacterFightDataSchema, error)

	PostCharacterGathering(
		ctx context.Context,
		character string,
		reqEditor ...RequestEditorFn,
	) (*actions.SkillDataSchema, error)

	PostCharacterCrafting(
		ctx context.Context,
		character string,
		request actions.CraftingRequestSchema,
		reqEditor ...RequestEditorFn,
	) (*actions.SkillDataSchema, error)

	PostCharacterBankDepositItem(
		ctx context.Context,
		character string,
		request actions.BankItemDepositRequestSchema,
		reqEditor ...RequestEditorFn,
	) (*actions.BankItemSchema, error)

	PostCharacterBankWithdrawItem(
		ctx context.Context,
		character string,
		request actions.BankItemWithdrawRequestSchema,
		reqEditor ...RequestEditorFn,
	) (*actions.BankItemSchema, error)

	PostCharacterBankDepositGold(
		ctx context.Context,
		character string,
		request actions.BankGoldDepositRequestSchema,
		reqEditor ...RequestEditorFn,
	) (*actions.GoldTransactionSchema, error)

	PostCharacterBankWithdrawGold(
		ctx context.Context,
		character string,
		request actions.BankGoldWithdrawRequestSchema,
		reqEditor ...RequestEditorFn,
	) (*actions.GoldTransactionSchema, error)
}

func (c *Client) PostCharacterMove(
	ctx context.Context,
	character string,
	request actions.PositionSchema,
	reqEditors ...RequestEditorFn,
) (*actions.CharacterMovementDataSchema, error) {
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

	var characterMovementDataResponse *actions.CharacterMovementDataSchema
	if err := json.Unmarshal(responseContainer.Data, &characterMovementDataResponse); err != nil {
		return nil, err
	}

	return characterMovementDataResponse, nil
}

func NewPostMyCharacterActionMoveRequest(
	server string,
	character string,
	request actions.PositionSchema,
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
