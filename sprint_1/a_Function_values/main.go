package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	// y = a * x^2 + b * x + c
	arr := make([]int, 4) // a, x, b, c

	for i := range 4 {
		fmt.Fscan(reader, &arr[i])
	}

	res := evaluateFunction(arr[0], arr[1], arr[2], arr[3])

	fmt.Println(res)

}

func evaluateFunction(a, x, b, c int) int {
	res := a*x*x + b*x + c

	return res
}
