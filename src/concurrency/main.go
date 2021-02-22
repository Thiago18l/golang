package concurrency

// VerifyWebSite verify an url, and return a boolean
type VerifyWebSite func(string) bool

// VerificaWebSite will receive a type VerifyWebSite and a slice
func VerificaWebSite(vw VerifyWebSite, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = vw(url)
	}
	return results
}
