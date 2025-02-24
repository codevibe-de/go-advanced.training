package main

import "fmt"

func main() {
	go add(1, 2)

	go func() {
		fmt.Println("Hello from your Goroutine!")
	}()
	// Output:
}

func add(a, b int) int {
	return a + b
}
