package test

import (
	"testing"

	"github.com/elzinko/go-fizzbuzz/internal/api/services"
	"github.com/google/go-cmp/cmp"
)

func Test_FizzBuzz_Service_OK(t *testing.T) {
	got := services.FizzBuzz(3, 5, 15, "Fizz", "Buzz")
	want := "1,2,Fizz,4,Buzz,Fizz,7,8,Fizz,Buzz,11,Fizz,13,14,FizzBuzz"
	check(t, want, got)
}

func check(t *testing.T, want string, got string) {
	if !cmp.Equal(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
