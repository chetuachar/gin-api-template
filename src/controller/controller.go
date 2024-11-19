package controller

import (
	"github.com/gin-gonic/gin"
)

func SayHelow(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hi back"})
}
