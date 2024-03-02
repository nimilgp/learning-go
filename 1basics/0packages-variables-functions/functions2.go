package main

import "fmt"

func add(x, y int)/*both x and y is int*/ int { 
	return x + y
}

func main() {
	fmt.Println(add(5, 7))
}
