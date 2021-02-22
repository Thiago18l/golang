package concurrency

// VerifyWebSite verify an url, and return a boolean
type VerifyWebSite func(string) bool

type result struct {
	string
	bool
}

// VerificaWebSite will receive a type VerifyWebSite and a slice
func VerificaWebSite(vw VerifyWebSite, urls []string) map[string]bool {
	results := make(map[string]bool)
	pipeResult := make(chan result)

	for _, url := range urls {
		go func(u string) {
			pipeResult <- result{u, vw(u)}
		}(url)
	}
	for i := 0; i < len(urls); i++ {
		result := <-pipeResult
		results[result.string] = result.bool
	}
	return results
}
