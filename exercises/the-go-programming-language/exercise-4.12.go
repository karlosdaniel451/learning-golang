package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type XKCD struct {
	Month      string
	Num        int
	Link       string
	Year       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Img        string
	Title      string
	Day        string
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("XKCD number: ")
	xkcdNumberInString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	xkcdNumberInString = strings.TrimSpace(xkcdNumberInString)
	xkcdNumber, err := strconv.Atoi(xkcdNumberInString)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	xkcd, err := GetXKCDByNumber(xkcdNumber)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("Ttitle: %s\nDate: %s-%s-%s\nTranscript:\n%s\n\nAlt: %s\nImg: %s\n",
		xkcd.Title, xkcd.Year, xkcd.Month, xkcd.Day, xkcd.Transcript, xkcd.Alt, xkcd.Img)
}

// GetXKCDByNumber retrieves information about an XKCD comic by its number.
// The function returns a pointer to an XKCD struct and an error. If the
// function encounters an error, it will return nil for the XKCD struct and
// the error message. Otherwise, it will return a pointer to the XKCD struct
// containing the comic information and nil for the error message.
func GetXKCDByNumber(number int) (*XKCD, error) {
	requestURL := fmt.Sprintf("https://xkcd.com/%d/info.0.json", number)

	response, err := http.Get(requestURL)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed with status %d", response.StatusCode)
	}

	var XKCD XKCD

	err = json.NewDecoder(response.Body).Decode(&XKCD)
	if err != nil {
		return nil, err
	}

	return &XKCD, nil
}
