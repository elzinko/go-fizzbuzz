package api

import (
	"github.com/elzinko/go-fizzbuzz/internal/api/router"
	"github.com/gin-gonic/gin"
)

func SetupServer() *gin.Engine {
	web := router.Setup()
	return web
}

func Run() {
	web := SetupServer()
	_ = web.Run()
}
