package main

import "fmt"

func swap(str1, str2 string)(string,string) {
	return str2, str1
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
