package main

import "testing"

func TestLongestSubstr(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected int
	}{
		{
			"Similar letters",
			"kkkk",
			1,
		},
		{
			"Test2",
			"abcabcbb",
			3,
		},
		{
			"Test3",
			"pwwkew",
			3,
		},
		{
			"Test4",
			"ojodx",
			4,
		},
		{
			"Test5",
			"klmohjodx",
			6,
		},
		{
			"Test6",
			"klmohjodxzy",
			7,
		},
		{
			"Test7",
			"fprarfpoz",
			6,
		},
	}

	for _, tt := range tests {
		res := longestSubstr(tt.line)
		if res != tt.expected {
			t.Errorf("%s: expected: %d, result: %d", tt.name, tt.expected, res)
		}
	}
}
