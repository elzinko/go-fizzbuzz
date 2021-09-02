package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"

	"github.com/elzinko/go-fizzbuzz/internal/api"
	"github.com/google/go-cmp/cmp"
)

func Test_FizzBuzz_Controller_OK(t *testing.T) {
	want := "1,2,Fizz,4,Buzz,Fizz,7,8,Fizz,Buzz,11,Fizz,13,14,FizzBuzz"
	got := FizzBuzz(t, http.StatusOK, "3", "5", "15", "Fizz", "Buzz")
	check(t, want, got)
}

func Test_FizzBuzz_Controller_KO_int1_empty(t *testing.T) {
	want := "int1 shouldn't be empty"
	got := FizzBuzz(t, http.StatusPreconditionFailed, "", "5", "15", "Fizz", "Buzz")
	check(t, want, got)
}

func Test_FizzBuzz_Controller_KO_int1_should_be_integer(t *testing.T) {
	want := "int1 should be set to a positive integer"
	got := FizzBuzz(t, http.StatusPreconditionFailed, "shouldnt_be_a_string", "5", "15", "Fizz", "Buzz")
	check(t, want, got)
}

func Test_FizzBuzz_Controller_KO_int2_empty(t *testing.T) {
	want := "int2 shouldn't be empty"
	got := FizzBuzz(t, http.StatusPreconditionFailed, "3", "", "15", "Fizz", "Buzz")
	check(t, want, got)
}

func Test_FizzBuzz_Controller_KO_int2_should_be_integer(t *testing.T) {
	want := "int2 should be set to a positive integer"
	got := FizzBuzz(t, http.StatusPreconditionFailed, "3", "shouldnt_be_a_string", "15", "Fizz", "Buzz")
	check(t, want, got)
}

func Test_FizzBuzz_Controller_KO_limit_empty(t *testing.T) {
	want := "limit shouldn't be empty"
	got := FizzBuzz(t, http.StatusPreconditionFailed, "3", "5", "", "Fizz", "Buzz")
	check(t, want, got)
}

func Test_FizzBuzz_Controller_KO_limit_should_be_integer(t *testing.T) {
	want := "limit should be set to a positive integer"
	got := FizzBuzz(t, http.StatusPreconditionFailed, "3", "5", "shouldnt_be_a_string", "Fizz", "Buzz")
	check(t, want, got)
}

func Test_FizzBuzz_Controller_KO_str1_empty(t *testing.T) {
	want := "str1 shouldn't be empty"
	got := FizzBuzz(t, http.StatusPreconditionFailed, "3", "5", "15", "", "Buzz")
	check(t, want, got)
}

func Test_FizzBuzz_Controller_KO_str2_empty(t *testing.T) {
	want := "str2 shouldn't be empty"
	got := FizzBuzz(t, http.StatusPreconditionFailed, "3", "5", "15", "Fizz", "")
	check(t, want, got)
}

func FizzBuzz(t *testing.T, httpCode int, int1 string, int2 string, limit string, str1 string, str2 string) string {

	web := api.SetupServer("config.yml")
	ts := httptest.NewServer(web)
	defer ts.Close()

	params := url.Values{}
	params.Add("int1", int1)
	params.Add("int2", int2)
	params.Add("limit", limit)
	params.Add("str1", str1)
	params.Add("str2", str2)

	resp, err := http.Get(fmt.Sprintf("%s/api/fizzbuzz?%s", ts.URL, params.Encode()))

	if err != nil {
		fmt.Printf("%s", err)
	}

	defer resp.Body.Close()

	if !cmp.Equal(resp.StatusCode, httpCode) {
		t.Errorf("got %q, wanted %q", resp.StatusCode, httpCode)
	}

	got := extractGot(resp)

	return got
}

func extractGot(resp *http.Response) string {

	got, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("%s", err)
	}

	gotUnQuoted, err := strconv.Unquote(string(got))

	if err != nil {
		fmt.Printf("%s", err)
	}

	return gotUnQuoted

}
