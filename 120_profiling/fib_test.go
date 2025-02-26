package main

import (
	"fmt"
	"local/_120_profiling/fib_fast"
	"local/_120_profiling/fib_memwaste"
	"local/_120_profiling/fib_recursive"
	"testing"
)

// performTest is a helper function to test the performance of the Fibonacci function.
// It checks if the result for the 30th Fibonacci number is correct (repeatedly for "count" times).
// The actual Fibonacci function to call is passed as the fib parameter.
func performTest(t testing.TB, count int, fib func(int) int) {
	for i := 0; i < count; i++ {
		if fib(30) != 832040 {
			t.Error("Incorrect!")
		}
	}
}

func TestRecursive(t *testing.T) {
	performTest(t, 1_000, fib_recursive.Get)
}

func BenchmarkRecursive(b *testing.B) {
	fmt.Println("BenchmarkRecursive")
	performTest(b, b.N, fib_recursive.Get)
}

func BenchmarkMemwaste(b *testing.B) {
	performTest(b, b.N, fib_memwaste.Get)
}

func BenchmarkFast(b *testing.B) {
	performTest(b, b.N, fib_fast.Get)
}
