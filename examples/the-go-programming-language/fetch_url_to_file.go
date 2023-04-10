package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

// Make an HTTP request at the given url and append its content to the given file.
// Return the number of written bytes and an error.
func fetchURLToFile(url, filename string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("HTTP response status: %s", resp.Status)
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	n, err := f.Write(content)
	if err != nil {
		return 0, nil
	}
	return n, nil
}

// var url = flag.String("", "", "HTTP URL to make the request")
var filename = flag.String("f", "", "filename to append the response content")

func main() {
	flag.Parse()

	// Check if the URL argument was passed.
	if len(flag.Args()) < 1 {
		log.Fatal("empty url")
	}

	url := flag.Args()[0]

	_, err := fetchURLToFile(url, *filename)
	if err != nil {
		log.Fatal(err)
	}
}
