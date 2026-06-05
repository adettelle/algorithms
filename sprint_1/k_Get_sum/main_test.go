package main

import (
	"reflect"
	"testing"
)

func TestFactorize(t *testing.T) {
	tests := []struct {
		name     string
		lenX     int
		x        []int
		k        int
		expected []int
	}{
		{
			"Test 1",
			4,
			[]int{1, 2, 0, 0},
			34,
			[]int{1, 2, 3, 4},
		},
		{
			"Test 2",
			2,
			[]int{9, 5},
			17,
			[]int{1, 1, 2},
		},
		{
			"Test 3",
			30,
			[]int{2, 0, 4, 3, 3, 2, 1, 8, 1, 9, 6, 0, 0, 1, 3, 3, 8, 9, 0, 8, 3, 8, 6, 3, 7, 9, 4, 0, 2, 6},
			349,
			[]int{2, 0, 4, 3, 3, 2, 1, 8, 1, 9, 6, 0, 0, 1, 3, 3, 8, 9, 0, 8, 3, 8, 6, 3, 7, 9, 4, 3, 7, 5},
		},
		{
			"Test 4",
			1,
			[]int{3},
			7991,
			[]int{7, 9, 9, 4},
		},
	}

	for _, tt := range tests {
		var x string
		res := getSum(tt.x, tt.k)

		if !reflect.DeepEqual(res, tt.expected) {
			t.Errorf("%s: expected result: %v, result: %v", tt.name, tt.expected, x)
		}
	}
}
