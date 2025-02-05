package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/NChitty/archaeologist/cmd/cli/token"
	"github.com/NChitty/archaeologist/pkg/account"
	"github.com/NChitty/archaeologist/pkg/actors"
	"github.com/NChitty/archaeologist/pkg/character"
	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	client, err := artifactsmmo.NewClientWithResponses("https://api.artifactsmmo.com/")
	if err != nil {
		slog.Error("Could not create new web client for artifacts mmo:", err)
		os.Exit(1)
	}

	token := token.GetToken(client)

	client, err = artifactsmmo.NewClientWithResponses(
		"https://api.artifactsmmo.com/",
		artifactsmmo.WithRequestEditorFn(artifactsmmo.NewBearerAuthorizationRequestFunc(token)),
	)
	if err != nil {
		slog.Error("Could not create new web client for artifacts mmo: ", err)
		os.Exit(1)
	}

	account, err := account.New(client)
	if err != nil {
		os.Exit(1)
	}

	characters, err := account.GetCharacters()
	if err != nil {
		slog.Error("Could not retrieve characters for given account token: ", err)
		os.Exit(1)
	}

	for i, char := range characters {
		fmt.Printf("[%1d] Name: %-20s\tLevel: %3d\n", i+1, char.Name, char.Level)
	}

	var choice int
	fmt.Println("Pick your character from the list.")
	_, err = fmt.Scanf("%d\n", &choice)
	if err != nil {
		slog.Error("Did not understand the choice:", err)
		os.Exit(1)
	}

	if choice <= 0 || choice > len(characters) {
		slog.Error("Invalid choice")
		os.Exit(1)
	}

	selectedCharacter, err := character.New(client, characters[choice-1].Name)
	if err != nil {
		slog.Error("Could not select character:", err)
		os.Exit(1)
	}

	var actor actors.Actor
	actorQueue := make(chan actors.Actor, 5)
	defer close(actorQueue)
	go actorsDo(actorQueue)

	for {
		fmt.Println("[1] Gather")
		fmt.Println("[2] Craft")
		_, err = fmt.Scanf("%d\n", &choice)
		if err != nil {
			slog.Error("Did not understand the choice:", err)
			os.Exit(1)
		}
		switch choice {
		case 1:
			code, qty := ItemInput()
			actor, err = actors.NewGatherActor(selectedCharacter, code, qty)
			if err != nil {
				continue
			}
			actorQueue <- actor
		case 2:
			code, qty := ItemInput()
			actor, err = actors.NewCraftingActor(selectedCharacter, code, qty)
			if err != nil {
				continue
			}
			actorQueue <- actor
		default:
			break
		}
	}
}

func ItemInput() (string, int) {
	var code string
	fmt.Print("Type code of item you would like to gather: ")
	_, err := fmt.Scanf("%s\n", &code)
	if err != nil {
		slog.Error("Did not understand the input:", err)
		os.Exit(1)
	}

	var quantity int
	fmt.Print("Type amount of the item you would like to gather: ")
	_, err = fmt.Scanf("%d\n", &quantity)
	if err != nil {
		slog.Error("Did not understand the input:", err)
		os.Exit(1)
	}
	return code, quantity
}

func actorsDo(actor chan actors.Actor) {
	slog.Debug("Waiting for actor")
	pop := <-actor
	slog.Debug("Acquired actor", "actor", actor)
	for {
		err := pop.Do()
		if err != nil {
			slog.Error("Failed to perform action", err)
		}
		pop = <-actor
	}
}
