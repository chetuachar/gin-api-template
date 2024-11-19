package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SayHelow(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hi back"})
}

func MonitorEmailService(c *gin.Context) {
	if domain := c.Query("domain"); domain == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "In request query domain name not found",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{"message": domain})
	}
}
