package config

import (
	"encoding/json"
	"log"
	"os"
)

//go:generate mockgen -source=config.go -destination=../mocks/mock_config.go -package=mocks
type Config interface {
	GetClientId() string
	GetClientSecretKey() string
	GetTokenIssuerUrl() string
	GetSpotifySearchApi() string
}

func NewConfig() Config {
	return &config{}
}

func (c *config) GetSpotifySearchApi() string {
	return c.SpotifySearchApi
}

func (c *config) GetClientId() string {
	return c.ClientId
}

func (c *config) GetClientSecretKey() string {
	return c.ClientSecret
}

func (c *config) GetTokenIssuerUrl() string {
	return c.TokenUrl
}

type config struct {
	ClientId         string `json:"client_id"`
	ClientSecret     string `json:"client_secret"`
	TokenUrl         string `json:"token_url"`
	SpotifySearchApi string `json:"spotify_search_api_url"`
}

func Load(configPath string) (Config, error) {
	var config config
	raw, err := os.ReadFile(configPath)
	if err != nil {
		log.Println("Error occured while reading config")
		return nil, err
	}
	_ = json.Unmarshal(raw, &config)

	return &config, nil
}
