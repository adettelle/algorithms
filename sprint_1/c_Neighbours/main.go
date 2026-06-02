package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)

	var m int
	fmt.Fscan(reader, &m)

	matrix := make([][]int, n)

	for i := range matrix {
		matrix[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Fscan(reader, &matrix[i][j])
		}
	}

	var row, col int
	fmt.Fscan(reader, &row)
	fmt.Fscan(reader, &col)

	res := getNeighbours(matrix, row, col)
	for _, num := range res {
		fmt.Print(num, " ")
	}
}

func getNeighbours(matrix [][]int, row int, col int) []int {
	m := len(matrix[0])
	n := len(matrix)

	neighbours := []int{}
	if row-1 >= 0 {
		neighbours = append(neighbours, matrix[row-1][col])
	}
	if row+1 < n {
		neighbours = append(neighbours, matrix[row+1][col])
	}
	if col-1 >= 0 {
		neighbours = append(neighbours, matrix[row][col-1])
	}
	if col+1 < m {
		neighbours = append(neighbours, matrix[row][col+1])
	}
	if len(neighbours) <= 1 {
		return neighbours
	}
	slices.Sort(neighbours)
	return neighbours
}
