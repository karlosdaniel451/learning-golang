package main

import (
	"encoding/json"
	"fmt"
)

type Stock struct {
	Ticker   string  `json:"ticker"`
	Dividend float64 `json:"dividend"`
}

// Iterates through a JSON string of stocks unmarshalls them into a struct
// and returns the ticket of stock with the highest dividend
func HighestDividend(stocksJson string) (stockTicket string) {
	stocks := []Stock{}
	err := json.Unmarshal([]byte(stocksJson), &stocks)
	if err != nil {
		panic(err)
	}

	stockWithHighestDividend := Stock{}

	for _, stock := range stocks {
		if stock.Dividend > stockWithHighestDividend.Dividend {
			stockWithHighestDividend = stock
		}
	}

	return stockWithHighestDividend.Ticker
}

func main() {
	stocksJson := `[
    {"ticker":"APPL","dividend":0.5},
    {"ticker":"GOOG","dividend":0.2},
    {"ticker":"MSFT","dividend":0.3}
  ]`

	highestDividend := HighestDividend(stocksJson)
	fmt.Println(highestDividend)
}
