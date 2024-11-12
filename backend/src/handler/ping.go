package handler

import (
	"github.com/gin-gonic/gin"
)

// PingHandler handles the ping endpoint
func PingHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
