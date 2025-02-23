package monsters

import (
	"context"
	"errors"
	"log/slog"
	"time"

	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
)

type MonsterAccessor interface {
	GetAllMonsters(
		minLevel *int,
		maxLevel *int,
		drop *string,
		page *int,
		size *int,
	) (*[]artifactsmmo.MonsterSchema, error)
	GetMonster(code string) (*artifactsmmo.MonsterSchema, error)
}

type clientMonsterAccessor struct {
	client *artifactsmmo.ClientWithResponses
	logger *slog.Logger
}

func NewClientMonsterAccessor(client *artifactsmmo.ClientWithResponses, logger *slog.Logger) MonsterAccessor {
	return &clientMonsterAccessor{client, logger}
}

func (accessor *clientMonsterAccessor) GetAllMonsters(
	minLevel *int,
	maxLevel *int,
	drop *string,
	page *int,
	size *int,
) (*[]artifactsmmo.MonsterSchema, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	accessor.logger.Info("Searching for monsters", "minLevel", *minLevel, "maxLevel", *maxLevel, "drop", *drop, "page", *page, "size", *size)

	resp, err := accessor.client.GetAllMonstersMonstersGetWithResponse(
		ctx,
		&artifactsmmo.GetAllMonstersMonstersGetParams{
			MinLevel: minLevel,
			MaxLevel: maxLevel,
			Drop:     drop,
			Page:     page,
			Size:     size,
		})
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		accessor.logger.Error("Received non-200 status code", "code", resp.StatusCode(), "body", string(resp.Body))
		return nil, errors.New(string(resp.Body))
	}

	accessor.logger.Info("Found monsters", "monsters", resp.JSON200.Data)

	return &resp.JSON200.Data, nil
}

func (accessor *clientMonsterAccessor) GetMonster(code string) (*artifactsmmo.MonsterSchema, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

  accessor.logger.Info("Searching for monster", "code", code)

	resp, err := accessor.client.GetMonsterMonstersCodeGetWithResponse(ctx, code)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != 200 {
		accessor.logger.Error("Received non-200 status code", "code", resp.StatusCode(), "body", string(resp.Body))
		return nil, errors.New(string(resp.Body))
	}

	accessor.logger.Info("Found monster", "monster", resp.JSON200.Data)

	return &resp.JSON200.Data, nil
}
