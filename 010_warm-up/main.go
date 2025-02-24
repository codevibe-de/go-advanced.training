package main

import "fmt"

func main() {
	list := ListItem{value: "initial"}
	list.add("second").add("third")

	list.forEach(func(v string) { fmt.Println("> ", v) })

	var res []string
	list.forEach(func(v string) { res = append(res, v) })
	fmt.Println(res)
}

// --- type ListItem :

type ListItem struct {
	next  *ListItem
	value string
}

func (l *ListItem) add(v string) *ListItem {
	l.next = &ListItem{value: v}
	return l.next
}

func (l *ListItem) forEach(consumer func(v string)) {
	current := l
	for ; current != nil; current = current.next {
		consumer(current.value)
	}
}
