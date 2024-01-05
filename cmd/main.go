package main

import (
	"spotify-api/config"
	"spotify-api/middleware"
	"spotify-api/spotify/save/controller"
	"spotify-api/spotify/save/service"

	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
)

const CONFIG_PATH = "../config/config.json"

func main() {
	r := gin.Default()

	config, err := config.Load(CONFIG_PATH)
	if err != nil {
		panic("error while loading the config" + err.Error())
	}

	r.Use(middleware.AuthMiddleware(config))

	request := gorequest.New()
	service := service.NewService(request, config)

	controller := controller.NewController(service)

	authRouter := r.Group("/api/spotify")
	authRouterGroup := authRouter.Group("/")
	{
		authRouterGroup.POST("song", controller.Save)
	}

	r.Run()
}
