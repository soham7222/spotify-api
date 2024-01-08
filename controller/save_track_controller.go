package controller

import (
	"net/http"
	"spotify-api/error"
	"spotify-api/model/request"
	"spotify-api/service"

	"github.com/gin-gonic/gin"
)

type SaveController interface {
	Save(context *gin.Context)
}

type saveController struct {
	service service.SaveService
}

func NewSaveController(service service.SaveService) SaveController {
	return saveController{
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
func (c saveController) Save(context *gin.Context) {
	req := request.SaveSongRequest{}
	if bindErr := context.ShouldBindJSON(&req); bindErr != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, error.SpotyfyErrors[error.BadFormattedJSONError])
		return
	}

	response, err := c.service.FetchFromSpotifyAndInsertIntoDB(context, req.ISRC)

	if err != nil {
		context.AbortWithStatusJSON(err.HttpStatusCode, err)
		return
	}
	context.JSON(http.StatusCreated, response)
}
