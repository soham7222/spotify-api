package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"spotify-api/client/model"
	"spotify-api/config"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=spotify_client.go -destination=../mocks/mock_spotify_client.go -package=mocks
type SpotifyClient interface {
	FetchTrackDetailsBasedOnISRC(ctx *gin.Context, isrc string) (model.TracksSearchResponse, error)
}

type spotifyClient struct {
	client *http.Client
	config config.Config
}

func NewSpotifyClient(client *http.Client, config config.Config) SpotifyClient {
	return spotifyClient{
		client: client,
		config: config,
	}
}

func (s spotifyClient) FetchTrackDetailsBasedOnISRC(ctx *gin.Context, isrc string) (model.TracksSearchResponse, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(s.config.GetSpotifySearchApi(), isrc), nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return model.TracksSearchResponse{}, err
	}

	bearerToken := "Bearer " + ctx.Request.Header.Get("Authorization")
	req.Header.Set("Authorization", bearerToken)

	resp, err := s.client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return model.TracksSearchResponse{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return model.TracksSearchResponse{}, err
	}

	var res model.TracksSearchResponse
	marshalErr := json.Unmarshal([]byte(body), &res)
	if marshalErr != nil {
		fmt.Errorf("unable to un marshal . error: %v", marshalErr)
		return model.TracksSearchResponse{}, err
	}

	return res, nil
}
