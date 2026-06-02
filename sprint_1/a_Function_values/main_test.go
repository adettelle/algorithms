package main

import "testing"

func TestQuadraticEquation(t *testing.T) {
	tests := []struct {
		name      string
		numOfTest string
		arr       []int
		expected  int
	}{
		{
			"Positive test",
			"Test N1",
			[]int{-8, -5, -2, 7}, // a, x, b, c
			-183,
		},
		{
			"Positive test",
			"Test N2",
			[]int{8, 2, 9, -10}, // a, x, b, c
			40,
		},
		{
			"Zero test",
			"Test N1",
			[]int{0, 0, 0, 0},
			0,
		},
	}

	for _, tt := range tests {
		res := evaluateFunction(tt.arr[0], tt.arr[1], tt.arr[2], tt.arr[3])
		if res != tt.expected {
			t.Errorf("N%s: expected result: %d, result: %d", tt.numOfTest, tt.expected, res)
		}
	}
}
