package api

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/elzinko/go-fizzbuzz/internal/api/pkg/config"
	"github.com/elzinko/go-fizzbuzz/internal/api/router"
	"github.com/gin-gonic/gin"
)

func Run(basePath string, configPath string) {
	web := SetupServer(basePath, configPath)
	log.Info("Go FizzBuzz REST API Running on port" + config.GetConfig().Server.Port)
	_ = web.Run(":" + config.GetConfig().Server.Port)
}

func SetupServer(basePath string, configPath string) *gin.Engine {
	if configPath == "" {
		configPath = "data/config.yml"
	}
	setConfiguration(basePath, configPath)
	initializeLogging(basePath+"log/fuzzbuzz.log", config.GetConfig().Log.Level, config.GetConfig().Log.AppendToLogFile)
	web := router.Setup(basePath)
	return web
}

func setConfiguration(basePath string, configPath string) {
	config.Setup(basePath, configPath)
	gin.SetMode(config.GetConfig().Server.Mode)
}

func initializeLogging(logFile string, logLevel string, appendToLogFile bool) {
	mw := io.MultiWriter(os.Stdout)

	// If specified, log to a file
	if appendToLogFile && logFile != "" {
		file, err := os.Create(logFile)
		if err != nil {
			log.Errorf("Could not open log file : " + err.Error())
		} else {
			mw = io.MultiWriter(os.Stdout, file)
		}
	}

	gin.DefaultWriter = io.MultiWriter(mw)
	gin.DisableConsoleColor()

	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(mw)

	setLogLevel(logLevel)
}

func setLogLevel(level string) {
	if level == "debug" {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
	log.Infof("Log level set to %s", level)
}
