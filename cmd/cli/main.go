package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/NChitty/archaeologist/cmd/cli/token"
	"github.com/NChitty/archaeologist/pkg/accounts"
	"github.com/NChitty/archaeologist/pkg/actors"
	"github.com/NChitty/archaeologist/pkg/characters"
	"github.com/NChitty/archaeologist/pkg/fights"
	"github.com/NChitty/archaeologist/pkg/items"
	"github.com/NChitty/archaeologist/pkg/items/effects"
	"github.com/phsym/console-slog"
	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
)

var itemService *items.ItemService
var effectAccumulator *effects.EffectAccumulator
var fightService *fights.FightService

func main() {
	logFile, err := os.Create(fmt.Sprintf("%s.log", time.Now().Format("2006-01-02_15-04")))
	if err != nil {
		log.Fatal("Could not create log file")
	}

	logger := slog.New(console.NewHandler(logFile, &console.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	client, err := artifactsmmo.NewClientWithResponses("https://api.artifactsmmo.com/")
	if err != nil {
		logger.Error("Could not create new web client for artifacts mmo", "error", err)
		os.Exit(1)
	}

	token := token.GetToken(client)

	client, err = artifactsmmo.NewClientWithResponses(
		"https://api.artifactsmmo.com/",
		artifactsmmo.WithRequestEditorFn(artifactsmmo.NewBearerAuthorizationRequestFunc(token)),
	)
	if err != nil {
		logger.Error("Could not create new web client for artifacts mmo", "error", err)
		os.Exit(1)
	}

	account, err := accounts.New(client)
	if err != nil {
		os.Exit(1)
	}

	accountCharacters, err := account.GetCharacters()
	if err != nil {
		logger.Error("Could not retrieve characters for given account token", "error", err)
		os.Exit(1)
	}

	for i, char := range accountCharacters {
		fmt.Printf("[%1d] Name: %-20s\tLevel: %3d\n", i+1, char.Name, char.Level)
	}

	var choice int
	fmt.Println("Pick your character from the list.")
	_, err = fmt.Scanf("%d\n", &choice)
	if err != nil {
		logger.Error("Did not understand the choice", "error", err)
		os.Exit(1)
	}

	if choice <= 0 || choice > len(accountCharacters) {
		logger.Error("Invalid choice")
		os.Exit(1)
	}

	selectedCharacter, err := characters.New(logger, client, accountCharacters[choice-1].Name)
	if err != nil {
		logger.Error("Could not select character", "error", err)
		os.Exit(1)
	}

	itemService = items.DefaultItemService()
	effectAccumulator = effects.New(slog.Default())
	fightService = fights.NewFightService(effectAccumulator, slog.Default(), itemService)

	var actor actors.Actor
	actorQueue := make(chan actors.Actor, 5)
	defer close(actorQueue)
	go actorsDo(logger, actorQueue)

	for {
		fmt.Println("[1] Gather")
		fmt.Println("[2] Craft")
		_, err = fmt.Scanf("%d\n", &choice)
		if err != nil {
			logger.Error("Did not understand the choice", "error", err)
			os.Exit(1)
		}
		switch choice {
		case 1:
			code, qty := ItemInput(logger)
			actor, err = actors.NewGatherActor(code, qty, selectedCharacter, itemService, logger)
			if err != nil {
				continue
			}
			actorQueue <- actor
		case 2:
			code, qty := ItemInput(logger)
			actor, err = actors.NewCraftingActor(code, qty, selectedCharacter, itemService, logger)
			if err != nil {
				continue
			}
			actorQueue <- actor
		default:
			break
		}
	}
}

func ItemInput(logger *slog.Logger) (string, int) {
	var code string
	fmt.Print("Type code of item you would like to gather: ")
	_, err := fmt.Scanf("%s\n", &code)
	if err != nil {
		logger.Error("Did not understand the input", "error", err)
	}

	var quantity int
	fmt.Print("Type amount of the item you would like to gather: ")
	_, err = fmt.Scanf("%d\n", &quantity)
	if err != nil {
		logger.Error("Did not understand the input", "error", err)
	}
	return code, quantity
}

func actorsDo(logger *slog.Logger, actor chan actors.Actor) {
	logger.Debug("Waiting for actor")
	pop := <-actor
	logger.Debug("Acquired actor", "actor", actor)
	for {
		err := pop.Do()
		if err != nil {
			logger.Error("Failed to perform action", "error", err)
		}
		logger.Debug("Waiting for actor")
		pop = <-actor
	}
}
