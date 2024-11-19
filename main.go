package main

import (
	"gin-api-template/src/config"
	"gin-api-template/src/logger"
	"gin-api-template/src/router"
)

func main() {
	config.Init()
	config.Appconfig = config.GetConfig()
	logger.Init()
	logger.InfoLn("Config Logger and Router are  Initialized successfully")
	router.Init()
}
