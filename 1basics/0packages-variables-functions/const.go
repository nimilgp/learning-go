package main

import "fmt"

const Pi = 3.14 //for that never changes
// cat use :=

func main() {
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
