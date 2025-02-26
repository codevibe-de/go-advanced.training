package fib_fast

var f = [9999]int{}

func init() {
	f[0] = 0
	f[1] = 1
	for i := 2; i <= len(f); i++ {
		f[i] = f[i-1] + f[i-2]
	}
}

func Get(n int) int {
	return f[n]
}
