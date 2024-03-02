package main

import "fmt"

func add(x int, y int) int { //return type is declared after parameters
	return x + y
}

func main() {
	fmt.Println(add(5, 7))
}
