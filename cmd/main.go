package main

import (
	"net/http"
	"spotify-api/config"
	"spotify-api/constants"
	"spotify-api/database"
	"spotify-api/middleware"
	"spotify-api/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db := database.Initialize()

	config, err := config.Load(constants.APP_CONFIG_MOUTH_PATH)
	if err != nil {
		panic("error while loading the config" + err.Error())
	}

	client := &http.Client{}
	r.Use(middleware.AuthMiddleware(config, client))

	router.Init(r, db, config, client)
	r.Run()
}
