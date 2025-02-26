package fib_recursive

func Get(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Get(n-1) + Get(n-2)
}
