package main

import (
	"fmt"
	"net/http"
)

// StoragePlayer is a storage to save the points of players in memory
type StoragePlayer interface {
	GetPlayerPoints(name string) int
	RegisterOwn(name string)
}

// PlayerServer is an interface http for the data of players
type PlayerServer struct {
	storage StoragePlayer
}

func (p *PlayerServer) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	switch r.Method {
	case http.MethodPost:
		p.registerOwn(w, player)
	case http.MethodGet:
		p.showPoints(w, player)
	}
}

func (p *PlayerServer) showPoints(w http.ResponseWriter, player string) {

	points := p.storage.GetPlayerPoints(player)

	if points == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, points)
}

func (p *PlayerServer) registerOwn(w http.ResponseWriter, name string) {

	p.storage.RegisterOwn(name)
	w.WriteHeader(http.StatusAccepted)
}
