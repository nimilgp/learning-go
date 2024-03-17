package main

import (
	"fmt"
)

type rectangle struct {
	length, breadth float64
}

func (rect rectangle) area() float64 {
	return rect.length * rect.breadth
}

func main() {
	rect := rectangle{10, 20}
	fmt.Println(rect.area())
}
