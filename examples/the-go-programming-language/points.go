package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
}

func (p *Point) DistanceFromOrigin() float64 {
	return p.DistanceFromPoint(&Point{0, 0})
}

func (p *Point) DistanceFromPoint(q *Point) float64 {
	return math.Sqrt(math.Pow(p.X-q.X, 2) + math.Pow(p.Y-q.Y, 2))
}

func (p *Point) String() string {
	return fmt.Sprintf("X: %f, Y: %f", p.X, p.Y)
}

func main() {
	pointA := Point{1, 1}
	pointB := Point{2, 2}

	fmt.Println(pointA.DistanceFromOrigin())
	fmt.Println(pointB.DistanceFromOrigin())

	fmt.Println(pointA.DistanceFromPoint(&pointB))
	fmt.Println(pointB.DistanceFromPoint(&pointA))

	p := new(Point)
	fmt.Printf("%#v\n", *p)
	fmt.Printf("%s\n", *p)
	fmt.Printf("%s\n", p)
	fmt.Println(*p)
	fmt.Println(p)
}
