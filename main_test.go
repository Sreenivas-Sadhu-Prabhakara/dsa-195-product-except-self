package main

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	cases := []struct {
		in, want []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
		{[]int{0, 0}, []int{0, 0}}, // two zeros: everything zero
		{[]int{5, 2}, []int{2, 5}},
		{[]int{7}, []int{1}}, // empty product for a single element
		{[]int{}, []int{}},
	}
	for _, c := range cases {
		if got := ProductExceptSelf(c.in); !reflect.DeepEqual(got, c.want) {
			t.Errorf("ProductExceptSelf(%v) = %v, want %v", c.in, got, c.want)
		}
	}
}
