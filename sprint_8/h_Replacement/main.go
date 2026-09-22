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

	var sample string
	fmt.Fscan(reader, &sample)

	var new string
	fmt.Fscan(reader, &new)
	// fmt.Println("line:", line, "; sample:", sample, "; new:", new)

	newLine := strings.Replace(line, sample, new, -1)
	fmt.Println(newLine)
}
