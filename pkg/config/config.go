package config

import (
	"errors"
	"log/slog"
	"os"
	"time"
)

type Config struct {
	logger  *slog.Logger
	timeout time.Duration
	token   string
	url     string
}

func (config *Config) GetLogger() (*slog.Logger) {
  return config.logger
}

func (config *Config) GetTimeout() (time.Duration) {
  return config.timeout
}

func (config *Config) GetToken() (string) {
  return config.token
}

func (config *Config) GetUrl() (string) {
  return config.url
}

type LoggerProvider interface {
	ConfigureLogger() (*slog.Logger, error)
}

type TimeoutProvider interface {
	ConfigureTimeout() (time.Duration, error)
}

type TokenProvider interface {
	ConfigureToken() (string, error)
}

type UrlProvider interface {
	ConfigureUrl() (string, error)
}

type ConfigProviderChain struct {
	loggerProviderChain  []LoggerProvider
	timeoutProviderChain []TimeoutProvider
	tokenProviderChain   []TokenProvider
	urlProviderChain     []UrlProvider
}

func New() *ConfigProviderChain {
	return &ConfigProviderChain{}
}

func Default() *ConfigProviderChain {
	defaultProvider := DefaultConfigProvider{}

	return &ConfigProviderChain{
		loggerProviderChain:  []LoggerProvider{&defaultProvider},
		timeoutProviderChain: []TimeoutProvider{&defaultProvider},
		tokenProviderChain:   []TokenProvider{&defaultProvider},
		urlProviderChain:     []UrlProvider{&defaultProvider},
	}
}

func (providerChain *ConfigProviderChain) GetConfig() (*Config, error) {
  var finalErr error

  var logger *slog.Logger
  for _, provider := range providerChain.loggerProviderChain {
    provided, err := provider.ConfigureLogger()
    if err != nil {
      finalErr = errors.Join(err)
      continue
    }
    logger = provided
    finalErr = nil
  }

  var timeout time.Duration
  for _, provider := range providerChain.timeoutProviderChain {
    provided, err := provider.ConfigureTimeout()
    if err != nil {
      finalErr = errors.Join(err)
      continue
    }
    timeout = provided
    finalErr = nil
  }

  var token string
  for _, provider := range providerChain.tokenProviderChain {
    provided, err := provider.ConfigureToken()
    if err != nil {
      finalErr = errors.Join(err)
      continue
    }
    token = provided
    finalErr = nil
  }

  var url string
  for _, provider := range providerChain.urlProviderChain {
    provided, err := provider.ConfigureUrl()
    if err != nil {
      finalErr = errors.Join(err)
      continue
    }
    url = provided
    finalErr = nil
  }

  if finalErr != nil {
    return nil, finalErr
  }

  return &Config{
    logger: logger,
    timeout: timeout,
    token: token,
    url: url,
  }, nil
}

type DefaultConfigProvider struct {
}

func (provider *DefaultConfigProvider) ConfigureLogger() (*slog.Logger, error) {
	return slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{AddSource: true})), nil
}

func (provider *DefaultConfigProvider) ConfigureTimeout() (time.Duration, error) {
	return 30 * time.Second, nil
}

func (provider *DefaultConfigProvider) ConfigureToken() (string, error) {
	return "", errors.New("Must have a token configured.")
}

func (provider *DefaultConfigProvider) ConfigureUrl() (string, error) {
	return "https://api.artifactsmmo.com/", nil
}
