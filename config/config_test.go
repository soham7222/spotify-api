package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	config, err := Load("test_config.json")
	assert.Nil(t, err)

	assert.Equal(t, "mock_client_id", config.GetClientId())
	assert.Equal(t, "mock_client_secret", config.GetClientSecretKey())
	assert.Equal(t, "https://accounts.spotify.com/api/token", config.GetTokenIssuerUrl())
	assert.Equal(t, "https://api.spotify.com/v1/search?q=isrc:%s&type=track", config.GetSpotifySearchApi())
}
