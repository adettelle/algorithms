package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	n, _ := strconv.Atoi(str)
	open := 0
	close := 0
	bracketGen(n, "", open, close)
}

func bracketGen(n int, sequence string, open, close int) {
	if len(sequence) == 2*n {
		fmt.Println(sequence)
		return
	}

	if open < n {
		bracketGen(n, sequence+"(", open+1, close)
	}
	if close < open {
		bracketGen(n, sequence+")", open, close+1)
	}
}
