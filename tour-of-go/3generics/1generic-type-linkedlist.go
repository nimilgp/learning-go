package main

import "fmt"

type List[T any] struct {
	next *List[T]
	val  T
}

func NewList[T any](val T) *List[T] {
    return &List[T]{val: val}
}

func (l *List[T]) Insert(val T) {
    l.next = &List[T]{val: val, next: l.next}
}

func (l *List[T]) Print() {
    for current := l; current != nil; current = current.next {
        fmt.Println(current.val)
    }
}

func main() {
	l := List{nil,2}	
}
