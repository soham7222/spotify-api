package router

import (
	"database/sql"
	"net/http"
	"spotify-api/config"
	"spotify-api/spotify/client"
	"spotify-api/spotify/save/controller"
	"spotify-api/spotify/save/repository"
	"spotify-api/spotify/save/service"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine, db *sql.DB, config config.Config, httpClient *http.Client) {
	repository := repository.NewSaveRepository(db)
	spotifyClient := client.NewSpotifyClient(httpClient, config)
	service := service.NewSaveService(repository, spotifyClient)

	controller := controller.NewSaveController(service)

	authRouter := r.Group("/api/spotify")
	authRouterGroup := authRouter.Group("/")
	{
		authRouterGroup.POST("track", controller.Save)
	}
}
