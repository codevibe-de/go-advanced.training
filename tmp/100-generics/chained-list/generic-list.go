package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	list := List[string]{value: "initial"}
	list.add("second").add("third")
	list.forEach(func(v string) { fmt.Println("> ", v) })
	var res []string
	list.forEach(func(v string) { res = append(res, v) })
	fmt.Println(res)

	listPtr := &list
	listPtr = nil
	listPtr.add("fourth")

	// ---
	var rw io.ReadWriter = &bytes.Buffer{}

	listOfReaders := List[io.Reader]{}
	listOfReaders.add(rw)

	listOfReadWriters := List[io.ReadWriter]{}
	listOfReadWriters.add(rw)

	//listOfReaders = listOfReadWriters
}

// simple chained list

type List[T any] struct {
	next  *List[T]
	value T
}

func (l *List[T]) add(v T) *List[T] {
	if l != nil {
		l.next = &List[T]{value: v}
		return l.next
	} else {
		return nil
	}
}

func (l *List[T]) forEach(consumer func(v T)) {
	current := l
	for ; current != nil; current = current.next {
		consumer(current.value)
	}
}
