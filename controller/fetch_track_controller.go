package controller

import (
	"net/http"
	"spotify-api/service"

	"github.com/gin-gonic/gin"
)

type ReadController interface {
	GetTracksByArtist(context *gin.Context)
	GetTrackByISRC(context *gin.Context)
}

type readController struct {
	service service.ReadService
}

func NewReadController(service service.ReadService) ReadController {
	return readController{
		service: service,
	}
}

func (c readController) GetTracksByArtist(context *gin.Context) {
	artist := context.Param("artist")
	res, err := c.service.SelectTracksByArtist(context, artist)
	if err != nil {
		context.AbortWithStatusJSON(err.HttpStatusCode, err)
		return
	}

	if res == nil {
		context.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}

	context.JSON(http.StatusOK, res)
}

func (c readController) GetTrackByISRC(context *gin.Context) {
	isrc := context.Param("isrc")
	res, err := c.service.SelectTracksByISRC(context, isrc)
	if err != nil {
		context.AbortWithStatusJSON(err.HttpStatusCode, err)
		return
	}

	if res == nil {
		context.AbortWithStatusJSON(http.StatusNotFound, err)
		return
	}

	context.JSON(http.StatusOK, res)
}
