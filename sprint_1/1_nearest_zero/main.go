// https://contest.yandex.ru/contest/22450/run-report/162781650/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var b strings.Builder

	scanner := makeScanner()
	n := readInt(scanner)
	arr := readArray(scanner)

	zeros := Zeros(arr)

	res := getNearestZeroDistances(zeros, n)

	for _, num := range res {
		b.WriteString(strconv.Itoa(num))
		b.WriteByte(' ')
	}

	fmt.Print(b.String())
}

func getNearestZeroDistances(zeros []int, n int) []int {
	res := make([]int, n)

	// before zeros[0]
	for i := zeros[0]; i >= 0; i-- {
		res[i] = zeros[0] - i
	}

	// after zeros[0] and before zeros[last]
	for k := 0; k < len(zeros)-1; k++ {
		sl := res[zeros[k] : zeros[k+1]+1]

		middle := len(sl) / 2
		if len(sl)%2 != 0 {
			middle++
		}
		for j := 0; j < middle; j++ {
			sl[j] = j
			sl[len(sl)-j-1] = j
		}
	}

	// after zeros[last]
	num := 0
	for i := zeros[len(zeros)-1]; i < n; i++ {
		res[i] = num
		num++
	}
	return res
}

func Zeros(arr []int) []int {
	zeros := []int{}

	if len(arr) == 1 {
		arr[0] = 0
		return arr
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			zeros = append(zeros, i)
		}
	}
	return zeros
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
