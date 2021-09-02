package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetFizzBuzz(c *gin.Context) {

	int1Str := c.Query("int1")
	int2Str := c.Query("int2")
	limitStr := c.Query("limit")
	str1 := c.Query("str1")
	str2 := c.Query("str2")

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

	result := FizzBuzz(int1, int2, limit, str1, str2)

	c.JSON(http.StatusOK, result)
}

func FizzBuzz(int1 int, int2 int, limit int, str1 string, str2 string) string {
	result := ""

	for i := 1; i <= limit; i++ {
		if i%int1 == 0 {
			// multiple of in1
			result += str1
		}
		if i%int2 == 0 {
			// multiple of in2
			result += str2
		}
		if i%int1 != 0 && i%int2 != 0 {
			// other cases
			result += strconv.Itoa(i)
		}
		if i != limit {
			result += ","
		}
	}
	return result
}
