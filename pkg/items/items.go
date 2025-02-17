package items

import (
	"context"
	"errors"
	"log/slog"
	"time"

	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
)

func GetItem(logger *slog.Logger, name string) (*artifactsmmo.ItemSchema, error) {
	client, err := artifactsmmo.NewClientWithResponses("https://api.artifactsmmo.com/")
	if err != nil {
		logger.Error("Could not create authentication-less client", "error", err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	itemResp, err := client.GetItemItemsCodeGetWithResponse(ctx, name)
	if err != nil {
		logger.Error("Could not retrieve item", "error", err)
		return nil, err
	}

	return &itemResp.JSON200.Data, nil
}

func GetAllResources(logger *slog.Logger, skill *artifactsmmo.GatheringSkill, code *string) ([]artifactsmmo.ResourceSchema, error) {
	client, err := artifactsmmo.NewClientWithResponses("https://api.artifactsmmo.com/")
	if err != nil {
		logger.Error("Could not create authentication-less client", "error", err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	resourceResp, err := client.GetAllResourcesResourcesGetWithResponse(
		ctx,
		&artifactsmmo.GetAllResourcesResourcesGetParams{
			MinLevel: nil,
			MaxLevel: nil,
			Skill:    skill,
			Drop:     code,
			Page:     nil,
			Size:     nil,
		},
	)
	if err != nil {
		logger.Error("Could not retrieve gather resource", "error", err)
		return nil, err
	}
	if results, err := resourceResp.JSON200.Total.AsDataPageResourceSchemaTotal0(); err != nil {
		return nil, err
	} else if results < 1 {
		return nil, errors.New("No results.")
	}

	return resourceResp.JSON200.Data, nil
}

type Equipment struct {
  Item artifactsmmo.ItemSchema
  Quantity int
}
