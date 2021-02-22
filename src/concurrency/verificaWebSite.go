package concurrency

import "net/http"

// VerifyWebSite will return a bool if the url return a status code 200 'OK'
func VerifyWebSite(url string) bool {
	res, err := http.Head(url)
	if err != nil {
		return false
	}

	if res.StatusCode != http.StatusOK {
		return false
	}
	return true
}
