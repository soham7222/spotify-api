package repository

import (
	"spotify-api/model"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
)

type ReadTrackRepositoryTest struct {
	suite.Suite
	sqlmock    sqlmock.Sqlmock
	context    *gin.Context
	repository ReadRepository
}

func TestReadTrackRepositoryTest(t *testing.T) {
	suite.Run(t, new(ReadTrackRepositoryTest))
}

func (suite *ReadTrackRepositoryTest) SetupTest() {
	db, mock, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	suite.sqlmock = mock
	suite.repository = NewReadRepository(db)
}

func (suite *ReadTrackRepositoryTest) TestSelectTracksByISRCSuccess() {
	suite.sqlmock.ExpectQuery(SelectTracksByISRC).
		WithArgs("UAT").
		WillReturnRows(suite.sqlmock.NewRows([]string{"ID", "ISRC", "TITLE", "IMG_URI", "ARTISTS"}).AddRow(
			"1", "UAT", "TEST", "TEST.Com", `"{"artists":["Bruce Springsteen"]}"`,
		))

	res, err := suite.repository.SelectTracksByISRC(suite.context, "UAT")
	suite.Equal([]model.TrackDetailsResponse{
		{
			Isrc:   "UAT",
			Title:  "TEST",
			ImgURI: "TEST.Com",
		},
	}, res)
	suite.Nil(err)
}

func (suite *ReadTrackRepositoryTest) TestSelectTracksByArtistSuccess() {
	suite.sqlmock.ExpectQuery(SelectTracksByArtist).
		WithArgs("Bruce").
		WillReturnRows(suite.sqlmock.NewRows([]string{"ID", "ISRC", "TITLE", "IMG_URI", "ARTISTS"}).AddRow(
			"1", "UAT", "TEST", "TEST.Com", `"{"artists":["Bruce Springsteen"]}"`,
		))

	res, err := suite.repository.SelectTracksByArtist(suite.context, "Bruce")
	suite.Equal([]model.TrackDetailsResponse{
		{
			Isrc:   "UAT",
			Title:  "TEST",
			ImgURI: "TEST.Com",
		},
	}, res)
	suite.Nil(err)
}
