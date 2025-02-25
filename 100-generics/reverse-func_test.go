package main

import (
	"reflect"
	"testing"
)

func Test_reverse(t *testing.T) {
	type testCase[T any] struct {
		name string
		slc  []T
		want []T
	}
	tests := []testCase[int]{
		{"single element", []int{1}, []int{1}},
		{"two elements", []int{12, 23}, []int{23, 12}},
		{"five elements including negative", []int{1, -2, 3, -4, 5}, []int{5, -4, 3, -2, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.slc); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
