package _100_generics

func reverse[T any](s []T) []T {
	n := len(s)
	for i := 0; i < n/2; i++ {
		s[i], s[n-i-1] = s[n-i-1], s[i]
	}
	return s
}
