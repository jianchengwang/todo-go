package main

import "fmt"

type rect struct {
	with, high int
}

func (r *rect) area() int {
	r.with = 100
	return r.with * r.high
}

func (r rect) perim() int {
	r.with = 1000
	return 2*(r.with + r.high)
}

func main() {
	r := rect{10, 10}
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}


