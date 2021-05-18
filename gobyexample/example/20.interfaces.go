package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rect1 struct {
	with, high float64
}
type circle struct {
	radius float64
}

func (r rect1) area() float64 {
	return r.with * r.high
}
func (r rect1) perim() float64 {
	return 2*(r.with + r.high)
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main()  {
	r := rect1{10, 20}
	c := circle{10}
	measure(r)
	measure(c)
}


