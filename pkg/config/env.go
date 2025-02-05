package config

import (
	"errors"
	"log/slog"
	"os"
	"time"
)

type EnvironmentConfigProvider struct {
}

func (provider *EnvironmentConfigProvider) ConfigureLogger() (*slog.Logger, error) {
	return slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{AddSource: true})), nil
}

func (provider *EnvironmentConfigProvider) ConfigureTimeout() (time.Duration, error) {
	timeoutStr := os.Getenv("ARTIFACTSMMO_TIMEOUT_DURATION")

	if len(timeoutStr) == 0 {
		return 30 * time.Second, errors.New("ARTIFACTSMMO_TIMEOUT_DURATION environment variable is blank")
	}

	timeout, err := time.ParseDuration(timeoutStr)
	if err != nil {
		return 30 * time.Second, err
	}

	return timeout, nil
}

func (provider *EnvironmentConfigProvider) ConfigureToken() (string, error) {
	token := os.Getenv("ARTIFACTSMMO_TOKEN")

	if len(token) == 0 {
		return "", errors.New("ARTIFACTSMMO_TOKEN environment variable is blank")
	}

	return token, nil
}

func (provider *EnvironmentConfigProvider) ConfigureUrl() (string, error) {
	url := os.Getenv("ARTIFACTSMMO_URL")

	if len(url) == 0 {
		return "https://api.artifactsmmo.com/", errors.New("ARTIFACTSMMO_URL environment variable is blank")
	}

	return url, nil
}
