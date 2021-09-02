package test

import (
	"testing"

	"github.com/elzinko/go-fizzbuzz/internal/api/services"
	common_test "github.com/elzinko/go-fizzbuzz/test/common"
)

func Test_FizzBuzz_Service_OK(t *testing.T) {
	got := services.FizzBuzz(3, 5, 15, "Fizz", "Buzz")
	want := "1,2,Fizz,4,Buzz,Fizz,7,8,Fizz,Buzz,11,Fizz,13,14,FizzBuzz"
	common_test.Check(t, want, got)
}
