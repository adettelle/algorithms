package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := makeScanner()
	scanner.Scan()
	line1 := scanner.Text()
	scanner.Scan()
	line2 := scanner.Text()

	res := isEqual(line1, line2)
	if !res {
		fmt.Println("NO")
	} else {
		fmt.Println("YES")
	}
}

func makeScanner() *bufio.Scanner {
	const maxCapacity = 1024 * 1024
	buf := make([]byte, maxCapacity)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(buf, maxCapacity)
	return scanner
}

func isEqual(line1, line2 string) bool {
	if len(line1) != len(line2) {
		return false
	}
	chars := make(map[byte]byte)

	for i := 0; i < len(line1); i++ {
		_, ok := chars[line1[i]]
		if !ok {
			if _, ok := findKeyByValue(chars, line2[i]); !ok {
				chars[line1[i]] = line2[i]
			} else {
				return false
			}

		} else {
			if chars[line1[i]] != line2[i] {
				return false
			}
		}
	}
	return true
}

func findKeyByValue(data map[byte]byte, targetValue byte) (byte, bool) {
	for key, val := range data {
		if val == targetValue {
			return key, true
		}
	}
	return 0, false
}
