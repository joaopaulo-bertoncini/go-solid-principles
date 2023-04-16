package main

//Single Responsibility Principle - SRP
// — Definition: “A module should have one, and only one reason to change”

import (
	"fmt"
	"math"
)

type circle struct {
	radius float32
}

func (c circle) area() float32 {
	return math.Pi * c.radius * c.radius
}

type square struct {
	side float32
}

func (s square) area() float32 {
	return s.side * s.side
}

type shape interface {
	area() float32
}

type outPrinter struct{}

func (o outPrinter) print(s shape) string {
	return fmt.Sprintf("the area is: %f", s.area())
}

func main() {
	c := circle{radius: 5}
	s := square{side: 2}
	o := outPrinter{}
	fmt.Println(o.print(c))
	fmt.Println(o.print(s))
}
