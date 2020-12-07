package checkWebsites

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resulteChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			//results[u] = wc(u)
			resulteChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		result := <-resulteChannel
		results[result.string] = result.bool
	}

	//time.Sleep(2 * time.Second)

	return results
}
