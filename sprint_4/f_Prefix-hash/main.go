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

	var base int64
	fmt.Fscan(reader, &base)

	var module int64
	fmt.Fscan(reader, &module)

	var line string
	fmt.Fscan(reader, &line)

	var t int
	fmt.Fscan(reader, &t)

	arr := make([][2]int, t)

	for i := range arr {
		// var l, r int
		fmt.Fscan(reader, &arr[i][0], &arr[i][1])
		// arr[i] = []int{l, r}
	}
	fmt.Println(arr)
	/*
		scanner := makeScanner()
		scanner.Scan()
		base, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		module, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		line := scanner.Text() //
		scanner.Scan()
		t, _ := strconv.Atoi(scanner.Text())

		arr := make([][]int, t)
		for i := 0; i < t; i++ {
			row := readArray(scanner)
			arr[i] = row
		}
	*/
	// fmt.Println(base, module, line, t)
	// fmt.Println(arr)
	// fmt.Println()

	for _, row := range arr {
		// fmt.Println(row)
		subLine := line[row[0]-1 : row[1]]
		// fmt.Println(getHash(subLine, int64(base), int64(module)))
		fmt.Fprintln(writer, getHash(subLine, int64(base), int64(module)))
	}
	// fmt.Fprintln(writer)
}

func getHash(line string, base int64, module int64) int64 {
	var h int64
	for i := 0; i < len(line); i++ {
		h = (h*base + int64(line[i])) % module
	}
	return h
}

func makeScanner() *bufio.Scanner {
	scanner := bufio.NewScanner(os.Stdin)
	maxCap := 1024 * 1024 * 8
	buf := make([]byte, maxCap)
	scanner.Buffer(buf, maxCap)
	return scanner
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	x, _ := strconv.Atoi(scanner.Text())
	return x
}

func readArray(scanner *bufio.Scanner) []int {
	scanner.Scan()
	strArr := strings.Split(scanner.Text(), " ")
	res := make([]int, len(strArr))

	for i := 0; i < len(strArr); i++ {
		x, _ := strconv.Atoi(strArr[i])
		res[i] = x
	}
	return res
}
