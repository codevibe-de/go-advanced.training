package main

import "fmt"

func main() {
	set := NewSet[int]()
	set.Add(1)
	set.Add(23)
	set.Add(23)
	fmt.Println(set.Len())
	fmt.Println(set.String())
	set.Iter()(func(value int) bool {
		fmt.Println(value)
		return true
	})
}
