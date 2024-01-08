package router

import (
	"database/sql"
	"net/http"
	"spotify-api/client"
	"spotify-api/config"
	"spotify-api/controller"
	"spotify-api/repository"
	"spotify-api/service"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine, db *sql.DB, config config.Config, httpClient *http.Client) {
	readRepository := repository.NewReadRepository(db)
	readService := service.NewReadService(readRepository)
	readController := controller.NewReadController(readService)

	saveRepository := repository.NewSaveRepository(db)
	spotifyClient := client.NewSpotifyClient(httpClient, config)
	saveService := service.NewSaveService(saveRepository, spotifyClient)

	saveController := controller.NewSaveController(saveService)

	authRouter := r.Group("/api/spotify")
	authRouterGroup := authRouter.Group("/")
	{
		authRouterGroup.GET("track/artist/:artist", readController.GetTracksByArtist)
		authRouterGroup.GET("track/isrc/:isrc", readController.GetTrackByISRC)
	}
	{
		authRouterGroup.POST("track", saveController.Save)
	}
}
