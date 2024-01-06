package main

import (
	"spotify-api/config"
	"spotify-api/database"
	"spotify-api/middleware"
	saveRouter "spotify-api/spotify/save/router"

	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
)

const CONFIG_PATH = "../config/config.json"

func main() {
	r := gin.Default()

	db := database.Initialize()
	request := gorequest.New()

	config, err := config.Load(CONFIG_PATH)
	if err != nil {
		panic("error while loading the config" + err.Error())
	}

	r.Use(middleware.AuthMiddleware(config, request))

	saveRouter.Init(r, db, config, request)

	r.Run()
}
