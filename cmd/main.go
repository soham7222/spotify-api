package main

import (
	"spotify-api/config"
	"spotify-api/constants"
	"spotify-api/database"
	"spotify-api/middleware"
	readRouter "spotify-api/spotify/read/router"
	saveRouter "spotify-api/spotify/save/router"

	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
)

func main() {
	r := gin.Default()

	db := database.Initialize()
	request := gorequest.New()

	config, err := config.Load(constants.APP_CONFIG_MOUTH_PATH)
	if err != nil {
		panic("error while loading the config" + err.Error())
	}

	r.Use(middleware.AuthMiddleware(config, request))

	saveRouter.Init(r, db, config, request)
	readRouter.Init(r, db, config, request)

	r.Run()
}
