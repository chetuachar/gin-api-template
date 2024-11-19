package logger

import (
	"gin-api-template/src/config"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	logger "github.com/sirupsen/logrus"
)

// Init ..
func Init() {
	customFormatter := new(logger.JSONFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	logger.SetFormatter(customFormatter)
	// logger.SetReportCaller(true)
	logLevel := config.Appconfig.GetString("LOG_LEVEL")
	setLogLevel(logLevel)
	if config.Appconfig.GetBool("LOG_STDOUT") {
		logger.New().Out = os.Stdout
	} else {
		fullPath := filepath.Join(config.Appconfig.GetString("LOG_PATH"), config.Appconfig.GetString("LOG_FILE"))
		logger.SetOutput(&lumberjack.Logger{
			Filename:   fullPath,
			MaxSize:    config.Appconfig.GetInt("LOG_MAX_SIZE"),
			MaxBackups: config.Appconfig.GetInt("LOG_MAX_BACKUPS"),
			MaxAge:     config.Appconfig.GetInt("LOG_MAX_AGE"),
			Compress:   config.Appconfig.GetBool("LOG_COMPRESS"),
		})
	}
}

func setLogLevel(logLevel string) {
	switch strings.ToLower(logLevel) {
	case "debug":
		logger.SetLevel(logger.DebugLevel)
	case "info":
		logger.SetLevel(logger.InfoLevel)
	case "warn":
		logger.SetLevel(logger.WarnLevel)
	case "error":
		logger.SetLevel(logger.ErrorLevel)
	default:
		logger.SetLevel(logger.DebugLevel)
	}
}

// LogInfo ...
func LogInfo(message string, c *gin.Context) {
	logger.WithFields(logger.Fields{
		"path":         c.Request.RequestURI,
		"x-request-id": c.Request.Header.Get("x-request-id"),
		"version":      c.Request.Header.Get("VERSION"),
		"request-ip":   c.ClientIP(),
	}).Info(message)
}

// LogError ...
func LogError(message string, err error, c *gin.Context) {
	logger.WithFields(logger.Fields{
		"path":         c.Request.RequestURI,
		"error":        err.Error(),
		"x-request-id": c.Request.Header.Get("x-request-id"),
		"version":      config.Appconfig.GetString("VERSION"),
		"request-ip":   c.ClientIP(),
	}).Error(message)
}

// LogFatal ...
func LogFatal(message string, errors error, c *gin.Context) {
	logger.WithFields(logger.Fields{
		"path":         c.Request.RequestURI,
		"error":        errors.Error(),
		"x-request-id": c.Request.Header.Get("x-request-id"),
		"version":      config.Appconfig.GetString("VERSION"),
		"request-ip":   c.ClientIP(),
	}).Fatal(message)
}

// LogDebug ..
func LogDebug(message string, path string, xRequestID string, errors error) {
	logger.WithFields(logger.Fields{
		"path":         path,
		"x-request-id": xRequestID,
		"version":      config.Appconfig.GetString("VERSION"),
		"error":        "N/A",
	}).Debug(message)
}

// Panic will exit with status code 2
func PanicLn(message string) {
	logger.Panicln(message)
}

// Fatal will exit with status code 1
func FatalLn(message string) {
	logger.Fatalln(message)
}

// Just log the message as info
func InfoLn(message string) {
	logger.Infoln(message)
}

// Just log the message as warn
func WarnLn(message string) {
	logger.Warnln(message)
}

// Just log the message as debug
func DebugLn(message string) {
	logger.Debugln(message)
}

// Just log the message as debug
func PrintLn(message string) {
	logger.Print(message)
}

// Just log the message as Error
func ErrorLn(message string) {
	logger.Errorln(message)
}
