package main

import (
	"log"
	"net/http"
)

func main() {
	treater := http.HandlerFunc(PlayerServer)
	if err := http.ListenAndServe(":8080", treater); err != nil {
		log.Fatalf("Dont was possible to run on port 8080, %v", err)
	}
}
