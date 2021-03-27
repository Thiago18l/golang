package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPlayers(t *testing.T) {
	t.Run("return the result from thiago", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/thiago", nil)
		response := httptest.NewRecorder()

		PlayerServer(response, request)

		received := response.Body.String()
		waited := "20"

		if received != waited {
			t.Errorf("received '%s', but expect '%s'", received, waited)
		}
	})
}
