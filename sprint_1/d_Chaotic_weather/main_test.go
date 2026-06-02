package main

import "testing"

func TestWeather(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		arr      []int
		expected int
	}{
		{
			"Three days",
			7,
			[]int{-1, -10, -8, 0, 2, 0, 5}, // -1, 2, 5
			3,
		},
		{
			"One day",
			7,
			[]int{5},
			1,
		},
		{
			"Warm weather",
			5,
			[]int{1, 2, 5, 4, 8},
			2,
		},
		{
			"Cold weather",
			5,
			[]int{-1, -2, -5, -4, -8},
			2,
		},
		{
			"No values",
			0,
			[]int{},
			0,
		},
	}

	for _, tt := range tests {
		res := getWeatherRandomness(tt.arr)
		if res != tt.expected {
			t.Errorf("Test %s: expected result: %d, result: %d", tt.name, tt.expected, res)
		}
	}
}
