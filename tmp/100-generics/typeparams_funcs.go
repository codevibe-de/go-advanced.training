package _100_generics

func foo[T any](v T) T {
	return v
}

func isEqual[T comparable](v1 T, v2 T) bool {
	return v1 == v2
}
