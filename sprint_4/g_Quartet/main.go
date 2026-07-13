package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	n := readInt(scanner)
	s := readInt(scanner)

	arr := readArray(scanner)
	res := findQuartet(arr, n, s)
	fmt.Println(len(res))
	for r := 0; r < len(res); r++ {
		for c := 0; c < len(res[r]); c++ {
			fmt.Print(res[r][c], " ")
		}
		fmt.Println()
	}
}

func findQuartet(arr []int, n int, target int) [][]int {
	slices.Sort(arr)
	fmt.Println(arr)

	res := [][]int{}

	for i := 0; i < n-3; i++ {
		if i > 0 && arr[i] == arr[i-1] {
			continue
		}
		first := arr[i]

		for j := i + 1; j < n-2; j++ {
			if j > i+1 && arr[j] == arr[j-1] {
				continue
			}
			second := arr[j]

			need := target - first - second

			left := j + 1
			right := n - 1

			for left < right {
				sum := arr[left] + arr[right]

				if sum < need {
					left++
				} else if sum > need {
					right--
				} else {
					third := arr[left]
					forth := arr[right]

					row := []int{first, second, third, forth}
					res = append(res, row)
					left++
					right--
				}
			}
		}
	}

	return res
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	return n
}

func readArray(scanner *bufio.Scanner) []int {
	scanner.Scan()
	listStr := strings.Split(scanner.Text(), " ")
	arr := make([]int, len(listStr))

	for i := 0; i < len(listStr); i++ {
		arr[i], _ = strconv.Atoi(listStr[i])
	}
	return arr
}
