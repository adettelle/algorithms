package main

import (
	"reflect"
	"testing"
)

func TestNeightbours(t *testing.T) {
	tests := []struct {
		name     string
		n        int // rows
		m        int // column
		matrix   [][]int
		row      int
		col      int
		expected []int
	}{
		{
			"two elements_1",
			4,
			3,
			[][]int{
				{1, 2, 3},
				{0, 2, 6},
				{7, 4, 1},
				{2, 7, 0},
			},
			3,
			0,
			[]int{7, 7},
		},
		{
			"two elements_2",
			4,
			3,
			[][]int{
				{1, 2, 3},
				{0, 2, 6},
				{7, 4, 1},
				{2, 7, 0},
			},
			0,
			0,
			[]int{0, 2},
		},
		{
			"four elements",
			4,
			3,
			[][]int{
				{1, 2, 3},
				{0, 2, 6},
				{7, 4, 1},
				{2, 7, 0},
			},
			1,
			1,
			[]int{0, 2, 4, 6},
		},
		{
			"three elements",
			4,
			3,
			[][]int{
				{1, 2, 3},
				{0, 2, 6},
				{7, 4, 1},
				{2, 7, 0},
			},
			3,
			1,
			[]int{0, 2, 4},
		},
		{
			"one element",
			2,
			1,
			[][]int{
				{1},
				{9},
			},
			0,
			0,
			[]int{9},
		},
		{
			"no elements",
			1,
			1,
			[][]int{
				{1},
			},
			0,
			0,
			[]int{},
		},
		{
			"one element_2",
			1,
			2,
			[][]int{
				{-3, 8},
			},
			0,
			0,
			[]int{8},
		},
	}

	for _, tt := range tests {
		res := getNeighbours(tt.matrix, tt.row, tt.col)
		if !reflect.DeepEqual(res, tt.expected) {
			t.Errorf("Test: %s; expected result: %v, result: %v", tt.name, tt.expected, res)
		}
	}
}
