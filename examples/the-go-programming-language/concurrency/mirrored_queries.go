package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"http://mirror.pop-sc.rnp.br/mirror/ubuntu-releases/",
		"http://mirror.uepg.br/ubuntu-releases/",
		"http://mirror.pop-sc.rnp.br/ubuntu-releases/",
		"https://ubuntu.itsbrasil.net/ubuntu-releases/",
	}

	fastestResponseTime, err := getFastestResponseTime(urls)
	if err != nil {
		panic(err)
	}

	fmt.Printf("fastestResponseTime: %v\n", fastestResponseTime)
}

// asdfasdfs
func getFastestResponseTime(mirroredUrls []string) (time.Duration, error) {
	type ResponseTimeItem struct {
		duration time.Duration
		err error
	}
	responseTimes := make(chan ResponseTimeItem, len(mirroredUrls))
	for _, url := range mirroredUrls {
		responseTime, err := request(url)
		responseTimes <- ResponseTimeItem{responseTime, err}
	}
	fastestResponseTime := time.Duration(0)
	for responseTime := range responseTimes {
		if responseTime.err != nil {
			return time.Duration(0), responseTime.err
		}
		if fastestResponseTime < responseTime.duration {
			fastestResponseTime = responseTime.duration
		}
	}
	return fastestResponseTime, nil
}

func request(url string) (time.Duration, error) {
	start := time.Now()
	response, err := http.Get(url)
	responseTime := time.Since(start)
	// defer response.Body.Close()
	if err != nil {
		return time.Duration(0), err
	}

	if response.StatusCode != http.StatusOK {
		return time.Duration(0), fmt.Errorf("request failed with status %d", response.StatusCode)
	}
	fmt.Println(responseTime)
	return responseTime, nil
}
