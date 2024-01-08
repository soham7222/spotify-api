package service

import (
	"fmt"
	"spotify-api/client"
	"spotify-api/error"
	"spotify-api/model/response"
	"spotify-api/repository"
	"strings"

	"github.com/gin-gonic/gin"
)

//go:generate mockgen -source=save_service.go -destination=../mocks/mock_save_service.go -package=mocks
type SaveService interface {
	FetchFromSpotifyAndInsertIntoDB(context *gin.Context, isrc string) (response.CreateSongResponse, *error.ErrorResponse)
}

type saveService struct {
	spotifyClient client.SpotifyClient
	repository    repository.SaveRepository
}

func NewSaveService(repository repository.SaveRepository,
	spotifyClient client.SpotifyClient) SaveService {
	return &saveService{
		repository:    repository,
		spotifyClient: spotifyClient,
	}
}

func (service saveService) FetchFromSpotifyAndInsertIntoDB(context *gin.Context, isrc string) (response.CreateSongResponse, *error.ErrorResponse) {
	res, err := service.spotifyClient.FetchTrackDetailsBasedOnISRC(context, isrc)
	if err != nil {
		return response.CreateSongResponse{}, error.SpotyfyErrors[error.InternalServerError]
	}

	if res.Tracks.Total == 0 {
		return response.CreateSongResponse{}, error.SpotyfyErrors[error.NoTrackExistsError]
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
