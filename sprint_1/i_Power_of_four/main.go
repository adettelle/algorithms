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
	strInt := scanner.Text()

	n, _ := strconv.Atoi(strInt)
	if isPowerOfFour(n) {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

}

func isPowerOfFour(number int) bool {
	const base = 4

	if number == 1 {
		return true
	}

	if number%4 != 0 {
		return false
	}

	return isPowerOfFour(number / base)
}
