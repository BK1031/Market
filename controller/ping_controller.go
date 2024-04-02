package controller

import (
	"github.com/gin-gonic/gin"
	"market-mono/config"
	"net/http"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Market v" + config.Version + " is online!"})
}
