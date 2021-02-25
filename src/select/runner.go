package main

import (
	"net/http"
	"time"
)

// Runner verify what url is most fast
func Runner(a, b string) (woner string) {
	startA := time.Now()
	http.Get(a)
	durationA := time.Since(startA)

	startB := time.Now()
	http.Get(b)
	durationB := time.Since(startB)

	if durationA < durationB {
		return a
	}
	return b
}

func main() {

}
