package monsters

import (
	"context"
	"errors"
	"log/slog"
	"time"

	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
)

type ClientMonsterAccessor struct {
	client artifactsmmo.ClientWithResponsesInterface
	logger *slog.Logger
}

func NewClientMonsterAccessor(client artifactsmmo.ClientWithResponsesInterface, logger *slog.Logger) *ClientMonsterAccessor {
	return &ClientMonsterAccessor{client, logger}
}

func orDefault[T any](ptr *T, defaultValue T) T {
	if ptr != nil {
		return *ptr
	}
	return defaultValue
}

func (accessor *ClientMonsterAccessor) GetAllMonsters(
	minLevel *int,
	maxLevel *int,
	drop *string,
	page *int,
	size *int,
) (*[]artifactsmmo.MonsterSchema, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	accessor.logger.Info("Searching for monsters",
		"minLevel",
		orDefault(minLevel, 1),
		"maxLevel",
		orDefault(maxLevel, 40),
		"drop",
		orDefault(drop, ""),
		"page",
		orDefault(page, 1),
		"size",
		orDefault(size, 50))

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

func (accessor *ClientMonsterAccessor) GetMonster(code string) (*artifactsmmo.MonsterSchema, error) {
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
