package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var rows, cols int
	fmt.Fscan(reader, &rows, &cols)

	field := make([][]int, rows)

	for r := 0; r < rows; r++ {
		var str string
		fmt.Fscan(reader, &str)

		arr := ([]byte)(str)

		for c := 0; c < cols; c++ {
			field[r] = append(field[r], int(arr[c])-48)
		}
	}

	// fmt.Println(field)
	fmt.Println(findMaxFLowers(field, rows, cols))
}

func findMaxFLowers(field [][]int, rows, cols int) int {

	for r := rows - 1; r >= 0; r-- {
		for c := 0; c < cols; c++ {
			leftVal := 0
			bottomVal := 0

			if r < rows-1 {
				bottomVal = field[r+1][c]
			}
			if c > 0 {
				leftVal = field[r][c-1]
			}
			field[r][c] = max(bottomVal, leftVal) + field[r][c]
		}
	}
	// fmt.Println(field)
	return field[0][cols-1]
}
