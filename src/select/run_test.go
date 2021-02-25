package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRunner(t *testing.T) {
	serverSlow := createServerWithDelay(20 * time.Millisecond)
	serverFast := createServerWithDelay(0 * time.Millisecond)

	URLSlow := serverSlow.URL
	URLFast := serverFast.URL

	waited := URLFast
	result := Runner(URLSlow, URLFast)

	if result != waited {
		t.Errorf("result '%s', waited '%s'", result, waited)
	}
}

func createServerWithDelay(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		rw.WriteHeader(http.StatusOK)
	}))
}
