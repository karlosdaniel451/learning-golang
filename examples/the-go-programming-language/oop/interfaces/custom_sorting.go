package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func main() {
	var tracks = []*Track{
		{
			Title:  "Hemisphère",
			Artist: "Paradis",
			Album:  "Couleurs Primaires",
			Year:   2015,
			Length: GetLength("7m14s")},
		{
			Title:  "New Person, Same Old Mistakes",
			Artist: "Tame Impala",
			Album:  "Currents",
			Year:   2015,
			Length: GetLength("6m04s")},
		{
			Title:  "Thinking of A Place",
			Artist: "The War On Drugs",
			Album:  "A Deeper Understanding",
			Year:   2017,
			Length: GetLength("11m11s"),
		},
		{
			Title:  "How",
			Artist: "The Neighbourhood",
			Album:  "I Love You",
			Year:   2013,
			Length: GetLength("5m15s"),
		},
	}

	// Sort by title:
	sort.Slice(tracks, func (i, j int) bool {
		return tracks[i].Title < tracks[j].Title
	})
	fmt.Println("Sorted by title:\n")
	printTracks(tracks)

	// Sort by year:
	sort.Slice(tracks, func(i, j int) bool {
		return tracks[i].Year < tracks[j].Year
	})
	fmt.Println("\n\nSorted by year:\n")
	printTracks(tracks)

	// Sort by length:
	sort.Slice(tracks, func(i, j int) bool {
		return tracks[i].Length < tracks[j].Length
	})
	fmt.Println("\n\nSorted by length:\n")
	printTracks(tracks)
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

func GetLength(s string) time.Duration {
	duration, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}

	return duration
}
