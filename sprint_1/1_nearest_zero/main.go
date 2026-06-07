// https://contest.yandex.ru/contest/22450/run-report/162860542/
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var b strings.Builder

	scanner := makeScanner()
	n := readInt(scanner)
	arr := readArray(scanner)

  res := getNearestZeroDistances(arr, n)

	for _, num := range res {
		b.WriteString(strconv.Itoa(num))
		b.WriteByte(' ')
	}

	fmt.Print(b.String())
}

func getNearestZeroDistances(arr []int, n int) []int {
	maxInt := math.MaxInt32
	counter := 1
	afterZero := false

	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			afterZero = true
			counter = 1
			continue
		}
		if !afterZero {
			arr[i] = maxInt
		} else {
			arr[i] = counter
			counter++
		}
	}

	afterZero = false
	for i := n - 1; i >= 0; i-- {
		if arr[i] == 0 {
			afterZero = true
			counter = 1
			continue
		}
		if afterZero {
			if counter < arr[i] {
				arr[i] = counter
			}
			counter++
		}
	}
	return arr
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 1024 * 1024 * 11
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	strInt := scanner.Text()
	res, _ := strconv.Atoi(strInt)
	return res
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
