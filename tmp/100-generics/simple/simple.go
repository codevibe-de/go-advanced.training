package main

import (
	"cmp"
	"fmt"
)

// Go has a built-in function called "max" since 1.21
func getMax[T int8](v ...T) T {
	if len(v) == 0 {
		return 0
	}
	res := v[0]
	for _, val := range v {
		if val > res {
			res = val
		}
	}
	return res
}

func getMax2[T cmp.Ordered](v0 T, v ...T) T {
	res := v0
	for _, val := range v {
		if val > res {
			res = val
		}
	}

	return res
}

func main() {
	fmt.Println(getMax2(1, 2, -3, 7))
}
