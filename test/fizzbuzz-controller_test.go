package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/elzinko/go-fizzbuzz/internal/api"
	"github.com/google/go-cmp/cmp"
)

func Test_FizzBuzz_Controller_OK(t *testing.T) {

	web := api.SetupServer()
	ts := httptest.NewServer(web)
	defer ts.Close()

	resp, err := http.Get(fmt.Sprintf("%s/api/fizzbuzz", ts.URL))

	if err != nil {
		fmt.Printf("%s", err)
	}

	defer resp.Body.Close()

	if !cmp.Equal(resp.StatusCode, http.StatusNotImplemented) {
		t.Errorf("got %q, wanted %q", resp.StatusCode, http.StatusOK)
	}
}
