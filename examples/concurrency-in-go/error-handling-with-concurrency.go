package main

import (
	"fmt"
	"net/http"
	"sync"
)

type Result struct {
	Req  *http.Request
	Resp *http.Response
	Err  error
}

func requestUrls(done <-chan struct{}, urls ...string) <-chan Result {
	results := make(chan Result, len(urls))
	var wg sync.WaitGroup

	defer func() {
		// Check if goroutine was cancelled.
		select {
		case <-done:
			close(results)
			return
		default:
		}

		// If gouroutine was not cancelled, wait for all responses.
		wg.Wait()
		close(results)
	}()

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()

			// Check for cancellation request in the `done` channel.
			select {
			case <-done:
				return
			default:
			}

			fmt.Printf("requesting %s\n", url)
			req, err := http.NewRequest(http.MethodGet, url, nil)
			resp, err := http.DefaultClient.Do(req)
			results <- Result{Req: req, Resp: resp, Err: err}
		}(url)
	}

	return results
}

func main() {
	urls := []string{
		"https://www.google.com.br",
		"https://www.a.com",
		"https://go.dev",
		"https://github.com",
		"https://grpc.io",
		"https://gofiber.io/",
	}

	done := make(chan struct{})

	i := 1
	for result := range requestUrls(done, urls...) {
		if i == 3 {
			fmt.Printf("cancelling goroutine after %d responses received\n", i)
			close(done)
			break
		}
		if result.Err != nil {
			fmt.Printf("error from %s: %v\n", result.Req.URL, result.Err)
			continue
		}
		fmt.Printf("response from %s: %s\n", result.Req.URL, result.Resp.Status)
		i++
	}
}
