package _100_generics

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
	l.next = &List[T]{value: v}
	return l.next
}

func (l *List[T]) forEach(consumer func(v T)) {
	current := l
	for ; current != nil; current = current.next {
		consumer(current.value)
	}
}
