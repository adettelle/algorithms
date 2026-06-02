package main

import (
	"testing"
)

func TestParity(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected string
	}{
		{
			"Positive test 1",
			[]int{-1, -3, 3},
			"WIN",
		},
		{
			"Positive test 2",
			[]int{1, 11, 33},
			"WIN",
		},
		{
			"Positive test 3",
			[]int{0, 0, 2},
			"WIN",
		},
		{
			"Positive test 4",
			[]int{7, 11, 7},
			"WIN",
		},
		{
			"Positive test 5",
			[]int{6, -2, 0},
			"WIN",
		},
		{
			"Negative test",
			[]int{1, 2, -3},
			"FAIL",
		},
	}

	for _, tt := range tests {
		var x string
		res := checkParity(tt.nums[0], tt.nums[1], tt.nums[2])
		if res {
			x = "WIN"
		} else {
			x = "FAIL"
		}
		if x != tt.expected {
			t.Errorf("%s: expected result: %s, result: %s", tt.name, tt.expected, x)
		}
	}
}
