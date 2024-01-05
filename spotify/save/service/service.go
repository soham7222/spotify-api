package service

import (
	"encoding/json"
	"fmt"
	"spotify-api/config"
	"spotify-api/error"
	"spotify-api/spotify/save/model"

	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
)

type Service interface {
	FetchFromSpotify(context *gin.Context, isrc string) (model.TracksSearchResponse, *error.ErrorResponse)
}

type service struct {
	request *gorequest.SuperAgent
	config  config.Config
}

func NewService(request *gorequest.SuperAgent, config config.Config) Service {
	return &service{
		request: request,
		config:  config,
	}
}

func (service service) FetchFromSpotify(context *gin.Context, isrc string) (model.TracksSearchResponse, *error.ErrorResponse) {
	service.request.Get(fmt.Sprintf(service.config.GetSpotifySearchApi(), isrc))
	bearerToken := "Bearer " + context.Request.Header.Get("Authorization")
	service.request.Set("Authorization", bearerToken)

	_, body, err := service.request.End()
	if err != nil {
		fmt.Errorf("unable to fetch data. error: %v", err)
		return model.TracksSearchResponse{}, error.SpotyfyErrors[error.InternalServerError]
	}

	var response model.TracksSearchResponse
	marshalErr := json.Unmarshal([]byte(body), &response)
	if marshalErr != nil {
		fmt.Errorf("unable to un marshal . error: %v", marshalErr)
		return model.TracksSearchResponse{}, error.SpotyfyErrors[error.BadFormattedJSONError]
	}

	return response, nil
}
