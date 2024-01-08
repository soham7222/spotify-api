package repository

import (
	"database/sql"
	"fmt"
	"spotify-api/dto"

	"github.com/gin-gonic/gin"
)

const (
	Insert = "INSERT into TRACK_DETAILS (ISRC, TITLE, IMG_URI, ARTISTS) VALUES (:1, :2, :3, :4)"
)

//go:generate mockgen -source=save_repository.go -destination=../mocks/mock_save_repository.go -package=mocks
type SaveRepository interface {
	Insert(ctx *gin.Context, track dto.TrackDbModel) (int64, error)
}

type saveRepository struct {
	db *sql.DB
}

func NewSaveRepository(db *sql.DB) SaveRepository {
	return &saveRepository{
		db: db,
	}
}

func (r saveRepository) Insert(ctx *gin.Context, track dto.TrackDbModel) (int64, error) {
	res, err := r.db.ExecContext(ctx.Request.Context(), Insert, track.Isrc, track.Title, track.ImgURI, track.Artists)
	if err != nil {
		fmt.Errorf("error while inserting record %v", err)
		return 0, err
	}

	lastInsertedId, err := res.LastInsertId()
	if err != nil {
		fmt.Errorf("error while gettung the last inserted id %v", err)
		return 0, err
	}

	return lastInsertedId, nil
}
