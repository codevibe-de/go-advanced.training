package main

import "iter"

// All is a simplified version of the new slices.All() function (Go 1.23)
func All(s []string) iter.Seq2[int, string] {
	return func(yield func(int, string) bool) {
		for i, v := range s {
			if !yield(i, v) {
				return
			}
		}
	}
}

func main() {
	slc := []string{"a", "b", "c"}
	iterator := All(slc)

	// direkte nutzung:
	callback := func(i int, v string) bool {
		println(i, v)
		return true
	}
	iterator(callback)

	// alternativ:
	for i, v := range iterator {
		println(i, v)
	}
}

// Output:
//0 a
//1 b
//2 c
