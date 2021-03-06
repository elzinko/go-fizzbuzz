package services

import (
	"strconv"

	log "github.com/sirupsen/logrus"
)

func FizzBuzz(int1 int, int2 int, limit int, str1 string, str2 string) string {

	log.Debugf("call services.FizzBuzz(%s,%s, %s, %s, %s)", strconv.Itoa(int1), strconv.Itoa(int2), strconv.Itoa(limit), str1, str2)

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

	log.Debugf("result = %s", result)

	return result
}
