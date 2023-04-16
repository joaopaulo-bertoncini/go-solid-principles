package main

// Open/Close principle:

import (
	"fmt"
	"math"
)

type shape interface {
	area() float32
}

type square struct {
	sideLength float32
}

func (s square) area() float32 {
	return s.sideLength * s.sideLength
}

type circle struct {
	radius float32
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

type triangle struct {
	base   float32
	height float32
}

func (t triangle) area() float32 {
	return t.base * t.height * 0.5
}

type calculator struct {
	total float32
}

func (c calculator) sunAreas(shapes ...shape) float32 {
	var sum float32
	for _, shape := range shapes {
		sum += shape.area()
	}

	return sum
}

func main() {
	s := square{sideLength: 5}
	c := circle{radius: 5}
	t := triangle{base: 5, height: 5}

	calculator := calculator{}
	fmt.Println(calculator.sunAreas(s, c, t))
}
