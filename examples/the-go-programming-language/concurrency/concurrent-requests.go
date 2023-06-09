package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	imgURLs := []string{
		"https://images.unsplash.com/photo-1516117172878-fd2c41f4a759",
		"https://images.unsplash.com/photo-1532009324734-20a7a5813719",
		"https://images.unsplash.com/photo-1524429656589-6633a470097c",
		"https://images.unsplash.com/photo-1530224264768-7ff8c1789d79",
		"https://images.unsplash.com/photo-1564135624576-c5c88640f235",
		"https://images.unsplash.com/photo-1541698444083-023c97d3f4b6",
		"https://images.unsplash.com/photo-1522364723953-452d3431c267",
		"https://images.unsplash.com/photo-1493976040374-85c8e12f0c0e",
		"https://images.unsplash.com/photo-1504198453319-5ce911bafcde",
		"https://images.unsplash.com/photo-1530122037265-a5f1f91d3b99",
		"https://images.unsplash.com/photo-1516972810927-80185027ca84",
		"https://images.unsplash.com/photo-1550439062-609e1531270e",
		"https://images.unsplash.com/photo-1549692520-acc6669e2f0c",
	}

	start := time.Now()

	var wg sync.WaitGroup

	for _, imgURL := range imgURLs {
		wg.Add(1)
		go func (imgURL string) {
			defer wg.Done()
			resp, err := http.Get(imgURL)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s: %s\n", imgURL, resp.Status)
		}(imgURL)

	}

	wg.Wait()

	fmt.Printf("\nFinished in %v\n", time.Since(start))
}
