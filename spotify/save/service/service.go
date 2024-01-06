package service

import (
	"encoding/json"
	"fmt"
	"spotify-api/config"
	"spotify-api/error"
	"spotify-api/spotify/save/model"
	"spotify-api/spotify/save/model/response"
	"spotify-api/spotify/save/repository"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
)

type Service interface {
	FetchFromSpotifyAndInsertIntoDB(context *gin.Context, isrc string) (response.CreateSongResponse, *error.ErrorResponse)
}

type service struct {
	request    *gorequest.SuperAgent
	config     config.Config
	repository repository.Repository
}

func NewService(request *gorequest.SuperAgent,
	config config.Config,
	repository repository.Repository) Service {
	return &service{
		request:    request,
		config:     config,
		repository: repository,
	}
}

func (service service) FetchFromSpotifyAndInsertIntoDB(context *gin.Context, isrc string) (response.CreateSongResponse, *error.ErrorResponse) {
	service.request.Get(fmt.Sprintf(service.config.GetSpotifySearchApi(), isrc))
	bearerToken := "Bearer " + context.Request.Header.Get("Authorization")
	service.request.Set("Authorization", bearerToken)

	_, body, err := service.request.End()
	if err != nil {
		fmt.Errorf("unable to fetch data. error: %v", err)
		return response.CreateSongResponse{}, error.SpotyfyErrors[error.InternalServerError]
	}

	var res model.TracksSearchResponse
	marshalErr := json.Unmarshal([]byte(body), &res)
	if marshalErr != nil {
		fmt.Errorf("unable to un marshal . error: %v", marshalErr)
		return response.CreateSongResponse{}, error.SpotyfyErrors[error.BadFormattedJSONError]
	}

	insertedId, dbErr := service.repository.Insert(context, res.TransformToDbModel(isrc))
	if dbErr != nil {
		fmt.Printf("unable insert to DB . error: %v", dbErr.Error())
		if strings.Contains(dbErr.Error(), "unique constraint") {
			return response.CreateSongResponse{}, error.SpotyfyErrors[error.DupliacteISRCError]
		}

		return response.CreateSongResponse{}, error.SpotyfyErrors[error.DBInsertionError]
	}

	return response.CreateSongResponse{
		Id: insertedId,
	}, nil
}
