package main

import "fmt"

func split(sum int)(x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return //returns the var in the func ret declaration
}

func main() {
	fmt.Println(split(27))
}

