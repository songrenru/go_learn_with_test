package concurrency

// import "time"

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	// time.Sleep(2 * time.Second)
	len := len(urls)
	for i := 0; i < len; i++ {
		result := <-resultChannel
		results[result.string] = result.bool
	}

	return results
}