package main

import (
	"fmt"
	"net/http"
)

type Result struct {
	Error    error
	Response *http.Response
}

func main() {
	checkStatus := func(done <-chan interface{}, urls ...string) <-chan Result {
		results := make(chan Result)
		go func() {
			defer close(results)
			for _, url := range urls {
				resp, err := http.Get(url)
				result := Result{Error: err, Response: resp}
				select {
				case <-done:
					return
				case results <- result:
				}
			}
		}()
		return results
	}
	done := make(chan interface{})
	defer close(done)

	urls := []string{"https://google.com", "https://badtogatoga", "https://badtogatoga", "https://badtogatoga", "https://badtogatoga", "https://badtogatoga", "https://badtogatoga", "https://badtogatoga", "https://badtogatoga", "https://badtogatoga", "https://badtogatoga", "https://badtogatoga", "https://badtogatoga", "https://badtogatoga", "https://badtogatoga", "https://badtogatoga", "https://badtogatoga", "https://badtogatoga", "https://badtogatoga", "https://badtogatoga", "https://badtogatoga"}
	for result := range checkStatus(done, urls...) {
		if result.Error != nil {
			fmt.Printf("error: %v", result.Error)
			continue
		}
		fmt.Printf("Response: %v\n", result.Response)
	}
}
