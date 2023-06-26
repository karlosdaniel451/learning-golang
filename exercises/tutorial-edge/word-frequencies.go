package main

import (
	"fmt"
	"sort"
	"strings"
)

func CountWords(text string) map[string]int {
	// // Lowercase text and remove dots and commas
	// text = strings.ToLower(text)
	// text = strings.ReplaceAll(text, ",", "")
	// text = strings.ReplaceAll(text, ".", "")

	words := strings.Split(text, " ")
	wordFrequencies := map[string]int{}

	for _, word := range words {
		wordFrequencies[word]++
	}

	return wordFrequencies
}

func topNWords(wordmap map[string]int, n int) ([]WordWithFrequency, error) {
	if n <= 0 || n >= len(wordmap) {
		return nil, fmt.Errorf("error: invalid value for n")
	}

	wordsWithFrequencies := make([]WordWithFrequency, 0, len(wordmap))

	for word, frequency := range wordmap {
		wordsWithFrequencies = append(wordsWithFrequencies,
			WordWithFrequency{word, frequency})
	}

	sort.SliceStable(wordsWithFrequencies, func(i, j int) bool {
		return wordsWithFrequencies[i].frequency > wordsWithFrequencies[j].frequency
	})

	return wordsWithFrequencies[:n], nil
}

func Top5Words(wordmap map[string]int) []WordWithFrequency {
	top5WordsWithFrequencies, err := topNWords(wordmap, 5)
	if err != nil {
		panic(err)
	}

	return top5WordsWithFrequencies
}

type WordWithFrequency struct {
	word      string
	frequency int
}

func main() {
	text := `Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum.`

	wordFrequencies := CountWords(text)

	MostCommon := Top5Words(wordFrequencies)
	fmt.Println(MostCommon)
}
