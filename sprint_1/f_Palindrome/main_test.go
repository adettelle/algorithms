package main

import "testing"

func TestPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		expected bool
		res      string
	}{
		{
			"Positive test",
			"A man, a plan, a canal: Panama",
			true,
			"True",
		},
		{
			"Negative test",
			"zo",
			false,
			"False",
		},
		{
			"Null test",
			"",
			false,
			"False",
		},
		{
			"Null test 2",
			" ",
			false,
			"False",
		},
		{
			"Negative test 2",
			`-Luke, I'm your Father! -N00Oo! -oo00n -rehTAFruoymiekul
`,
			true,
			"True",
		},
	}

	for _, tt := range tests {
		res := isPalindrome(tt.line)
		if res != tt.expected {
			t.Errorf("Test %s: expected result: %v, result: %v", tt.name, tt.expected, res)
		}
	}
}
