package test

import (
	"testing"

	"github.com/elzinko/go-fizzbuzz/internal/api/services"
)

func Test_FizzBuzz_Service_OK(t *testing.T) {
	got := services.FizzBuzz(3, 5, 15, "Fizz", "Buzz")
	want := "1,2,Fizz,4,Buzz,Fizz,7,8,Fizz,Buzz,11,Fizz,13,14,FizzBuzz"
	Check(t, want, got)
}
