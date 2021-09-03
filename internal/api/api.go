package api

import (
	log "github.com/sirupsen/logrus"

	"github.com/elzinko/go-fizzbuzz/internal/api/pkg/config"
	"github.com/elzinko/go-fizzbuzz/internal/api/router"
	"github.com/gin-gonic/gin"
)

func Run(configPath string) {
	web := SetupServer(configPath)
	conf := config.GetConfig()
	log.Info("Go FizzBuzz REST API Running on port" + conf.Server.Port)
	_ = web.Run(":" + conf.Server.Port)
}

func SetupServer(configPath string) *gin.Engine {
	if configPath == "" {
		configPath = "data/config.yml"
	}
	setConfiguration(configPath)
	web := router.Setup()
	return web
}

func setConfiguration(configPath string) {
	config.Setup(configPath)
	setLogLevel(config.GetConfig().Log.Level)
	gin.SetMode(config.GetConfig().Server.Mode)
}

func setLogLevel(level string) {
	if level == "debug" {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
	log.Infof("Log level set to %s", level)
}
