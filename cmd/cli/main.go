package main

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"strconv"
	"time"

	"github.com/NChitty/archaeologist/cmd/cli/token"
	"github.com/NChitty/archaeologist/pkg/accounts"
	"github.com/NChitty/archaeologist/pkg/actors"
	"github.com/NChitty/archaeologist/pkg/characters"
	"github.com/NChitty/archaeologist/pkg/fights"
	"github.com/NChitty/archaeologist/pkg/items"
	"github.com/NChitty/archaeologist/pkg/items/effects"
	"github.com/NChitty/archaeologist/pkg/maps"
	"github.com/NChitty/archaeologist/pkg/monsters"
	"github.com/phsym/console-slog"
	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
)

var characterService *characters.CharacterService
var itemService items.ItemAccessor
var effectAccumulator *effects.EffectAccumulator
var fightService *fights.FightService
var mapAccessor maps.MapAccessor
var monsterAccessor monsters.MonsterAccessor

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

	account, err := accounts.New(client)
	if err != nil {
		os.Exit(1)
	}

	accountCharacters, err := account.GetCharacters(token)
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

	selectedCharacter := characters.FromSchema(accountCharacters[choice-1])
	selectedCharacter.Token = token
	characterService := characters.NewCharacterService(logger, client)
	itemService = items.DefaultItemService()
	effectAccumulator = effects.New(slog.Default())
	fightService = fights.NewFightService(effectAccumulator, slog.Default(), itemService)
	mapAccessor = maps.NewClientMapAccessor(client, slog.Default())
	monsterAccessor = monsters.NewClientMonsterAccessor(client, slog.Default())

	var actor actors.Actor
	actorQueue := make(chan struct {
		actors.Actor
		*characters.CharacterWrapper
	}, 5)
	defer close(actorQueue)
	go actorsDo(logger, actorQueue)

	for {
		fmt.Println("[1] Gather")
		fmt.Println("[2] Craft")
		fmt.Println("[3] Task")
		_, err = fmt.Scanf("%d\n", &choice)
		if err != nil {
			logger.Error("Did not understand the choice", "error", err)
			os.Exit(1)
		}
		switch choice {
		case 1:
			code, qty := ItemInput(logger, "gather", "gather")
			actor, err = actors.NewGatherActor(code, qty, characterService, itemService, mapAccessor, logger)
			if err != nil {
				fmt.Errorf("Error: %w", err)
				continue
			}
			actorQueue <- struct {
				actors.Actor
				*characters.CharacterWrapper
			}{actor, selectedCharacter}
		case 2:
			code, qty := ItemInput(logger, "craft", "craft")
			actor, err = actors.NewCraftingActor(code, qty, characterService, itemService, mapAccessor, logger)
			if err != nil {
				fmt.Errorf("Error: %w", err)
				continue
			}
			actorQueue <- struct {
				actors.Actor
				*characters.CharacterWrapper
			}{actor, selectedCharacter}
		case 3:
			var input string
			fmt.Printf("Would you like to quit fighting when you are out of healing items? (1/t/true): ")
			_, err := fmt.Scanf("%s\n", &input)
			if err != nil {
				logger.Error("Did not understand the input", "error", err)
				continue
			}
			exitOnRest, err := strconv.ParseBool(input)
			if err != nil {
				logger.Error("Did not understand the input", "error", err)
				continue
			}
			actor, err := actors.NewTaskFightingActor(selectedCharacter, exitOnRest, characterService, fightService, itemService, mapAccessor, monsterAccessor, logger)
			if err != nil {
				fmt.Printf("An error occurred: %v\n", err)
				continue
			}
			fmt.Println("Created fighting actor")
			actorQueue <- struct {
				actors.Actor
				*characters.CharacterWrapper
			}{actor, selectedCharacter}
		default:
			break
		}
	}
}

func ItemInput(logger *slog.Logger, itemPrompt string, quantityPrompt string) (string, int) {
	var code string
	fmt.Printf("Type code of item you would like to %s: ", itemPrompt)
	_, err := fmt.Scanf("%s\n", &code)
	if err != nil {
		logger.Error("Did not understand the input", "error", err)
	}

	var quantity int
	fmt.Printf("Type code of item you would like to %s: ", quantityPrompt)
	_, err = fmt.Scanf("%d\n", &quantity)
	if err != nil {
		logger.Error("Did not understand the input", "error", err)
	}
	return code, quantity
}

func actorsDo(logger *slog.Logger, actor chan struct {
	actors.Actor
	*characters.CharacterWrapper
}) {
	logger.Debug("Waiting for actor")
	pop := <-actor
	logger.Debug("Acquired actor", "actor", actor)
	for {
		err := pop.Do(pop.CharacterWrapper)
		if err != nil {
			logger.Error("Failed to perform action", "error", err)
		}
		logger.Debug("Waiting for actor")
		pop = <-actor
	}
}
