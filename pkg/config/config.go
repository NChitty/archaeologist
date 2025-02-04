package config

import (
	"errors"
	"time"
)

type Config struct {
	timeout time.Duration
	token   string
	url     string
}

func (config *Config) GetTimeout() time.Duration {
  return config.timeout
}

func (config *Config) GetToken() string {
  return config.token
}

func (config *Config) GetUrl() string {
  return config.url
}

type ConfigProviderFn func() (*Config, error)

type ConfigProvider interface {
	GetTimeout() (time.Duration, error)
	GetToken() (string, error)
	GetUrl() (string, error)
}

type ConfigProviderChain struct {
	providers []ConfigProviderFn
}

func (providerChain *ConfigProviderChain) GetConfig() (*Config, error) {
	var err error
	for _, providerFn := range providerChain.providers {
		config, provderErr := providerFn()
		if err != nil {
			err = errors.Join(err, provderErr)
			continue
		}
		return config, nil
	}
	return nil, err
}
