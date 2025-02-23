package maps

import (
	"context"
	"log/slog"
	"time"

	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
)

type MapAccessor interface {
	GetAllMaps(contentType *string, contentCode *string, page *int, size *int) ([]artifactsmmo.MapSchema, error)
	GetMap(x int, y int) (*artifactsmmo.MapSchema, error)
}

type clientMapAccessor struct {
	client artifactsmmo.ClientWithResponsesInterface
	logger *slog.Logger
}

func NewClientMapAccessor(client artifactsmmo.ClientWithResponsesInterface, logger *slog.Logger) MapAccessor {
	return &clientMapAccessor{
		client: client,
		logger: logger,
	}
}

func (accessor *clientMapAccessor) GetAllMaps(
	contentType *string,
	contentCode *string,
	page *int,
	size *int,
) ([]artifactsmmo.MapSchema, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	mapResp, err := accessor.client.GetAllMapsMapsGetWithResponse(
		ctx,
		&artifactsmmo.GetAllMapsMapsGetParams{
			ContentType: (*artifactsmmo.MapContentTypeAZAZ09)(contentType),
			ContentCode: contentCode,
			Page:        page,
			Size:        size,
		},
	)

	if err != nil {
		accessor.logger.Error("Could not retrieve item", "error", err)
		return nil, err
	}

	return mapResp.JSON200.Data, nil
}

func (accessor *clientMapAccessor) GetMap(x int, y int) (*artifactsmmo.MapSchema, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	accessor.logger.Info("")

	mapResp, err := accessor.client.GetMapMapsXYGetWithResponse(ctx, x, y)

	if err != nil {
		accessor.logger.Error("Could not retrieve item", "error", err)
		return nil, err
	}

	return &mapResp.JSON200.Data, nil
}
