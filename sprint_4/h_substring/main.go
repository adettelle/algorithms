package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	fmt.Println(longestSubstr(line))
}

func longestSubstr(line string) int {
	count := 0
	maxLen := 0
	letters := make(map[byte]int)

	for i := 0; i < len(line); i++ {
		if oldI, ok := letters[line[i]]; !ok {
			letters[line[i]] = i
			count++
			if count > maxLen {
				maxLen = count
			}
		} else {
			count = min(count, i-oldI-1)
			count++
			letters[line[i]] = i
			if count > maxLen {
				maxLen = count
			}
		}
	}
	return maxLen
}
