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

	defer serverSlow.Close()
	defer serverFast.Close()

	URLSlow := serverSlow.URL
	URLFast := serverFast.URL

	waited := URLFast
	result, err := Runner(URLSlow, URLFast, 0*time.Millisecond)
	if err != nil {
		t.Fatalf("waited an error, but collect: '%v'", err)
	}
	if result != waited {
		t.Errorf("result '%s', waited '%s'", result, waited)
	}

	t.Run("return an error if the server don't answer in ten seconds", func(t *testing.T) {
		server := createServerWithDelay(11 * time.Second)

		defer server.Close()

		_, err := Config(server.URL, server.URL, 20*time.Second)
		if err == nil {
			t.Error("was waiting a error")
		}
	})
}

func createServerWithDelay(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		rw.WriteHeader(http.StatusOK)
	}))
}
