package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"spotify-api/dto"
	"spotify-api/model"

	"github.com/gin-gonic/gin"
)

const (
	SelectTracksByArtist = "SELECT * FROM TRACK_DETAILS WHERE DBMS_LOB.INSTR(ARTISTS, :1) > 0"
	SelectTracksByISRC   = "SELECT * FROM TRACK_DETAILS WHERE ISRC=:1"
)

//go:generate mockgen -source=read_repository.go -destination=../mocks/mock_read_repository.go -package=mocks
type ReadRepository interface {
	SelectTracksByArtist(ctx *gin.Context, artist string) ([]model.TrackDetailsResponse, error)
	SelectTracksByISRC(ctx *gin.Context, artist string) ([]model.TrackDetailsResponse, error)
}

type readRepository struct {
	db *sql.DB
}

func NewReadRepository(db *sql.DB) ReadRepository {
	return &readRepository{
		db: db,
	}
}

func (r readRepository) SelectTracksByArtist(ctx *gin.Context, artist string) ([]model.TrackDetailsResponse, error) {
	rows, err := r.db.Query(SelectTracksByArtist, artist)
	if err != nil {
		fmt.Errorf("error fetching db: %w", err)
		return nil, err
	}

	defer rows.Close()
	return serializeSqlRowData(rows)
}

func (r readRepository) SelectTracksByISRC(ctx *gin.Context, isrc string) ([]model.TrackDetailsResponse, error) {
	rows, err := r.db.Query(SelectTracksByISRC, isrc)
	if err != nil {
		fmt.Errorf("error fetching db: %w", err)
		return nil, err
	}

	defer rows.Close()
	return serializeSqlRowData(rows)
}

func serializeSqlRowData(rows *sql.Rows) ([]model.TrackDetailsResponse, error) {
	result := make([]model.TrackDetailsResponse, 0)
	for rows.Next() {
		var id int
		var isrc string
		var imgURI string
		var title string
		var artists string

		err := rows.Scan(&id, &isrc, &title, &imgURI, &artists)
		if err != nil {
			fmt.Errorf("error while rows.Scan() %w", err)
			return nil, err
		}

		var artistsModel dto.TrackArtistsData
		err = json.Unmarshal([]byte(artists), &artistsModel)
		if err != nil {
			fmt.Errorf("error while un marshal %w", err)
		}

		track := model.TrackDetailsResponse{
			Title:   title,
			Isrc:    isrc,
			ImgURI:  imgURI,
			Artists: artistsModel.Artist,
		}

		result = append(result, track)
	}

	if err := rows.Err(); err != nil {
		fmt.Errorf("error while rows.Next() %v", err)
		return nil, err
	}

	return result, nil
}
