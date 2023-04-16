package main

import "fmt"

type transport interface {
	getName() string
}

type vehicle struct {
	name string
}

func (v *vehicle) getName() string {
	return v.name
}

type car struct {
	vehicle
	wheel int
	gates int
}

type motocycle struct {
	vehicle
	wheel int
}

type printer struct{}

func (p *printer) print(t transport) {
	fmt.Printf("Name: %s\n", t.getName())
}

func main() {
	v := vehicle{name: "vehicle"}

	c := car{
		vehicle: vehicle{name: "car"},
		wheel:   4,
		gates:   2,
	}

	m := motocycle{
		vehicle: vehicle{name: "motocycle"},
		wheel:   2,
	}

	p := printer{}
	p.print(&v)
	p.print(&c)
	p.print(&m)
}
