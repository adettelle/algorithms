package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var rows int
	fmt.Fscan(reader, &rows)
	matrix := make([][]int, rows)

	for row := range matrix {
		matrix[row] = make([]int, 2)
		for col := range 2 {
			fmt.Fscan(reader, &matrix[row][col])
		}
	}

	result := mergeSegments(matrix)

	for r := 0; r < len(result); r++ {
		for c := 0; c < len(result[r]); c++ {
			fmt.Fprint(writer, result[r][c], " ")
		}
		fmt.Fprintln(writer)
	}
}

func less(a, b []int) bool {
	if a[0] != b[0] {
		return a[0] < b[0]
	}
	return a[1] < b[1]
}

func mergeSortByComparator(arr [][]int) [][]int { // , less func(int, int) bool
	if len(arr) <= 1 { // базовый случай рекурсии
		return arr
	}

	mid := len(arr) / 2
	left := mergeSortByComparator(arr[:mid])
	right := mergeSortByComparator(arr[mid:])

	return mergeByComparator(left, right)
}

// сливаем результаты
func mergeByComparator(left, right [][]int) [][]int {
	res := [][]int{} //

	i, j := 0, 0
	for i < len(left) && j < len(right) {
		// выбираем, из какого массива забрать минимальный элемент
		if less(left[i], right[j]) {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}
	//  Добавляем остатки, если в одном из списков еще есть элементы
	res = append(res, left[i:]...)
	res = append(res, right[j:]...)
	return res
}

/*
func insertionSortByComparator(arr [][]int) { // , less func(int, int) bool

		for i := 1; i < len(arr); i++ {
			itemToInsert := arr[i]
			j := i
			for j > 0 && less(itemToInsert, arr[j-1]) {
				arr[j] = arr[j-1]
				j--
			}
			arr[j] = itemToInsert
		}
	}
*/

func mergeSegments(arr [][]int) [][]int {
	arr = mergeSortByComparator(arr)
	res := [][]int{}

	current := arr[0]

	for i := 1; i < len(arr); i++ {
		if current[1] >= arr[i][0] {
			if current[1] < arr[i][1] {
				current[1] = arr[i][1]
			}
			continue
		} else {
			res = append(res, current)
		}
		current = arr[i]
	}
	res = append(res, current)
	return res
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())
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
