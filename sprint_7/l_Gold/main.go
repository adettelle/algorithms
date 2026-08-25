package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, capacity int
	fmt.Fscan(reader, &n, &capacity)

	weights := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &weights[i])
	}

	// fmt.Println(weights)
	fmt.Println(FindTotalWeight(weights, capacity, n))
}

func FindTotalWeight(weights []int, capacity int, n int) int {
	// r — это количество рассматриваемых слитков
	//  с — вместимость рюкзака
	// dp[r][c] - максимальный вес золота, который можно получить,
	// если рассматривать первые r слитков и иметь рюкзак вместимостью c кг.
	dp := make([]int, capacity+1)

	for r := 0; r < n; r++ {
		w := weights[r]
		// Идём справа налево
		for c := capacity; c >= w; c-- {
			dp[c] = max(dp[c], dp[c-w]+w)

		}
	}

	return dp[capacity]
}
