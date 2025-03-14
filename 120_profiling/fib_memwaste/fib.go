package fib_memwaste

func Get(n int) int {
	f := make([]int, 10_000)
	f[0] = 0
	f[1] = 1
	for i := 2; i <= n; i++ {
		f[i] = f[i-1] + f[i-2]
	}
	return f[n]
}
