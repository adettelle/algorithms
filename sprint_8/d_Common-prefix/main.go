package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)

	data := make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &data[i])
	}

	fmt.Println(CommonPrefix(data))
}

func CommonPrefix(data []string) int {
	word := data[0]
	prefix := 0

	for i := 0; i < len(word); i++ {
		for j := 1; j < len(data); j++ {
			if data[j][i] != word[i] {
				return prefix
			}
		}
		prefix++
	}
	return prefix
}
