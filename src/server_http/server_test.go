package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type PlayerStorage struct {
	points map[string]int
}

func (p *PlayerStorage) GetPlayerPoints(name string) int {
	points := p.points[name]
	return points
}

func TestGetPlayers(t *testing.T) {

	storage := PlayerStorage{
		map[string]int{
			"thiago": 20,
			"pedro":  10,
		},
	}
	server := &PlayerServer{&storage}
	t.Run("return the result from thiago", func(t *testing.T) {

		request := newRequisitionPoints("thiago")
		response := httptest.NewRecorder()

		server.ServerHTTP(response, request)

		received := response.Body.String()
		waited := "20"
		verifyStatusCode(t, response.Code, http.StatusOK)
		verifyBodyRequest(t, received, waited)
	})

	t.Run("return the result from pedro", func(t *testing.T) {
		request := newRequisitionPoints("pedro")
		response := httptest.NewRecorder()

		server.ServerHTTP(response, request)

		received := response.Body.String()
		waited := "10"
		verifyStatusCode(t, response.Code, http.StatusOK)
		verifyBodyRequest(t, received, waited)
	})

	t.Run("should return 404 if the player is not found", func(t *testing.T) {
		request := newRequisitionPoints("John")
		response := httptest.NewRecorder()

		server.ServerHTTP(response, request)

		received := response.Code
		waited := http.StatusNotFound

		verifyStatusCode(t, received, waited)
	})
}

func newRequisitionPoints(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return request
}

func verifyBodyRequest(t *testing.T, received, waited string) {
	t.Helper()
	if received != waited {
		t.Errorf(`received '%s', but expect '%s'`, received, waited)
	}
}

func verifyStatusCode(t *testing.T, received, waited int) {
	t.Helper()
	if received != waited {
		t.Errorf("received %d, waited %d", received, waited)
	}
}
