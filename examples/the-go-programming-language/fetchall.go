package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func example_fetchall_main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Not enough CLI args\n")
		os.Exit(1)
	}

	var urls = os.Args[1:]
	ch := make(chan string)

	for _, url := range urls {
		go fetch(url, ch)
	}

	for _, url := range urls {
		fmt.Printf("\nResponse for request at \"%s\":\n%s\n", url, <-ch)
	}
}

func fetchAll(url string, ch chan<- string) {
	fmt.Printf("\nSending GET request at \"%s\"\n", url)
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	statusCode := resp.StatusCode

	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()

	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	ch <- fmt.Sprintf("%d\n%s", statusCode, body)
}
