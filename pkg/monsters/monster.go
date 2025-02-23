package monsters

import (
	"context"
	"errors"
	"log/slog"
	"time"

	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
)

type MonsterAccessor interface {
	GetAllMonsters(params *artifactsmmo.GetAllMonstersMonstersGetParams) (*[]artifactsmmo.MonsterSchema, error)
	GetMonster(code string) (*artifactsmmo.MonsterSchema, error)
}

type ClientMonsterAccessor struct {
	client *artifactsmmo.ClientWithResponses
	logger *slog.Logger
}

func NewClientMonsterAccessor(client *artifactsmmo.ClientWithResponses, logger *slog.Logger) *ClientMonsterAccessor {
	return &ClientMonsterAccessor{client, logger}
}

func (accessor *ClientMonsterAccessor) GetAllMonsters(
	params *artifactsmmo.GetAllMonstersMonstersGetParams,
) (*[]artifactsmmo.MonsterSchema, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	respChan := make(chan struct {
		*artifactsmmo.GetAllMonstersMonstersGetResponse
		error
	})
	go func(retVal chan struct {
		*artifactsmmo.GetAllMonstersMonstersGetResponse
		error
	}) {
		resp, err := accessor.client.GetAllMonstersMonstersGetWithResponse(ctx, params)
		retVal <- struct {
			*artifactsmmo.GetAllMonstersMonstersGetResponse
			error
		}{resp, err}
	}(respChan)

	accessor.logger.Info("Searching for monsters", "params", *params)

	resp := <-respChan

	if resp.error != nil {
		return nil, resp.error
	}
	if resp.StatusCode() != 200 {
		accessor.logger.Error("Received non-200 status code", "code", resp.StatusCode(), "body", string(resp.Body))
		return nil, errors.New(string(resp.Body))
	}

	accessor.logger.Info("Found monsters", "monsters", resp.JSON200.Data)

	return &resp.JSON200.Data, nil
}

func (accessor *ClientMonsterAccessor) GetMonster(code string) (*artifactsmmo.MonsterSchema, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	respChan := make(chan struct {
		*artifactsmmo.GetMonsterMonstersCodeGetResponse
		error
	})
	go func(retVal chan struct {
		*artifactsmmo.GetMonsterMonstersCodeGetResponse
		error
	}) {
		resp, err := accessor.client.GetMonsterMonstersCodeGetWithResponse(ctx, code)
		retVal <- struct {
			*artifactsmmo.GetMonsterMonstersCodeGetResponse
			error
		}{resp, err}
	}(respChan)

	accessor.logger.Info("Searching for monster", "code", code)

	resp := <-respChan

	if resp.error != nil {
		return nil, resp.error
	}
	if resp.StatusCode() != 200 {
		accessor.logger.Error("Received non-200 status code", "code", resp.StatusCode(), "body", string(resp.Body))
		return nil, errors.New(string(resp.Body))
	}

	accessor.logger.Info("Found monster", "monster", resp.JSON200.Data)

	return &resp.JSON200.Data, nil
}
