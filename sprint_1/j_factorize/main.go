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
	res := factorize(n)
	for _, elem := range res {
		fmt.Print(elem, " ")
	}
}

func factorize(number int) []int {
	res := []int{}

	for divider := 2; divider <= number; {
		remainder := number % divider

		if remainder != 0 {
			divider++
			continue
		} else {
			res = append(res, divider)
			number /= divider
		}
	}
	return res
}
