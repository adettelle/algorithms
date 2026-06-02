package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	nums := make([]int, 3)

	for i := range 3 {
		fmt.Fscan(reader, &nums[i])
	}

	if checkParity(nums[0], nums[1], nums[2]) {
		fmt.Println("WIN")
	} else {
		fmt.Println("FAIL")
	}
}

// проверка соответствия
func checkParity(a int, b int, c int) bool {
	all := checkEven(a)
	if checkEven(b) == all {
		fmt.Println("b:", checkEven(b), all)
		if checkEven(c) == all {
			return true
		}
	}
	return false
}

// проверка чётности
func checkEven(x int) bool {
	if x%2 == 0 {
		return true
	}
	return false
}
