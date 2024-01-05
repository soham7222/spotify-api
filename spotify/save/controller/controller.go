package controller

import (
	"net/http"
	"spotify-api/spotify/save/model"
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
	request := model.SaveSongRequest{}
	if bindErr := context.ShouldBindJSON(&request); bindErr != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, nil)
	}

	response, err := c.service.FetchFromSpotify(context, request.ISRC)

	if err != nil {
		context.AbortWithStatusJSON(err.HttpStatusCode, err)
	}
	context.JSON(http.StatusOK, response)
}
