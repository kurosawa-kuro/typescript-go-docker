package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// PingHandler handles the ping endpoint
func PingHandler(c *gin.Context) {
	fmt.Println("PingHandler called")

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
