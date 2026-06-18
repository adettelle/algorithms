package main

import (
	"reflect"
	"testing"
)

func TestBiggerComparator(t *testing.T) {
	tests := []struct {
		name     string
		s1       Student
		s2       Student
		expected bool
	}{
		{
			"Test negative",
			Student{name: "alla", points: 4, fine: 100},
			Student{name: "gena", points: 6, fine: 1000},
			false,
		},
		{
			"Test positive",
			Student{name: "rita", points: 9, fine: 100},
			Student{name: "gena", points: 6, fine: 1000},
			true,
		},
		{
			"Equal points",
			Student{name: "alla", points: 4, fine: 100},
			Student{name: "gena", points: 4, fine: 1000},
			true,
		},
		{
			"Equal points 2",
			Student{name: "rita", points: 4, fine: 100},
			Student{name: "gena", points: 4, fine: 1000},
			true,
		},
		{
			"Equal points and fines",
			Student{name: "rita", points: 4, fine: 500},
			Student{name: "gena", points: 4, fine: 500},
			false,
		},
	}

	for _, tt := range tests {
		res := biggerComparator(tt.s1, tt.s2)
		if res != tt.expected {
			t.Errorf("%s: expected: %v, result: %v", tt.name, tt.expected, res)
		}
	}
}

// quickSortWithComparator
func TestQuickSortWithComparator(t *testing.T) {
	tests := []struct {
		name       string
		students   []Student
		comparator func(Student, Student) bool
		expected   []Student
	}{
		{
			"Test 1",
			[]Student{
				{name: "alla", points: 4, fine: 100},
				{name: "gena", points: 6, fine: 1000},
				{name: "gosha", points: 2, fine: 90},
				{name: "rita", points: 2, fine: 90},
				{name: "timofey", points: 4, fine: 80},
			},
			biggerComparator,
			[]Student{
				{name: "gena", points: 6, fine: 1000},
				{name: "timofey", points: 4, fine: 80},
				{name: "alla", points: 4, fine: 100},
				{name: "gosha", points: 2, fine: 90},
				{name: "rita", points: 2, fine: 90},
			},
		},
		{
			"Test names",
			[]Student{
				{name: "alla", points: 0, fine: 0},
				{name: "rita", points: 0, fine: 0},
				{name: "gosha", points: 0, fine: 0},
				{name: "gena", points: 0, fine: 0},
				{name: "timofey", points: 0, fine: 0},
			},
			biggerComparator,
			[]Student{
				{name: "alla", points: 0, fine: 0},
				{name: "gena", points: 0, fine: 0},
				{name: "gosha", points: 0, fine: 0},
				{name: "rita", points: 0, fine: 0},
				{name: "timofey", points: 0, fine: 0},
			},
		},
	}

	for _, tt := range tests {
		res := quickSortWithComparator(tt.students, tt.comparator)
		if !reflect.DeepEqual(res, tt.expected) {
			t.Errorf("%s: expected: %v, result: %v", tt.name, tt.expected, res)
		}
	}
}
