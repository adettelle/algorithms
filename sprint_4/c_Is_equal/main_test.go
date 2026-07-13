package main

import "testing"

func TestMain(t *testing.T) {
	tests := []struct {
		name     string
		line1    string
		line2    string
		expected bool
	}{
		// {
		// 	"Test_1",
		// 	"aag",
		// 	"xxd",
		// 	true,
		// },
		// {
		// 	"Test_2",
		// 	"aag",
		// 	"xdu",
		// 	false,
		// },
		// {
		// 	"Test_3",
		// 	"mxyskaoghi",
		// 	"qodfrgmslc",
		// 	true,
		// },
		{
			"Test_4",
			"aabc",
			"xxyx",
			false,
		},
		{
			"Test_5",
			"aaba",
			"xxyx",
			true,
		},
		{
			"Test_6",
			"abcabc",
			"xxxxxx",
			false,
		},
		{
			"Test_7",
			"abacaba",
			"abacabac",
			false,
		},
	}

	for _, tt := range tests {
		res := isEqual(tt.line1, tt.line2)
		if res != tt.expected {
			t.Errorf("%s, expected: %v, result: %v", tt.name, tt.expected, res)
		}
	}
}
