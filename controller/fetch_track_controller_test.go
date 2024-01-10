package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"spotify-api/error"
	"spotify-api/mocks"
	"spotify-api/model"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type ReadTrackControllerTest struct {
	suite.Suite
	mockService  *mocks.MockReadService
	context      *gin.Context
	controller   ReadController
	mockCtrl     *gomock.Controller
	mockRecorder *httptest.ResponseRecorder
}

func TestReadTrackControllerTest(t *testing.T) {
	suite.Run(t, new(ReadTrackControllerTest))
}

func (suite *ReadTrackControllerTest) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.mockRecorder = httptest.NewRecorder()
	suite.context, _ = gin.CreateTestContext(suite.mockRecorder)
	suite.context.Request, _ = http.NewRequest("GET", "", nil)
	suite.mockService = mocks.NewMockReadService(suite.mockCtrl)
	suite.controller = NewReadController(suite.mockService)
}

func (suite *ReadTrackControllerTest) TestGetTracksByArtistReturnsStatusOkWhenSreviceReturnSuccess() {
	url := fmt.Sprintf("api/spotify/track/artist/%s", "omar")
	suite.context.Request, _ = http.NewRequest(http.MethodGet, url, nil)
	suite.context.Params = gin.Params{
		gin.Param{
			Key:   "artist",
			Value: "omar",
		},
	}
	res := []model.TrackDetailsResponse{{
		Isrc:    "UTA",
		ImgURI:  "II",
		Title:   "test-tile",
		Artists: []string{"omar ray"},
	}}

	expectedResponseinBytes, _ := json.Marshal(res)

	suite.mockService.EXPECT().SelectTracksByArtist(suite.context, "omar").Return(res, nil).Times(1)
	suite.controller.GetTracksByArtist(suite.context)
	suite.Equal(http.StatusOK, suite.mockRecorder.Code)
	suite.Equal(string(expectedResponseinBytes), suite.mockRecorder.Body.String())
}

func (suite *ReadTrackControllerTest) TestGetTracksByArtistReturnsStatus500WhenSreviceReturnsError() {
	url := fmt.Sprintf("api/spotify/track/artist/%s", "omar")
	suite.context.Request, _ = http.NewRequest(http.MethodGet, url, nil)
	suite.context.Params = gin.Params{
		gin.Param{
			Key:   "artist",
			Value: "omar",
		},
	}

	suite.mockService.EXPECT().SelectTracksByArtist(suite.context, "omar").Return(nil, error.SpotyfyErrors[error.InternalServerError]).Times(1)
	suite.controller.GetTracksByArtist(suite.context)
	suite.Equal(http.StatusInternalServerError, suite.mockRecorder.Code)
}

func (suite *ReadTrackControllerTest) TestGetTrackByISRCReturnsStatusOkWhenSreviceReturnSuccess() {
	url := fmt.Sprintf("api/spotify/track/isrc/%s", "UTA")
	suite.context.Request, _ = http.NewRequest(http.MethodGet, url, nil)
	suite.context.Params = gin.Params{
		gin.Param{
			Key:   "isrc",
			Value: "UTA",
		},
	}
	res := []model.TrackDetailsResponse{{
		Isrc:    "UTA",
		ImgURI:  "II",
		Title:   "test-tile",
		Artists: []string{"omar ray"},
	}}

	expectedResponseinBytes, _ := json.Marshal(res)

	suite.mockService.EXPECT().SelectTracksByISRC(suite.context, "UTA").Return(res, nil).Times(1)
	suite.controller.GetTrackByISRC(suite.context)
	suite.Equal(http.StatusOK, suite.mockRecorder.Code)
	suite.Equal(string(expectedResponseinBytes), suite.mockRecorder.Body.String())
}

func (suite *ReadTrackControllerTest) TestGetTrackByISRCReturnsStatus500WhenSreviceReturnsError() {
	url := fmt.Sprintf("api/spotify/track/isrc/%s", "UTA")
	suite.context.Request, _ = http.NewRequest(http.MethodGet, url, nil)
	suite.context.Params = gin.Params{
		gin.Param{
			Key:   "isrc",
			Value: "UTA",
		},
	}

	suite.mockService.EXPECT().SelectTracksByISRC(suite.context, "UTA").Return(nil, error.SpotyfyErrors[error.InternalServerError]).Times(1)
	suite.controller.GetTrackByISRC(suite.context)
	suite.Equal(http.StatusInternalServerError, suite.mockRecorder.Code)
}
