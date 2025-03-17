package config

import (
	"github.com/spf13/viper"
)

var config *viper.Viper
var Appconfig *viper.Viper

func Init() {
	config = viper.New()
	setDefaults()
	config.AutomaticEnv()

}

func setDefaults() {
	config.SetDefault("LOG_LEVEL", "debug")
	config.SetDefault("LOG_STDOUT", "false")
	config.SetDefault("LOG_FILE", "gin-api.log")
	config.SetDefault("LOG_PATH", "./log")
	config.SetDefault("LOG_MAX_SIZE", 10)
	config.SetDefault("LOG_MAX_BACKUPS", 3)
	config.SetDefault("LOG_MAX_AGE", 7)
	config.SetDefault("LOG_COMPRESS", false)
	config.SetDefault("SERVER_PORT", 8086)
	config.SetDefault("VERSION", "v1.0.0")
	config.SetDefault("SELF_API_TOKEN", "f144a0da7c41f4725d83c50ea4fdf1bd")
	config.SetDefault("REQ_RATE_LIMIT", 10)
	config.SetDefault("REQ_RATE_LIMIT_TIME", 5)
	config.SetDefault("LOG_SHOW_PASSWD", "Chethan@123")
}

func GetConfig() *viper.Viper {
	return config
}
