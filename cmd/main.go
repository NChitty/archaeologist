package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/NChitty/artifactsmmo/pkg/clients"
	"github.com/NChitty/artifactsmmo/pkg/schemas/responses"
)

func main() {
	client, _ := clients.NewClient("https://api.artifactsmmo.com/")
	httpResponse, _ := client.GetStatus(context.TODO())

	defer httpResponse.Body.Close()

	var responseContainer responses.ResponseContainer

	body, _ := io.ReadAll(httpResponse.Body)
	json.Unmarshal(body, &responseContainer)

	var statusResponse responses.StatusSchema
	json.Unmarshal(responseContainer.Data, &statusResponse)

	fmt.Printf("Schema: %+v\n", statusResponse)
}
