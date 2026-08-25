package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod int = 1000000007

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)
	fmt.Println(Fibonacci(n))
}

func Fibonacci(n int) int {
	m := n + 1
	arr := make([]int, m)

	arr[0] = 1
	arr[1] = 1
	for i := 2; i < m; i++ {
		arr[i] = (arr[i-1] + arr[i-2]) % mod
	}
	return arr[n]
}
