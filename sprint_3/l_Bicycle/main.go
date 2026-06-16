package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := makeScanner()
	_ = readInt(scanner)
	amounts := readArray(scanner)
	price := readInt(scanner)

	firstDay := findDayBinarySeach(amounts, 0, price)
	secondDay := findDayBinarySeach(amounts, firstDay, price*2)

	if firstDay == -1 {
		fmt.Println(-1, -1)
	} else {
		fmt.Println(firstDay, secondDay)
	}
}

func findDayBinarySeach(arr []int, start int, price int) int {
	left := start
	right := len(arr) - 1
	res := -1

	for left <= right {
		middle := left + (right-left)/2

		if arr[middle] >= price {
			res = middle
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	if res == -1 {
		return -1
	}
	return res + 1 // дни в условии нумеруются с 1, а не с 0
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 3 * 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
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

/*
// naive O(n)
func findDayNaive(amountsPerDay []int, price int) int { // days int,
	for i := range len(amountsPerDay) {
		fmt.Println("day =", i, "amountsPerDay[i] =", amountsPerDay[i])
		if amountsPerDay[i] >= price {
			return i + 1
		}
	}
	return -1
}
*/
