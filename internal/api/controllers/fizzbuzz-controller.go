package controllers

import (
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/elzinko/go-fizzbuzz/internal/api/services"
	"github.com/gin-gonic/gin"
)

const FIZZBUZZ_PATH = "/fizzbuzz"

// FizzFuzz godoc
// @Summary Get FizzBuzz
// @Description Returns a fizzbuzz strings with numbers from 1 to limit, where all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.
// @Produce json
// @Param int1 query integer true "modulo for str1 parameter"
// @Param int2 query integer true "modulo for str2 parameter"
// @Param limit query integer true "Loop limit generation"
// @Param str1 query string false "Fizz string value"
// @Param str2 query string false "Buzz string value"
// @Success 200 {object} string "ok"
// @Failure 412 {string} string	"Parameter problem"
// @Router /fizzbuzz [get]
func GetFizzFuzz(c *gin.Context) {

	log.Info("controllers.GetFizzFuzz")

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
