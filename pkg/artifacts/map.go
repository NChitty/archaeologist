package artifacts

import (
	"context"
	"log/slog"
	"time"

	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
)

func GetAllMaps(logger *slog.Logger, contentType *string, contentCode *string) ([]artifactsmmo.MapSchema, error) {
	client, err := artifactsmmo.NewClientWithResponses("https://api.artifactsmmo.com/")
	if err != nil {
		logger.Error("Could not create authentication-less client:", err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	mapResp, err := client.GetAllMapsMapsGetWithResponse(
		ctx,
		&artifactsmmo.GetAllMapsMapsGetParams{
			ContentCode: contentCode,
			ContentType: (*artifactsmmo.MapContentType)(contentType),
			Page:        nil,
			Size:        nil,
		})
	if err != nil {
		logger.Error("Could not retrieve item:", err)
		return nil, err
	}
	logger.Debug("Map response", "map", mapResp.Body)

	return mapResp.JSON200.Data, nil
}
