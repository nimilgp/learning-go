package main

import (
	"fmt"
	"math"
)

type shape interface {
	area()
}

type rectangle struct {
	length, breadth float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() {
	fmt.Println(r.length * r.breadth)
}

func (c circle) area() {
	fmt.Println(math.Pi * c.radius * c.radius)
}

func calc_area(s shape) {
	s.area()
}

func main() {
	rect := rectangle{10,20}
	circ := circle{10}
	calc_area(rect)
	calc_area(circ)
}
