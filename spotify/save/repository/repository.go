package repository

import (
	"database/sql"
	"fmt"
	"spotify-api/spotify/save/dto"

	"github.com/gin-gonic/gin"
)

const (
	Insert = "INSERT into TRACK_DETAILS (ISRC, METADATA) VALUES (:1, :2)"
)

type Repository interface {
	Insert(ctx *gin.Context, track dto.TrackDbModel) (int64, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

func (r repository) Insert(ctx *gin.Context, track dto.TrackDbModel) (int64, error) {
	res, err := r.db.ExecContext(ctx.Request.Context(), Insert, track.Isrc, track.Metadata)
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
