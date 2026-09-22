package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	words := strings.Fields(input)

	for i := len(words) - 1; i >= 0; i-- {
		fmt.Print(words[i], " ")
	}
}
