package main

import (
	"reflect"
	"testing"
)

func TestFactorize(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		expected []int
	}{
		{
			"Test 1",
			8,
			[]int{2, 2, 2},
		},
		{
			"Test 2",
			13,
			[]int{13},
		},
		{
			"Test 3",
			100,
			[]int{2, 2, 5, 5},
		},
		{
			"Test 4",
			2,
			[]int{2},
		},
		{
			"Test 5",
			25,
			[]int{5, 5},
		},
		{
			"Test 6",
			21,
			[]int{3, 7},
		},
		{
			"Test 7",
			105,
			[]int{3, 5, 7},
		},
	}

	for _, tt := range tests {
		var x string
		res := factorize(tt.n)

		if !reflect.DeepEqual(res, tt.expected) {
			t.Errorf("%s: expected result: %v, result: %v", tt.name, tt.expected, x)
		}
	}
}
