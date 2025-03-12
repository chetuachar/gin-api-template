package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Application health status check.
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func SayHelow(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hi back"})
}
