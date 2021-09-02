package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFizzFuzz(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, "Not yet implemented")
}
