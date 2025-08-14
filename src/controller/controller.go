package controller

import (
	"gin-api-template/src/config"
	"gin-api-template/src/utils"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// Application health status check.
// @Summary Check the gin server health check
// @Schemes
// @Description do health check
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /health [get]
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}


// Say Helow for the gin api
// @Summary Test api sending
// @Tags v1
// @Accept json
// @Produce json
// @Success 200 {object} map[string]string
// @Router /v1/hi [get]
// @Security ApiKeyAuth
func SayHelow(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Hi back"})
}


// GetLogs godoc
// @Summary View last 100 logs
// @Description Returns the last 100 log entries in HTML format. Requires a password for authentication.
// @Tags logs
// @Param passwd query string true "Password for log access"
// @Produce html
// @Success 200 {string} string "HTML page with logs"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {string} string "Failed to load logs"
// @Router /logs [get]
func GetLogs(c *gin.Context) {
	// Authentication check
	passwd := c.Query("passwd")
	if passwd != config.Appconfig.GetString("LOG_SHOW_PASSWD") {
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
