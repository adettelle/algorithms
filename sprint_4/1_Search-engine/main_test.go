package main

import (
	"reflect"
	"testing"
)

/*
func TestMakeMap(t *testing.T) {
	tests := []struct {
		Name string
		Docs [][]string
	}{
		{
			"Test1",
			[][]string{
				{"buy", "flat", "in", "moscow"},
				{"want", "flat", "in", "moscow", "like", "crazy"},
			},
		},
	}

	// map[string][]WordInDocs
	for _, tt := range tests {
		res := makeMap(tt.Docs)

	}
}
*/

func TestCheck(t *testing.T) {
	tests := []struct {
		Name      string
		relevance []WordInDocs
		elem      WordInDocs
		limit     int
		expected  []WordInDocs
	}{
		// {
		// 	"Nothing to change",
		// 	[]WordInDocs{
		// 		{DocPosition: 9, WordPoints: 12},
		// 		{DocPosition: 2, WordPoints: 10},
		// 		{DocPosition: 5, WordPoints: 8},
		// 		{DocPosition: 3, WordPoints: 6},
		// 		{DocPosition: 7, WordPoints: 5},
		// 	},
		// 	WordInDocs{DocPosition: 10, WordPoints: 4},
		// 	5,
		// 	[]WordInDocs{
		// 		{DocPosition: 9, WordPoints: 12},
		// 		{DocPosition: 2, WordPoints: 10},
		// 		{DocPosition: 5, WordPoints: 8},
		// 		{DocPosition: 3, WordPoints: 6},
		// 		{DocPosition: 7, WordPoints: 5},
		// 	},
		// },
		{
			"Test1",
			[]WordInDocs{
				{DocPosition: 9, WordPoints: 12},
				{DocPosition: 2, WordPoints: 10},
				{DocPosition: 5, WordPoints: 8},
				{DocPosition: 3, WordPoints: 6},
				{DocPosition: 7, WordPoints: 5},
			},
			WordInDocs{DocPosition: 15, WordPoints: 7},
			5,
			[]WordInDocs{
				{DocPosition: 9, WordPoints: 12},
				{DocPosition: 2, WordPoints: 10},
				{DocPosition: 5, WordPoints: 8},
				{DocPosition: 15, WordPoints: 7},
				{DocPosition: 3, WordPoints: 6},
			},
		},
		{
			"Test change 1",
			[]WordInDocs{
				{DocPosition: 9, WordPoints: 12},
				{DocPosition: 2, WordPoints: 10},
				{DocPosition: 5, WordPoints: 8},
				{DocPosition: 3, WordPoints: 6},
				{DocPosition: 7, WordPoints: 5},
			},
			WordInDocs{DocPosition: 15, WordPoints: 7},
			5,
			[]WordInDocs{
				{DocPosition: 9, WordPoints: 12},
				{DocPosition: 2, WordPoints: 10},
				{DocPosition: 5, WordPoints: 8},
				{DocPosition: 15, WordPoints: 7},
				{DocPosition: 3, WordPoints: 6},
			},
		},
		{
			"Test change in short slice",
			[]WordInDocs{
				{DocPosition: 5, WordPoints: 8},
				{DocPosition: 3, WordPoints: 6},
				{DocPosition: 7, WordPoints: 5},
			},
			WordInDocs{DocPosition: 15, WordPoints: 7},
			5,
			[]WordInDocs{
				{DocPosition: 5, WordPoints: 8},
				{DocPosition: 15, WordPoints: 7},
				{DocPosition: 3, WordPoints: 6},
				{DocPosition: 7, WordPoints: 5},
			},
		},
	}

	for _, tt := range tests {
		collectTop(tt.relevance, tt.limit, tt.elem)
		if reflect.DeepEqual(tt.relevance, tt.expected) {
			t.Errorf("%s: expected: %v, result: %v", tt.Name, tt.expected, tt.relevance)
		}
	}
}

func TestGetTop(t *testing.T) {
	tests := []struct {
		Name      string
		relevance []WordInDocs
		elem      WordInDocs
		limit     int
		expected  []WordInDocs
	}{
		{
			"Nothing to change",
			[]WordInDocs{
				{DocPosition: 9, WordPoints: 12},
				{DocPosition: 2, WordPoints: 10},
				{DocPosition: 5, WordPoints: 8},
				{DocPosition: 3, WordPoints: 6},
				{DocPosition: 7, WordPoints: 5},
			},
			WordInDocs{DocPosition: 10, WordPoints: 4},
			5,
			[]WordInDocs{
				{DocPosition: 9, WordPoints: 12},
				{DocPosition: 2, WordPoints: 10},
				{DocPosition: 5, WordPoints: 8},
				{DocPosition: 3, WordPoints: 6},
				{DocPosition: 7, WordPoints: 5},
			},
		},
		{
			"Test1",
			[]WordInDocs{
				{DocPosition: 9, WordPoints: 12},
				{DocPosition: 2, WordPoints: 10},
				{DocPosition: 5, WordPoints: 8},
				{DocPosition: 3, WordPoints: 6},
				{DocPosition: 7, WordPoints: 5},
			},
			WordInDocs{DocPosition: 15, WordPoints: 7},
			5,
			[]WordInDocs{
				{DocPosition: 9, WordPoints: 12},
				{DocPosition: 2, WordPoints: 10},
				{DocPosition: 5, WordPoints: 8},
				{DocPosition: 15, WordPoints: 7},
				{DocPosition: 3, WordPoints: 6},
			},
		},
		{
			"Test change 1",
			[]WordInDocs{
				{DocPosition: 9, WordPoints: 12},
				{DocPosition: 2, WordPoints: 10},
				{DocPosition: 5, WordPoints: 8},
				{DocPosition: 3, WordPoints: 6},
				{DocPosition: 7, WordPoints: 5},
			},
			WordInDocs{DocPosition: 15, WordPoints: 7},
			5,
			[]WordInDocs{
				{DocPosition: 9, WordPoints: 12},
				{DocPosition: 2, WordPoints: 10},
				{DocPosition: 5, WordPoints: 8},
				{DocPosition: 15, WordPoints: 7},
				{DocPosition: 3, WordPoints: 6},
			},
		},
		{
			"Test change in short slice",
			[]WordInDocs{
				{DocPosition: 5, WordPoints: 8},
				{DocPosition: 3, WordPoints: 6},
				{DocPosition: 7, WordPoints: 5},
			},
			WordInDocs{DocPosition: 15, WordPoints: 7},
			5,
			[]WordInDocs{
				{DocPosition: 5, WordPoints: 8},
				{DocPosition: 15, WordPoints: 7},
				{DocPosition: 3, WordPoints: 6},
				{DocPosition: 7, WordPoints: 5},
			},
		},
	}

	// map[string][]WordInDocs
	for _, tt := range tests {
		got := collectTop(tt.relevance, tt.limit, tt.elem)
		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("%s: expected: %v, result: %v", tt.Name, tt.expected, tt.relevance)
		}
	}
}
