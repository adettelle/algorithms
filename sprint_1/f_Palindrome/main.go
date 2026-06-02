package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	scanner := makeScanner()
	scanner.Scan()
	line := scanner.Text()

	if isPalindrome(line) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 1024 * 1024
	buf := make([]byte, maxCapacity)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func isPalindrome(line string) bool {
	res := strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			return r
		}
		return -1
	}, line)
	if res == "" {
		return false
	}
	res = strings.ToLower(res)
	palindrome := true

	lineInRunes := []rune(res)
	n := len(lineInRunes)

	for i := 0; i < n/2; i++ {
		if lineInRunes[i] != lineInRunes[n-1-i] {
			palindrome = false
		}
	}
	return palindrome
}
