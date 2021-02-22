package concurrency

import (
	"reflect"
	"testing"
)

func mockVerifyWebSite(url string) bool {
	if url == "waat://fuhur.geds" {
		return false
	}
	return true
}

func TestVerifyWebSite(t *testing.T) {
	websites := []string{
		"http://google.com",
		"https://github.com",
		"waat://fuhur.geds",
	}

	waited := map[string]bool{
		"http://google.com":  true,
		"https://github.com": true,
		"waat://fuhur.geds":  false,
	}

	result := VerificaWebSite(mockVerifyWebSite, websites)

	if !reflect.DeepEqual(waited, result) {
		t.Fatalf("waited '%v', result '%v'", waited, result)
	}
}
