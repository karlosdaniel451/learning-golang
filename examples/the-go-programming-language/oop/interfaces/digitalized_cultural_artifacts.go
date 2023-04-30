package main

import (
	"fmt"
	"io"
	"time"
)

type Artifact interface {
	Title() string
	Creators() []string
	Created() time.Time
}

type Text interface {
	Pages() int
	Words() int
	PageSize() int
}

type Streamer interface {
	Stream() (io.ReadCloser, error)
	RunningTime() int
	Format() string
}

type Audio interface {
	Stream() Streamer
}

type Video interface {
	Stream() Streamer
}

type Book struct {
	title string
}

type Movie struct {
}

type Podcast struct {
}

type TVEpisode struct {
}

type Track struct {
}

func (book *Book) Title() string {
	return book.title
}

func main() {
	fmt.Println("Hello, world")
}
