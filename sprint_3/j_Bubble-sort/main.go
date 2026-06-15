package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	nums := make([]int, n)
	for i := range n {
		fmt.Fscan(reader, &nums[i])
	}
	bubbleSort(nums, writer)
}

func isSorted(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}
	return true
}

func bubbleSort(arr []int, writer *bufio.Writer) {
	if isSorted(arr) {
		for i := 0; i < len(arr); i++ {
			fmt.Print(arr[i], " ")
		}
		return
	}

	for j := 0; j < len(arr)-1; j++ {
		changed := false
		for i := 1; i < len(arr)-j; i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				changed = true
			}
		}

		if !changed {
			break
		}

		for i := 0; i < len(arr); i++ {
			fmt.Fprint(writer, arr[i], " ")
		}
		fmt.Fprintln(writer)
	}
}
