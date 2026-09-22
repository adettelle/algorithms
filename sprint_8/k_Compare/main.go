package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var line1, line2 string
	fmt.Fscan(reader, &line1, &line2)

	// fmt.Println(line1)
	// fmt.Println(line2)

	newLine1 := NewLine(line1)
	newLine2 := NewLine(line2)

	if newLine1 < newLine2 {
		fmt.Println(-1)
	} else if newLine1 == newLine2 {
		fmt.Println(0)
	} else {
		fmt.Println(1)
	}
}

func NewLine(s string) string {
	newLine := ""

	for i := 0; i < len(s); i++ {
		if s[i]%2 == 0 {
			newLine = newLine + string(s[i])
		}
	}
	return newLine
}
