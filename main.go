package main

import (
	"gin-api-template/src/config"
	"gin-api-template/src/logger"
	"gin-api-template/src/router"
)

func main() {
	// @title Gin API Template
	// @version 1.0.0
	// @description This is a sample Gin FAST API basic Setup.

	// @contact.name Chethan NV
	// @contact.url https://chetuachar.github.io
	// @contact.email chethuchethan72@gamil.com


	// @BasePath /
	// @securityDefinitions.apikey ApiKeyAuth
	// @in header
	// @name Auth-Token

	
	config.Init()
	config.Appconfig = config.GetConfig()
	logger.Init()
	logger.InfoLn("Config Logger and Router are  Initialized successfully")
	router.Init()
}
