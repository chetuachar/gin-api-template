package middleware

import (
	"encoding/json"
	"fmt"
	"gin-api-template/src/config"
	"gin-api-template/src/logger"
	"time"

	"github.com/gin-gonic/gin"
	ginratelimit "github.com/ljahier/gin-ratelimit"
)

// CORSMiddleware function to handle the CORS.
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers",
			"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Auth-Token, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

// Loggin the https requiest.
func LogRequestInfo() gin.HandlerFunc {
	return func(c *gin.Context) {
		// accessLogMap := make(map[string]string)
		accessLogMap := make(map[string]interface{})
		accessLogMap["request_time"] = time.Now().String()
		accessLogMap["request_method"] = c.Request.Method
		accessLogMap["request_uri"] = c.Request.RequestURI
		accessLogMap["request_proto"] = c.Request.Proto
		accessLogMap["request_ua"] = c.Request.UserAgent()
		accessLogMap["request_referer"] = c.Request.Referer()
		accessLogMap["request_post_data"] = c.Request.PostForm.Encode()
		accessLogMap["request_client_ip"] = c.ClientIP()
		accessLogMap["x-request-id"] = c.Request.Header.Get("X-Request-Id")
		accessLogMap["version"] = config.Appconfig.GetString("VERSION")
		logData, err := json.Marshal(accessLogMap)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		accessLogJSON := string(logData)
		logger.InfoLn(accessLogJSON)
		// log.Println(accessLogJSON)
		c.Next()
	}
}

// Validation of Auth-Token
func Authenticate(c *gin.Context) {
	if !(c.Request.Header.Get("Auth-Token") == config.Appconfig.GetString("SELF_API_TOKEN")) {
		c.AbortWithStatusJSON(401, gin.H{
			"Message": "Token is missing or Not a valid token",
		})
		return
	}
	c.Next()
}

// RateLimit the end points
func RateLimitMiddleware(tb *ginratelimit.TokenBucket) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !tb.Allow(c.ClientIP()) {
			c.JSON(429, gin.H{
				"error": "Too Many Requests",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
