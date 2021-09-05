package test

import (
	"net/http"
	"testing"
)

func Test_FizzBuzz_Controller_OK(t *testing.T) {
	want := "1,2,Fizz,4,Buzz,Fizz,7,8,Fizz,Buzz,11,Fizz,13,14,FizzBuzz"
	got := FizzBuzz(t, http.StatusOK, "3", "5", "15", "Fizz", "Buzz")
	Check(t, want, got)
}

func Test_FizzBuzz_Controller_KO_int1_empty(t *testing.T) {
	want := "int1 shouldn't be empty"
	got := FizzBuzz(t, http.StatusPreconditionFailed, "", "5", "15", "Fizz", "Buzz")
	Check(t, want, got)
}

func Test_FizzBuzz_Controller_KO_int1_should_be_integer(t *testing.T) {
	want := "int1 should be set to a positive integer"
	got := FizzBuzz(t, http.StatusPreconditionFailed, "shouldnt_be_a_string", "5", "15", "Fizz", "Buzz")
	Check(t, want, got)
}

func Test_FizzBuzz_Controller_KO_int2_empty(t *testing.T) {
	want := "int2 shouldn't be empty"
	got := FizzBuzz(t, http.StatusPreconditionFailed, "3", "", "15", "Fizz", "Buzz")
	Check(t, want, got)
}

func Test_FizzBuzz_Controller_KO_int2_should_be_integer(t *testing.T) {
	want := "int2 should be set to a positive integer"
	got := FizzBuzz(t, http.StatusPreconditionFailed, "3", "shouldnt_be_a_string", "15", "Fizz", "Buzz")
	Check(t, want, got)
}

func Test_FizzBuzz_Controller_KO_limit_empty(t *testing.T) {
	want := "limit shouldn't be empty"
	got := FizzBuzz(t, http.StatusPreconditionFailed, "3", "5", "", "Fizz", "Buzz")
	Check(t, want, got)
}

func Test_FizzBuzz_Controller_KO_limit_should_be_integer(t *testing.T) {
	want := "limit should be set to a positive integer"
	got := FizzBuzz(t, http.StatusPreconditionFailed, "3", "5", "shouldnt_be_a_string", "Fizz", "Buzz")
	Check(t, want, got)
}

func Test_FizzBuzz_Controller_KO_str1_empty(t *testing.T) {
	want := "str1 shouldn't be empty"
	got := FizzBuzz(t, http.StatusPreconditionFailed, "3", "5", "15", "", "Buzz")
	Check(t, want, got)
}

func Test_FizzBuzz_Controller_KO_str2_empty(t *testing.T) {
	want := "str2 shouldn't be empty"
	got := FizzBuzz(t, http.StatusPreconditionFailed, "3", "5", "15", "Fizz", "")
	Check(t, want, got)
}
