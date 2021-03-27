package main

import (
	"log"
	"net/http"
)

// PlayerStorageInMemory store in memory the data about the players
type PlayerStorageInMemory struct{}

// GetPlayerPoints get the points from a player
func (p *PlayerStorageInMemory) GetPlayerPoints(name string) int {
	return 123
}

func main() {
	server := &PlayerServer{&PlayerStorageInMemory{}}
	treater := http.HandlerFunc(server.ServerHTTP)
	if err := http.ListenAndServe(":8080", treater); err != nil {
		log.Fatalf("Dont was possible to run on port 8080, %v", err)
	}
}
