package router

import (
	"database/sql"
	"spotify-api/config"
	"spotify-api/spotify/read/controller"
	"spotify-api/spotify/read/repository"
	"spotify-api/spotify/read/service"

	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine, db *sql.DB, config config.Config) {
	repository := repository.NewReadRepository(db)
	service := service.NewReadService(repository)
	controller := controller.NewReadController(service)

	authRouter := r.Group("/api/spotify")
	authRouterGroup := authRouter.Group("/")
	{
		authRouterGroup.GET("track/artist/:artist", controller.GetTracksByArtist)
		authRouterGroup.GET("track/isrc/:isrc", controller.GetTrackByISRC)
	}
}
