package main

import (
	"bufio"
	"fmt"
	"os"
)

const mod int = 1000000007

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscan(reader, &n, &k)
	fmt.Println(Stairs(n, k))
}

func Stairs(n, k int) int {
	arr := make([]int, n+1)
	arr[0] = 1

	for i := 1; i < n+1; i++ {
		for j := 1; j < k+1; j++ {
			if i-j >= 0 {
				arr[i] = (arr[i] + arr[i-j]) % mod
			}
		}
	}
	// fmt.Println(arr)
	return arr[n-1]
}
