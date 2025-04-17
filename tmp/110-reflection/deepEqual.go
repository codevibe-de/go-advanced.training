package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int16 = 1
	var b int8 = 1
	fmt.Println(reflect.DeepEqual(a, b)) // false

	slc1 := []string{"a", "b", "c"}
	slc2 := []string{"a", "b", "c"}
	fmt.Println(reflect.DeepEqual(slc1, slc2)) // true
}
