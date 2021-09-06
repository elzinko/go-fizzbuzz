package api

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/elzinko/go-fizzbuzz/docs"
	"github.com/elzinko/go-fizzbuzz/internal/api/pkg/config"
	"github.com/elzinko/go-fizzbuzz/internal/api/router"
	"github.com/gin-gonic/gin"
)

// @contact.name Thomas Couderc
// @contact.email thomas.couderc@gmail.com
// @license.name MIT
// @license.url https://github.com/elzinko/go-fizzbuzz/blob/master/LICENSE
// @termsOfService http://swagger.io/terms/

const api_base = "/api"
const api_version = "0.6.1"
const API_PATH = api_base + "/" + api_version

func Run(basePath string, configPath string) {
	web := SetupServer(basePath, configPath)
	log.Info("Go FizzBuzz REST API Running on port" + config.GetConfig().Server.Port)
	_ = web.Run(":" + config.GetConfig().Server.Port)
}

func SetupServer(basePath string, configPath string) *gin.Engine {
	setConfiguration(basePath, configPath)
	initializeLogging(basePath+"log/fuzzbuzz.log", config.GetConfig().Log.Level, config.GetConfig().Log.AppendToLogFile)
	setSwaggerInfos()

	return router.Setup(basePath, API_PATH)
}

func setConfiguration(basePath string, configPath string) {
	config.Load(basePath, configPath)
	gin.SetMode(config.GetConfig().Server.Mode)
}

func initializeLogging(logFile string, logLevel string, appendToLogFile bool) {
	mw := io.MultiWriter(os.Stdout)

	// If specified in config.yml, log to a file
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

func setSwaggerInfos() {
	docs.SwaggerInfo.Title = "FizzBuzz generator using REST"
	docs.SwaggerInfo.Description = "FizzFuzz REST API in Golang with Gin Framework."
	docs.SwaggerInfo.Version = api_version
	docs.SwaggerInfo.BasePath = API_PATH
	docs.SwaggerInfo.Schemes = []string{"https"}
}
