package artifacts

import (
	"context"
	"log/slog"
	"time"

	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
)

func GetItem(name string) (*artifactsmmo.ItemSchema, error) {
	client, err := artifactsmmo.NewClientWithResponses("https://api.artifactsmmo.com/")
	if err != nil {
		slog.Error("Could not create authentication-less client:", err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	itemResp, err := client.GetItemItemsCodeGetWithResponse(ctx, name)
	if err != nil {
		slog.Error("Could not retrieve item:", err)
		return nil, err
	}

	return &itemResp.JSON200.Data, nil
}

