package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"spotify-api/error"
	"spotify-api/mocks"
	"spotify-api/model/request"
	"spotify-api/model/response"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type SaveTrackControllerTest struct {
	suite.Suite
	mockService  *mocks.MockSaveService
	context      *gin.Context
	controller   SaveController
	mockCtrl     *gomock.Controller
	mockRecorder *httptest.ResponseRecorder
}

func TestSaveTrackControllerTest(t *testing.T) {
	suite.Run(t, new(SaveTrackControllerTest))
}

func (suite *SaveTrackControllerTest) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.mockRecorder = httptest.NewRecorder()
	suite.context, _ = gin.CreateTestContext(suite.mockRecorder)
	suite.context.Request, _ = http.NewRequest("POST", "", nil)
	suite.mockService = mocks.NewMockSaveService(suite.mockCtrl)
	suite.controller = NewSaveController(suite.mockService)
}

func (suite *SaveTrackControllerTest) TestSaveReturnsStatusOkWhenSreviceReturnSuccess() {
	url := fmt.Sprint("api/spotify/track")
	req := request.SaveSongRequest{
		ISRC: "123",
	}

	data, _ := json.Marshal(req)
	suite.context.Request, _ = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))

	suite.mockService.EXPECT().FetchFromSpotifyAndInsertIntoDB(suite.context, "123").Return(response.CreateSongResponse{
		Id: 123,
	}, nil).Times(1)
	suite.controller.Save(suite.context)
	suite.Equal(http.StatusCreated, suite.mockRecorder.Code)
}

func (suite *SaveTrackControllerTest) TestSaveReturnsStatus500WhenSreviceReturnsError() {
	url := fmt.Sprint("api/spotify/track")
	req := request.SaveSongRequest{
		ISRC: "123",
	}

	data, _ := json.Marshal(req)
	suite.context.Request, _ = http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))

	suite.mockService.EXPECT().FetchFromSpotifyAndInsertIntoDB(suite.context, "123").Return(response.CreateSongResponse{}, error.SpotyfyErrors[error.InternalServerError]).Times(1)
	suite.controller.Save(suite.context)
	suite.Equal(http.StatusInternalServerError, suite.mockRecorder.Code)
}
