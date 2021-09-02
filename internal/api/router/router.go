package router

import (
	"fmt"
	"io"
	"os"

	"github.com/elzinko/go-fizzbuzz/internal/api/controllers"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	app := gin.New()

	// Logging to a file
	f, _ := os.Create("log/api.log")
	gin.DisableConsoleColor()
	gin.DefaultWriter = io.MultiWriter(f)

	// Logging format
	app.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - - [%s] \"%s %s %s %d %s \" \" %s\" \" %s\"\n",
			param.ClientIP,
			param.TimeStamp,
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))
	app.Use(gin.Recovery())

	app.GET("/api/fizzbuzz", controllers.GetFizzFuzz)
	return app
}
