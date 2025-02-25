package main

import (
	"fmt"
	"iter"
)

type Set[T comparable] struct {
	values map[T]struct{}
}

func NewSet[T comparable]() *Set[T] {
	return &Set[T]{values: make(map[T]struct{})}
}

func (s *Set[T]) Add(value T) {
	s.values[value] = struct{}{}
}

func (s *Set[T]) Remove(value T) {
	delete(s.values, value)
}

func (s *Set[T]) Contains(value T) bool {
	_, exists := s.values[value]
	return exists
}

func (s *Set[T]) Len() int {
	return len(s.values)
}

func (s *Set[T]) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Set[T]) String() string {
	tmp := make([]T, 0, s.Len())
	s.Iter()(func(value T) bool {
		tmp = append(tmp, value)
		return true
	})
	return fmt.Sprintf("%v", tmp)
}

func (s *Set[T]) Iter() iter.Seq[T] {
	return func(yield func(T) bool) {
		for value := range s.values {
			if !yield(value) {
				break
			}
		}
	}
}
