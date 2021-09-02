package router

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/elzinko/go-fizzbuzz/internal/api/controllers"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	app := gin.New()
	initializeLogging("log/fizzbuzz.log")
	app.Use(gin.Recovery())

	app.GET("/api/fizzbuzz", controllers.GetFizzFuzz)
	return app
}

func initializeLogging(logFile string) {
	file, err := os.Create("log/api.log")
	// var file, err = os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Errorf("Could Not Open Log File : " + err.Error())
	}

	mw := io.MultiWriter(os.Stdout, file)

	log.SetOutput(mw)
	log.SetFormatter(&log.JSONFormatter{})
}
