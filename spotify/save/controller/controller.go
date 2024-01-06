package controller

import (
	"net/http"
	"spotify-api/spotify/save/model/request"
	"spotify-api/spotify/save/service"

	"github.com/gin-gonic/gin"
)

type Controller interface {
	Save(context *gin.Context)
}

type controller struct {
	service service.Service
}

func NewController(service service.Service) Controller {
	return controller{
		service: service,
	}
}

func (c controller) Save(context *gin.Context) {
	req := request.SaveSongRequest{}
	if bindErr := context.ShouldBindJSON(&req); bindErr != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, nil)
	}

	response, err := c.service.FetchFromSpotifyAndInsertIntoDB(context, req.ISRC)

	if err != nil {
		context.AbortWithStatusJSON(err.HttpStatusCode, err)
		return
	}
	context.JSON(http.StatusCreated, response)
}
