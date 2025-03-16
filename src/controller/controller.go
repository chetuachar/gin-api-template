package controller

import (
	"gin-api-template/src/config"
	"gin-api-template/src/utils"
	"net/http"
	"path/filepath"

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

func GetLogs(c *gin.Context) {
	// Authentication check
	passwd := c.Query("passwd")
	if passwd != "Chethan@123" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	// Read last 100 logs
	logFile := filepath.Join(
		config.Appconfig.GetString("LOG_PATH"),
		config.Appconfig.GetString("LOG_FILE"))
	logs, err := utils.ReadLastNLogs(logFile, 100)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to load logs")
		return
	}

	// Render logs in HTML
	c.HTML(http.StatusOK, "logs.html", gin.H{"logs": logs})

}
