package main

import "fmt"

type myint int64

type myfloat float64

type mytype interface{}

func main() {
	var i mytype
	
	s := i
	fmt.Println(s)

	s,ok := i.(float64)
	fmt.Println(s, ok)
}
