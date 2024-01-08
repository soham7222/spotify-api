package service

import (
	"spotify-api/error"
	"spotify-api/model"
	"spotify-api/repository"

	"github.com/gin-gonic/gin"
)

type readService struct {
	repository repository.ReadRepository
}

//go:generate mockgen -source=read_service.go -destination=../mocks/mock_read_service.go -package=mocks
type ReadService interface {
	SelectTracksByArtist(ctx *gin.Context, artist string) ([]model.TrackDetailsResponse, *error.ErrorResponse)
	SelectTracksByISRC(ctx *gin.Context, isrc string) ([]model.TrackDetailsResponse, *error.ErrorResponse)
}

func NewReadService(repository repository.ReadRepository) ReadService {
	return &readService{repository: repository}
}

func (r readService) SelectTracksByArtist(ctx *gin.Context, artist string) ([]model.TrackDetailsResponse, *error.ErrorResponse) {
	res, err := r.repository.SelectTracksByArtist(ctx, artist)
	if err != nil {
		return nil, error.SpotyfyErrors[error.InternalServerError]
	}

	return res, nil
}

func (r readService) SelectTracksByISRC(ctx *gin.Context, isrc string) ([]model.TrackDetailsResponse, *error.ErrorResponse) {
	res, err := r.repository.SelectTracksByISRC(ctx, isrc)
	if err != nil {
		return nil, error.SpotyfyErrors[error.InternalServerError]
	}

	return res, nil
}
