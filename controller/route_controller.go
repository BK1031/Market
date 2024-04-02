package controller

import (
	"github.com/gin-gonic/gin"
	"market-mono/utils"
)

func InitializeRoutes(router *gin.Engine) {
	router.GET("/ping", Ping)
}

func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		utils.SugarLogger.Infoln("REQUEST ID: " + c.GetHeader("Request-ID"))
		c.Next()
	}
}

func AuthChecker() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") != "" {
			utils.SugarLogger.Infoln("User token provided", c.GetHeader("Authorization"))
		} else {
			utils.SugarLogger.Infoln("No user token provided")
		}
		c.Next()
	}
}
