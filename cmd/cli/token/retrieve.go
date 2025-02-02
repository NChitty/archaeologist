package token

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"syscall"

	artifactsmmo "github.com/promiseofcake/artifactsmmo-go-client/client"
	"golang.org/x/term"
)

func GetToken(client *artifactsmmo.ClientWithResponses) string {
	token, isSet := os.LookupEnv("ARTIFACTSMMO_TOKEN")
	if isSet && len(token) > 0 {
		return token
	}

	slog.Debug("ARTIFACTSMMO_TOKEN environment variable is empty or unset, searching for user.token file.")
	if userTokenFile, err := os.Open("user.token"); errors.Is(err, os.ErrNotExist) {
		slog.Debug("No token file found, lets get logged in.")
		return login(client)
	} else {
		return readLoginFile(userTokenFile)
	}
}

func login(client *artifactsmmo.ClientWithResponses) string {
	fmt.Print("Email: ")
	var username string
	_, err := fmt.Scanf("%s\n", &username)
	if err != nil {
		slog.Error("Could not read username: ", err)
		os.Exit(1)
	}

	fmt.Print("Enter Password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		slog.Error("Could read password: ", err)
		os.Exit(1)
	}

	password := string(bytePassword)

	tokenResponse, err := client.GenerateTokenTokenPostWithResponse(
		context.TODO(),
		artifactsmmo.NewBasicAuthorizationRequestFunc(username, password),
	)
	if err != nil {
		slog.Error("Could not retrieve token response: ", err)
		os.Exit(1)
	}

	token := tokenResponse.JSON200.Token

	slog.Debug("Retrieved token, saving to \"user.token\"")

	file, err := os.Create("user.token")
	if err != nil {
		slog.Error("Could not create token file: ", err)
		return token
	}

	_, err = file.WriteString(token)
	if err != nil {
		slog.Error("Could not write token to file: ", err)
	}

	return token
}

func readLoginFile(file *os.File) string {
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text()
}
