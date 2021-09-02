package router

import (
	"github.com/elzinko/go-fizzbuzz/internal/api/controllers"
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	app := gin.New()
	app.GET("/api/fizzbuzz", controllers.GetFizzFuzz)
	return app
}
