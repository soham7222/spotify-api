package router

import (
	"database/sql"
	"spotify-api/config"
	"spotify-api/spotify/save/controller"
	"spotify-api/spotify/save/repository"
	"spotify-api/spotify/save/service"

	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
)

func Init(r *gin.Engine, db *sql.DB, config config.Config, request *gorequest.SuperAgent) {
	repository := repository.NewRepository(db)
	service := service.NewService(request, config, repository)

	controller := controller.NewController(service)

	authRouter := r.Group("/api/spotify")
	authRouterGroup := authRouter.Group("/")
	{
		authRouterGroup.POST("song", controller.Save)
	}
}
