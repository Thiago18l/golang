package main

import (
	"net/http"
	"time"
)

// Runner verify what url is most fast
func Runner(a, b string) (woner string) {
	durationA := timeResponse(a)
	durationB := timeResponse(b)

	if durationA < durationB {
		return a
	}
	return b
}

func timeResponse(url string) time.Duration {
	begin := time.Now()
	http.Get(url)
	return time.Since(begin)
}

func main() {

}
