package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int = 0

type Response struct {
	Count int `json:"count"`
}

func example_server2_main() {
	http.HandleFunc("/count", func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		count++
		mu.Unlock()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{
			Count: count,
		})
	})

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
