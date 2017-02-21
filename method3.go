package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func (n *NamePoint) Abs() float64 {
	return n.Point.Abs() * 100
}

type NamePoint struct {
	Point
	name string
}

func main() {
	n := &NamePoint{Point{3, 4}, "Pythagoras"}
	fmt.Println(n.Abs())
}

//10.6.6
