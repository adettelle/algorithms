package main

import (
	"fmt"
	"reflect"
)

func merge_sort(arr []int, lf int, rg int) {
	if len(arr[lf:rg]) <= 1 {
		return
	}

	mid := lf + (rg-lf)/2

	merge_sort(arr, lf, mid)
	merge_sort(arr, mid, rg)

	merged := merge(arr, lf, mid, rg)
	copy(arr[lf:rg], merged)

}

// сливает две отсортированные части одного и того же массива
// в один отсортированный массив
func merge(arr []int, left int, mid int, right int) (result []int) {
	leftArr := arr[left:mid]
	rightArr := arr[mid:right]

	i, j := 0, 0
	for i < len(leftArr) && j < len(rightArr) {
		if leftArr[i] < rightArr[j] {
			result = append(result, leftArr[i])
			i++
		} else {
			result = append(result, rightArr[j])
			j++
		}
	}
	result = append(result, leftArr[i:]...)
	result = append(result, rightArr[j:]...)

	return result
}

func test() {
	a := []int{1, 4, 9, 2, 10, 11}
	b := merge(a, 0, 3, 6) // [1, 4, 9] + [2, 10, 11] --> 1 2 4 9 10 11
	expected := []int{1, 2, 4, 9, 10, 11}
	if !reflect.DeepEqual(b, expected) {
		panic("WA. Merge")
	}

	c := []int{1, 4, 2, 10, 1, 2}
	merge_sort(c, 0, 6)
	expected = []int{1, 1, 2, 2, 4, 10}
	if !reflect.DeepEqual(c, expected) {
		panic("WA. MergeSort")
	}
}

func main() {
	a := []int{1, 4, 9, 2, 10, 11}
	b := merge(a, 0, 3, 6)
	fmt.Println(b)
	// expected := []int{1, 2, 4, 9, 10, 11}
	// test()

	c := []int{1, 4, 2, 10, 1, 2}
	merge_sort(c, 0, 6)
	fmt.Println(c) // []int{1, 1, 2, 2, 4, 10}
	// test()
}
