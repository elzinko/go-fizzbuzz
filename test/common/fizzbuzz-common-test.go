package common_test

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

func Check(t *testing.T, want string, got string) {
	if !cmp.Equal(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
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
