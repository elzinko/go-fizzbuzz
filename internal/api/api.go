package api

import (
	log "github.com/sirupsen/logrus"

	"github.com/elzinko/go-fizzbuzz/internal/api/pkg/config"
	"github.com/elzinko/go-fizzbuzz/internal/api/router"
	"github.com/gin-gonic/gin"
)

func setConfiguration(configPath string) {
	config.Setup(configPath)
	gin.SetMode(config.GetConfig().Server.Mode)
}

func SetupServer(configPath string) *gin.Engine {
	if configPath == "" {
		configPath = "data/config.yml"
	}
	setConfiguration(configPath)
	conf := config.GetConfig()
	web := router.Setup()
	log.Info("Go FizzBuzz REST API Running on port " + conf.Server.Port)
	return web
}

func Run(configPath string) {
	web := SetupServer(configPath)
	conf := config.GetConfig()
	_ = web.Run(":" + conf.Server.Port)
}
