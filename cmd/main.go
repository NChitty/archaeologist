package main

import (
	"github.com/NChitty/artifactsmmo/pkg/clients"
)

func main() {
	_, _ = clients.NewClient("https://api.artifactsmmo.com/")
}
