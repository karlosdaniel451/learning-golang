/*
Exercise 1.7: The function call io.Copy(dst, src) reads from src and writes to dst.
Use it instead of ioutil.ReadAll to copy the response body to os.Stdout without
requiring a buffer large enough to hold the entire stream. Be sure to check the error
result of io.Copy.
*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func exercise_1_7_main() {
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

		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error by reading response body: %v\n", err)
			break
		}

	}
}
