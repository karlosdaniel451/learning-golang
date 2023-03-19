/*
Exercise 1.8: Mo diy fetch to add the prefix http:// to each argument URL if it is missing.
You might want to use strings.HasPrefix.
*/

package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func exercise_1_8_main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Not enough CLI args\n")
		os.Exit(1)
	}

	var urls = os.Args[1:]

	for _, url := range urls {
		if strings.HasPrefix(url, "http://") == false {
			url = "http://" + url
		}

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
