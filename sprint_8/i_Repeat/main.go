package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var line string
	fmt.Fscan(reader, &line)

	fmt.Println(FindRepeat(line))
}

func FindRepeat(line string) int {
	// sample := string(line[0])
	var sample string

	for i := 0; i < len(line); i++ {
		sample = sample + string(line[i])
		// fmt.Println("sample:", sample)

		ok := checkSample(line, string(sample))
		if ok {
			// fmt.Println(len(line))
			return len(line) / len(sample)
		}
	}
	return len(line)
}

func checkSample(line string, sample string) bool {
	// fmt.Println("line:", line, "; sample:", sample)
	new := "1"

	newLine := strings.Replace(line, sample, new, -1)
	// fmt.Println(newLine)
	// _, err := strconv.ParseUint(newLine, 10, 64)
	// if err != nil {
	// 	fmt.Println(" --- ")
	// 	return false
	// }

	ok := isNumber(newLine)
	if !ok {
		return false
	}
	return true
}

func isNumber(s string) bool {
	if s == "" {
		return false
	}

	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}

	return true
}
