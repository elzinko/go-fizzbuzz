package controllers

import (
	"net/http"
	"strconv"

	"github.com/elzinko/go-fizzbuzz/internal/api/services"
	"github.com/gin-gonic/gin"
)

func GetFizzFuzz(c *gin.Context) {
	int1Str := c.Query("int1")
	int2Str := c.Query("int2")
	limitStr := c.Query("limit")
	str1 := c.Query("str1")
	str2 := c.Query("str2")

	if int1Str == "" {
		c.JSON(http.StatusPreconditionFailed, "int1 shouldn't be empty")
		return
	}
	if int2Str == "" {
		c.JSON(http.StatusPreconditionFailed, "int2 shouldn't be empty")
		return
	}
	if limitStr == "" {
		c.JSON(http.StatusPreconditionFailed, "limit shouldn't be empty")
		return
	}
	if str1 == "" {
		c.JSON(http.StatusPreconditionFailed, "str1 shouldn't be empty")
		return
	}
	if str2 == "" {
		c.JSON(http.StatusPreconditionFailed, "str2 shouldn't be empty")
		return
	}

	int1, err := strconv.Atoi(int1Str)
	if err != nil {
		c.JSON(http.StatusPreconditionFailed, "int1 should be set to a positive integer")
		return
	}

	int2, err := strconv.Atoi(int2Str)
	if err != nil {
		c.JSON(http.StatusPreconditionFailed, "int2 should be set to a positive integer")
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusPreconditionFailed, "limit should be set to a positive integer")
		return
	}

	result := services.FizzBuzz(int1, int2, limit, str1, str2)

	c.JSON(http.StatusOK, result)
}
