package service

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"spotify-api/client/model"
	"spotify-api/mocks"
	"spotify-api/model/response"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type SaveTrackServicerTest struct {
	suite.Suite
	mockRepository *mocks.MockSaveRepository
	context        *gin.Context
	service        SaveService
	mockCtrl       *gomock.Controller
	mockClient     *mocks.MockSpotifyClient
	mockRecorder   *httptest.ResponseRecorder
}

func TestSaveTrackControllerTest(t *testing.T) {
	suite.Run(t, new(SaveTrackServicerTest))
}

func (suite *SaveTrackServicerTest) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.mockRepository = mocks.NewMockSaveRepository(suite.mockCtrl)
	suite.mockClient = mocks.NewMockSpotifyClient(suite.mockCtrl)
	suite.service = NewSaveService(suite.mockRepository, suite.mockClient)
	suite.context, _ = gin.CreateTestContext(suite.mockRecorder)
	suite.context.Request, _ = http.NewRequest("POST", "", nil)
}

func (suite *SaveTrackServicerTest) TestReturnsErrorWhenFetchTrackDetailsBasedOnISRCRetrunsNoTracks() {
	suite.mockClient.EXPECT().FetchTrackDetailsBasedOnISRC(suite.context,
		"123").Return(model.TracksSearchResponse{
		Tracks: model.TracksCollection{
			Total: 0,
		},
	}, nil)

	res, err := suite.service.FetchFromSpotifyAndInsertIntoDB(suite.context, "123")
	suite.Equal(response.CreateSongResponse{Id: 0}, res)
	suite.NotNil(err)
}

func (suite *SaveTrackServicerTest) TestReturnsSuccessWhenFetchTrackDetailsBasedOnISRCRetrunsSuccess() {
	jsonData, err := ioutil.ReadFile("../mockTestData/data1.json")
	suite.Nil(err)

	var res model.TracksSearchResponse
	err = json.Unmarshal(jsonData, &res)
	suite.Nil(err)

	suite.mockClient.EXPECT().FetchTrackDetailsBasedOnISRC(suite.context, "QZHNA1928786").Return(res, nil)
	suite.mockRepository.EXPECT().Insert(suite.context, res.TransformToDbModel("QZHNA1928786")).Return(int64(12), nil)
	result, serErr := suite.service.FetchFromSpotifyAndInsertIntoDB(suite.context, "QZHNA1928786")
	suite.Nil(serErr)
	suite.Equal(response.CreateSongResponse{
		Id: int64(12),
	}, result)
}

func (suite *SaveTrackServicerTest) TestReturnsErrorWhenFetchTrackDetailsBasedOnISRCRetrunsError() {
	suite.mockClient.EXPECT().FetchTrackDetailsBasedOnISRC(suite.context, "QZHNA1928786").Return(model.TracksSearchResponse{}, nil)
	_, serErr := suite.service.FetchFromSpotifyAndInsertIntoDB(suite.context, "QZHNA1928786")
	suite.NotNil(serErr)
}

func (suite *SaveTrackServicerTest) TestReturnsSuccessWhenInsertFails() {
	jsonData, err := ioutil.ReadFile("../mockTestData/data1.json")
	suite.Nil(err)

	var res model.TracksSearchResponse
	err = json.Unmarshal(jsonData, &res)
	suite.Nil(err)

	suite.mockClient.EXPECT().FetchTrackDetailsBasedOnISRC(suite.context, "QZHNA1928786").Return(res, nil)
	suite.mockRepository.EXPECT().Insert(suite.context, res.TransformToDbModel("QZHNA1928786")).Return(int64(0), errors.New("something is wrong"))
	_, serErr := suite.service.FetchFromSpotifyAndInsertIntoDB(suite.context, "QZHNA1928786")
	suite.NotNil(serErr)
}
