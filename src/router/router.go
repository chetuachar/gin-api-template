package router

import (
	"fmt"
	"gin-api-template/src/config"
	"gin-api-template/src/controller"
	"gin-api-template/src/logger"
	"gin-api-template/src/middleware"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	ginratelimit "github.com/ljahier/gin-ratelimit"
)

func Init() {
	router := NewRouter()
	server := &http.Server{
		Addr:         ":" + config.Appconfig.GetString("SERVER_PORT"),
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server.ListenAndServe()

}

func NewRouter() *gin.Engine {
	router := gin.New()
	logger.InfoLn("Listening and serving HTTP on port:" + fmt.Sprintf("%d", config.Appconfig.GetInt("SERVER_PORT")))
	tb := ginratelimit.NewTokenBucket(
		config.Appconfig.GetInt("REQ_RATE_LIMIT"),
		config.Appconfig.GetDuration("REQ_RATE_LIMIT_TIME")*time.Minute)

	// CORS middleware for Front end application
	router.Use(middleware.CORSMiddleware())
	// Load HTML templates
	router.LoadHTMLGlob("./src/templates/*")
	v1 := router.Group("/v1")
	v1.Use(middleware.LogRequestInfo(), middleware.Authenticate)
	{
		v1.GET("/health", controller.HealthCheck)
		v1.GET("/hi", middleware.RateLimitMiddleware(tb), controller.SayHelow)
	}
	router.GET("/logs", controller.GetLogs)

	return router
}
