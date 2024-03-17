package main

import (
	"fmt"
)

type rectangle struct {
	length, breadth int
}

func (rect rectangle) area() int {
	return rect.length * rect.breadth
}

func (rect *rectangle) scale(x int){
	rect.length *= x
	rect.breadth *= x
}

func (rect rectangle) printrect() {
	fmt.Println(rect)
}

func main() {
	rect := rectangle{10, 20}
	rect.scale(10)
	rect.printrect()
}
