package services

import (
	"strconv"
)

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
