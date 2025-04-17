package main

import (
	"fmt"
	"local/040_testing/stack"
)

func ExampleStack_Push() {
	st := stack.NewStack()
	st.Push("foo")
	fmt.Printf("%q\n", *st)
	// Output: {["foo"]}
}
