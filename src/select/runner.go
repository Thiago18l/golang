package main

import (
	"fmt"
	"net/http"
	"time"
)

var timeMax = 10 * time.Second

// Runner verify what url is most fast
func Runner(a, b string, timeLimit time.Duration) (woner string, err error) {
	return Config(a, b, timeLimit)
}

// Config will compare the time of response of url's
func Config(a, b string, timeLimit time.Duration) (woner string, err error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeMax):
		return "", fmt.Errorf("time limit of wait for '%s', '%s'", a, b)
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}

func main() {

}
