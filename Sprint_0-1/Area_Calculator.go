package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}
type Rectangle struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (c Circle) Area() float64 {
	return math.Pi * (c.radius * c.radius)
}

func main() {
	figure_1 := Circle{radius: 1.0}
	fmt.Println(figure_1.Area())
	figure_2 := Rectangle{width: 57.2, height: 10.2}
	fmt.Println(figure_2.Area())
}
