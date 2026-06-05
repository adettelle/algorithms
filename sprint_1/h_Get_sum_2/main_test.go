package main

import (
	"testing"
)

func TestFactorize(t *testing.T) {
	tests := []struct {
		name         string
		firstNumber  string
		secondNumber string
		expected     string
	}{
		{
			"Test 1",
			"1010",
			"1011",
			"10101",
		},
		{
			"Test 1",
			"1",
			"1",
			"10",
		},
		{
			"Test 3",
			"1010",
			"1",
			"1011",
		},
		{
			"Test 4",
			"10",
			"1011",
			"1101",
		},
	}

	for _, tt := range tests {
		res := getSum(tt.firstNumber, tt.secondNumber)

		if res != tt.expected {
			t.Errorf("%s: expected result: %s, result: %s", tt.name, tt.expected, res)
		}
	}
}
