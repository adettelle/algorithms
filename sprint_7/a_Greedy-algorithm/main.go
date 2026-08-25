package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &arr[i])
	}

	res := findMaxIncome(arr)
	fmt.Println(res)
}

func findMaxIncome(arr []int) int {
	if len(arr) < 2 {
		return 0
	}

	res := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] > arr[i-1] {
			res += arr[i] - arr[i-1]
		}
	}

	return res
}
