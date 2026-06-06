// https://contest.yandex.ru/contest/22450/run-report/162784284/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	var k int
	fmt.Sscan(scanner.Text(), &k)

	matrix := make([][]int, 4)

	for i := 0; i < 4; i++ {
		scanner.Scan()
		line := scanner.Text()

		row := make([]int, 4)
		for j := 0; j < 4; j++ {
			if line[j] == '.' {
				row[j] = 0
			} else {
				row[j] = int(line[j] - '0')
			}
		}
		matrix[i] = row
	}
	fmt.Println("k =", k)
	fmt.Println(matrix)
	x := Search(k, matrix)
	fmt.Println(x)
}

func Search(k int, matrix [][]int) int {
	n := k * 2
	count := 0
	nums := map[int]int{}

	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[row]); col++ {
			elem := matrix[row][col]
			if elem != 0 {
				_, ok := nums[elem]
				if !ok {
					nums[elem] = 1
				} else {
					nums[elem]++
				}
			}
		}
	}

	for _, val := range nums {
		if val <= n {
			count++
		}
	}
	return count
}
