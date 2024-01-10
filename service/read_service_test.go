package service

import (
	"errors"
	"spotify-api/error"
	"spotify-api/mocks"
	"spotify-api/model"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type ReadTrackServicerTest struct {
	suite.Suite
	mockRepository *mocks.MockReadRepository
	context        *gin.Context
	service        ReadService
	mockCtrl       *gomock.Controller
}

func TestReadTrackControllerTest(t *testing.T) {
	suite.Run(t, new(ReadTrackServicerTest))
}

func (suite *ReadTrackServicerTest) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.mockRepository = mocks.NewMockReadRepository(suite.mockCtrl)
	suite.service = NewReadService(suite.mockRepository)
}

func (suite *ReadTrackServicerTest) TestSelectTracksByISRCReturnsStatusOkWhenRepositoryReturnSuccess() {
	suite.mockRepository.EXPECT().SelectTracksByISRC(suite.context, "omar").Return([]model.TrackDetailsResponse{{
		Isrc:    "UTA",
		ImgURI:  "II",
		Title:   "test-tile",
		Artists: []string{"omar ray"},
	}}, nil).Times(1)
	res, err := suite.service.SelectTracksByISRC(suite.context, "omar")
	suite.Equal([]model.TrackDetailsResponse{{
		Isrc:    "UTA",
		ImgURI:  "II",
		Title:   "test-tile",
		Artists: []string{"omar ray"},
	}}, res)
	suite.Nil(err)
}

func (suite *ReadTrackServicerTest) TestSelectTracksByISRCReturnsErrorWhenRepositoryReturnError() {
	suite.mockRepository.EXPECT().SelectTracksByISRC(suite.context, "omar").Return(nil, errors.New("new error")).Times(1)
	res, err := suite.service.SelectTracksByISRC(suite.context, "omar")
	suite.Nil(res)
	suite.Equal(err, error.SpotyfyErrors[error.InternalServerError])
}

func (suite *ReadTrackServicerTest) TestSelectTracksByArtistReturnsStatusOkWhenRepositoryReturnSuccess() {
	suite.mockRepository.EXPECT().SelectTracksByArtist(suite.context, "omar").Return([]model.TrackDetailsResponse{{
		Isrc:    "UTA",
		ImgURI:  "II",
		Title:   "test-tile",
		Artists: []string{"omar ray"},
	}}, nil).Times(1)
	res, err := suite.service.SelectTracksByArtist(suite.context, "omar")
	suite.Equal([]model.TrackDetailsResponse{{
		Isrc:    "UTA",
		ImgURI:  "II",
		Title:   "test-tile",
		Artists: []string{"omar ray"},
	}}, res)
	suite.Nil(err)
}

func (suite *ReadTrackServicerTest) TestSelectTracksByArtistReturnsErrorWhenRepositoryReturnError() {
	suite.mockRepository.EXPECT().SelectTracksByArtist(suite.context, "omar").Return(nil, errors.New("new error")).Times(1)
	res, err := suite.service.SelectTracksByArtist(suite.context, "omar")
	suite.Nil(res)
	suite.Equal(err, error.SpotyfyErrors[error.InternalServerError])
}
