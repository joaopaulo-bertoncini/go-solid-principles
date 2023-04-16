package main

import "math"

// Interface Segregation - ISP
// — Definition: “Clients should not be forced to depend on methods they don’t use”

type shape interface {
	area() float64
}

type object interface {
	shape
	volume() float64
}

type square struct {
	sideLength float64
}

func (s square) area() float64 {
	return math.Pow(s.sideLength, 2)
}

type cube struct {
	square
}

func (c cube) volume() float64 {
	return math.Pow(c.sideLength, 3)
}

func areaSum(shapes ...shape) float64 {
	var sum float64
	for _, s := range shapes {
		sum += s.area()
	}
	return sum
}

func areaVolumeSum(shapes ...object) float64 {
	var sum float64
	for _, s := range shapes {
		sum += s.area() + s.volume()
	}

	return sum
}

func main() {
	s := square{sideLength: 2}
	c := cube{square{sideLength: 2}}
	println(areaSum(s, c))
	println(areaVolumeSum(c))
}
