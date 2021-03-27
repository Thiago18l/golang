package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterOwnsAndFindOwns(t *testing.T) {
	storage := newStoragePlayerInMemory()
	server := PlayerServer{storage}

	player := "thiago"
	for i := 0; i < 3; i++ {
		server.ServerHTTP(httptest.NewRecorder(), newRequisitionRegister(player))
	}

	response := httptest.NewRecorder()
	server.ServerHTTP(response, newRequisitionPoints(player))
	verifyStatusCode(t, response.Code, http.StatusOK)
	verifyBodyRequest(t, response.Body.String(), "3")
}
