package _100_generics

type Box[T any] struct{ value T }

func (b Box[T]) get() T {
	return b.value
}

func main() {
	box := Box[int]{value: 42}
	var n int = box.get()
	_ = n
}
