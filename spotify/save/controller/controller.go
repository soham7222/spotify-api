package controller

import (
	"net/http"
	"spotify-api/error"
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

// Save godoc
// @Summary      Save Song details
// @Description  Save Song details based on ISRC
// @Tags         spotify
// @Accept       json
// @Produce      json
// @Param        requestBody body request.SaveSongRequest true "Request Body"
// @Success      201  {object}  response.CreateSongResponse
// @Failure      400  {object}  error.ErrorResponse
// @Failure      500  {object}  error.ErrorResponse
// @Router       /api/spotify/track [post]
func (c controller) Save(context *gin.Context) {
	req := request.SaveSongRequest{}
	if bindErr := context.ShouldBindJSON(&req); bindErr != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, error.SpotyfyErrors[error.BadFormattedJSONError])
	}

	response, err := c.service.FetchFromSpotifyAndInsertIntoDB(context, req.ISRC)

	if err != nil {
		context.AbortWithStatusJSON(err.HttpStatusCode, err)
		return
	}
	context.JSON(http.StatusCreated, response)
}
