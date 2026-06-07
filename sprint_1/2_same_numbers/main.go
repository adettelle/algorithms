// https://contest.yandex.ru/contest/22450/run-report/162784284/
package main

import (
	"bufio"
	"fmt"
	"os"
)

const players = 2

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	var k int
	fmt.Sscan(scanner.Text(), &k)

	const size = 4
	matrix := make([][]int, size)

	for i := 0; i < size; i++ {
		scanner.Scan()
		line := scanner.Text()

		row := make([]int, size)
		for j := 0; j < size; j++ {
			if line[j] == '.' {
				row[j] = 0
			} else {
				row[j] = int(line[j] - '0')
			}
		}
		matrix[i] = row
	}

	x := Search(k, matrix, players)
	fmt.Println(x)
}

func Search(k int, matrix [][]int, players int) int {

	n := k * players
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
