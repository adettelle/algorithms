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

	var rows int
	var cols int

	fmt.Fscan(reader, &rows)
	fmt.Fscan(reader, &cols)

	matrix := make([][]int, rows)

	for r := range matrix {
		matrix[r] = make([]int, cols)
		for c := range cols {
			fmt.Fscan(reader, &matrix[r][c])
		}
	}

	for col := range cols {
		for row := range rows {
			fmt.Fprint(writer, matrix[row][col])
			if row+1 < rows {
				fmt.Fprint(writer, " ")
			}
		}
		fmt.Fprintln(writer)
	}
}

/*
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	rows := readInt(scanner)
	cols := readInt(scanner)

	matrix := make([][]int, rows)
	for row := range matrix {
		matrix[row] = make([]int, cols)
		matrix[row] = readArray(scanner)
	}

	for col := range cols {
		for row := range rows {
			fmt.Print(matrix[row][col], " ")
		}
		fmt.Println()
	}
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	strInt := scanner.Text()
	x, _ := strconv.Atoi(strInt)
	return x
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
*/
