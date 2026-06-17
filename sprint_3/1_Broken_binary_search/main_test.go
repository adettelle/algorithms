package main

import (
	"testing"
)

func TestBinarySearchMax(t *testing.T) {
	tests := []struct {
		name        string
		arr         []int
		expectedIdx int
	}{
		{
			"Test 1",
			[]int{19, 21, 100, 101, 1, 4, 5, 7, 12},
			3,
		},
		{
			"Test 2",
			[]int{15, 17, 19, 21, 100, 101, 110, 1, 4, 5, 7, 12, 13},
			6,
		},
		{
			"Test 3", // TODO
			[]int{8, 10, 0, 2, 4},
			1,
		},
		{
			"Test 4",
			[]int{100, 0, 1, 4, 5, 7, 12, 15, 19},
			0,
		},
		{
			"Test 5",
			[]int{87, 89, 90, 95, 100, 101, 105, 110, 120},
			8,
		},
	}

	for _, tt := range tests {
		res := binarySearchMax(tt.arr)
		if res != tt.expectedIdx {
			t.Errorf("%s: expected: %d; result: %d", tt.name, tt.expectedIdx, res)
		}
	}
}

func TestBrokenSearch(t *testing.T) {
	tests := []struct {
		name        string
		arr         []int
		target      int
		expectedIdx int
	}{
		{
			"Test_second number",
			[]int{19, 21, 100, 101, 1, 4, 5, 7, 12},
			21,
			1,
		},
		{
			"Test_first number",
			[]int{15, 17, 19, 21, 100, 101, 110, 1, 4, 5, 7, 12, 13},
			15,
			0,
		},
		{
			"Test_last number", // TODO
			[]int{8, 10, 0, 2, 4},
			4,
			4,
		},
		{
			"Test_before last number",
			[]int{100, 0, 1, 4, 5, 7, 12, 15, 19},
			15,
			7,
		},
		{
			"Test_in the middle number",
			[]int{87, 89, 90, 95, 100, 101, 105, 110, 120},
			100,
			4,
		},
		{
			"Test_second num",
			[]int{19, 21, 100, 101, 1, 4, 5, 7, 12},
			21,
			1,
		},
	}

	for _, tt := range tests {
		res := brokenSearch(tt.arr, tt.target)
		if res != tt.expectedIdx {
			t.Errorf("%s: expected: %d; result: %d", tt.name, tt.expectedIdx, res)
		}
	}
}
