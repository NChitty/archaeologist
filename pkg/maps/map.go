package maps

import (
	"context"
	"log/slog"
	"time"

	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
)

type ClientMapAccessor struct {
	client artifactsmmo.ClientWithResponsesInterface
	logger *slog.Logger
}

func NewClientMapAccessor(client artifactsmmo.ClientWithResponsesInterface, logger *slog.Logger) *ClientMapAccessor {
	return &ClientMapAccessor{
		client: client,
		logger: logger,
	}
}

func (accessor *ClientMapAccessor) GetAllMaps(
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

func (accessor *ClientMapAccessor) GetMap(x int, y int) (*artifactsmmo.MapSchema, error) {
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
