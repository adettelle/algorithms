package main

import (
	"reflect"
	"testing"
)

func TestNearestZero(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		arr      []int
		expected []int
	}{
		{
			"Two zeros",
			5,
			[]int{0, 1, 4, 9, 0}, // 0 1 4 9 0
			[]int{0, 1, 2, 1, 0},
		},
		{
			"One zero",
			6,
			[]int{0, 7, 9, 4, 8, 20},
			[]int{0, 1, 2, 3, 4, 5},
		},
		{
			"Only zero",
			1,
			[]int{0},
			[]int{0},
		},
		{
			"Several zeros",
			3,
			[]int{0, 0, 0},
			[]int{0, 0, 0},
		},
		{
			"One last zero",
			5,
			[]int{1, 4, 9, 7, 0},
			[]int{4, 3, 2, 1, 0},
		},
		{
			"Three zeros",
			10,
			[]int{0, 1, 4, 9, 0, 2, 3, 0, 8, 5}, // 0 1 4 9 0 2 3 0 8 5
			[]int{0, 1, 2, 1, 0, 1, 1, 0, 1, 2},
		},
		{
			"Adjacent zeros",
			10,
			[]int{0, 0, 0, 9, 1, 2, 3, 0, 8, 5},
			[]int{0, 0, 0, 1, 2, 2, 1, 0, 1, 2},
		},
		{
			"Two zeros 2",
			9,
			[]int{98, 0, 10, 77, 0, 59, 28, 0, 94},
			[]int{1, 0, 1, 1, 0, 1, 1, 0, 1},
		},
		{
			"Many zeros",
			20, // 10 13 31 35 39 0 0 59 0 66 68 73 74 0 0 0 87 8 96 99
			[]int{10, 13, 31, 35, 39, 0, 0, 59, 0, 66, 68, 73, 74, 0, 0, 0, 87, 8, 96, 99},
			[]int{5, 4, 3, 2, 1, 0, 0, 1, 0, 1, 2, 2, 1, 0, 0, 0, 1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		res := getNearestZeroDistances(tt.arr, tt.n)
		if !reflect.DeepEqual(res, tt.expected) {
			t.Errorf("Test: %s, expected result: %v, result: %v", tt.name, tt.expected, res)
		}
	}
}
