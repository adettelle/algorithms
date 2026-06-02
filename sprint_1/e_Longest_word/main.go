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
	readInt(scanner)
	line := readLine(scanner)

	word := getLongestWord(line)
	fmt.Println(word)
	fmt.Println(len(word))
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	stringInt := scanner.Text()
	res, _ := strconv.Atoi(stringInt)
	return res
}

func readLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}

func getLongestWord(line string) string {
	maxLen := 0
	foundWord := ""

	words := strings.Split(line, " ")
	for _, word := range words {
		elems := len([]rune(word))
		if elems > maxLen {
			maxLen = elems
			foundWord = word
		}
	}
	return foundWord
}
