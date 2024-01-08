package service

import (
	"fmt"
	"spotify-api/error"
	"spotify-api/spotify/read/model"
	"spotify-api/spotify/read/repository"

	"github.com/gin-gonic/gin"
)

type service struct {
	repository repository.Repository
}

type Service interface {
	SelectTracksByArtist(ctx *gin.Context, artist string) ([]model.TrackDetailsResponse, *error.ErrorResponse)
	SelectTracksByISRC(ctx *gin.Context, isrc string) ([]model.TrackDetailsResponse, *error.ErrorResponse)
}

func NewService(repository repository.Repository) Service {
	return &service{repository: repository}
}

func (r service) SelectTracksByArtist(ctx *gin.Context, artist string) ([]model.TrackDetailsResponse, *error.ErrorResponse) {
	res, err := r.repository.SelectTracksByArtist(ctx, artist)
	if err != nil {
		return nil, error.SpotyfyErrors[error.InternalServerError]
	}

	return res, nil
}

func (r service) SelectTracksByISRC(ctx *gin.Context, isrc string) ([]model.TrackDetailsResponse, *error.ErrorResponse) {
	res, err := r.repository.SelectTracksByISRC(ctx, isrc)
	if err != nil {
		return nil, error.SpotyfyErrors[error.InternalServerError]
	}

	fmt.Printf("res", res)
	return res, nil
}
