package main

import "testing"

func TestLongestWord(t *testing.T) {
	test := []struct {
		name         string
		n            int // длина текста
		text         string
		expectedWord string
		expectedLen  int
	}{
		{
			"Positive test 1",
			19,
			"i love segment tree",
			"segment",
			7,
		},
		{
			"Positive test 2",
			21,
			"frog jumps from river",
			"jumps",
			5,
		},
		{
			"Test with leading and trailing spaces ",
			21,
			" frog jumps from river  ",
			"jumps",
			5,
		},
		{
			"Zero test",
			0,
			"",
			"",
			0,
		},
		{
			"One word",
			4,
			"word",
			"word",
			4,
		},
	}
	for _, tt := range test {
		word := getLongestWord(tt.text)
		wordLen := len([]rune(word))
		if word != tt.expectedWord && wordLen != tt.expectedLen {
			t.Errorf("Test %s: expected word: %s, result word: %s, expected len: %d, result len: %d",
				tt.name, tt.expectedWord, word, tt.expectedLen, wordLen)
		}
	}
}
