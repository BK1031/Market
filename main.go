package main

import (
	"github.com/gin-gonic/gin"
	"market-mono/config"
	"market-mono/controller"
	"market-mono/database"
	"market-mono/utils"
)

var router *gin.Engine

func setupRouter() *gin.Engine {
	if config.Env == "PROD" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.Default()
	r.Use(controller.RequestLogger())
	r.Use(controller.AuthChecker())
	return r
}

func main() {
	utils.InitializeLogger()
	defer utils.Logger.Sync()

	router = setupRouter()
	database.InitializeDB("disable", "UTC")

	controller.InitializeRoutes(router)
	router.Run(":" + config.Port)
}
