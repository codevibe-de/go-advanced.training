package main

import (
	"math"
	"testing"
)

func TestPrimeNumbers(t *testing.T) {
	tests := []struct {
		max    int
		primes []int
	}{
		{3, []int{2}},
		{10, []int{2, 3, 5, 7}},
		{100, []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}},
	}
	for _, test := range tests {
		primes := primeNumbers(test.max)
		if len(primes) != len(test.primes) {
			t.Errorf("primeNumbers(%d) returned %v, wanted %v", test.max, primes, test.primes)
		}
		for i, p := range primes {
			if p != test.primes[i] {
				t.Errorf("primeNumbers(%d) returned %v, wanted %v", test.max, primes, test.primes)
			}
		}
	}
}

func BenchmarkPrimeNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primeNumbers(i)
	}
}

// primeNumbers calculate the prime numbers up to max and returns them as a slice.
func primeNumbers(max int) []int {
	var primes []int
	for i := 2; i < max; i++ {
		isPrime := true
		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
		}
	}
	return primes
}
