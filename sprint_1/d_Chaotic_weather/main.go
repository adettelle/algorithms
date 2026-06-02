package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)

	temperatures := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &temperatures[i])
	}

	fmt.Println(getWeatherRandomness(temperatures))
}

func getWeatherRandomness(temperatures []int) int {
	n := len(temperatures)

	if n == 1 {
		return 1
	}
	if n == 0 {
		return 0
	}

	count := 0

	if temperatures[0] > temperatures[1] {
		count++
	}
	for i := 1; i < n-1; i++ {
		if temperatures[i] > temperatures[i-1] && temperatures[i] > temperatures[i+1] {
			count++
		}
	}
	if temperatures[n-1] > temperatures[n-2] {
		count++
	}
	return count
}
