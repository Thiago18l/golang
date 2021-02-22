package concurrency

import (
	"testing"
	"time"
)

func slowStubVerifyWebSite(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func benchmarkVerifyWebSite(b *testing.B) {
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "uma url"
	}

	for i := 0; i < b.N; i++ {
		VerificaWebSite(slowStubVerifyWebSite, urls)
	}
}
