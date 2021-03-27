package main

import (
	"fmt"
	"net/http"
)

// StoragePlayer is a storage to save the points of players in memory
type StoragePlayer interface {
	GetPlayerPoints(name string) int
}

// PlayerServer is an interface http for the data of players
type PlayerServer struct {
	storage StoragePlayer
}

func (p *PlayerServer) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	points := p.storage.GetPlayerPoints(player)
	if points == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, points)
}
