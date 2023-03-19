package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func example_fetch_main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Not enough CLI args\n")
		os.Exit(1)
	}

	var urls = os.Args[1:]

	for _, url := range urls {
		fmt.Printf("\nSending GET request at \"%s\"\n", url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error by sending request: %v\n", err)
			break
		}

		fmt.Printf("Status code: %d\n", resp.StatusCode)

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error by reading response body: %v\n", err)
			break
		}

		fmt.Printf("Response: %s\n", body)
	}
}
