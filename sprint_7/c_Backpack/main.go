package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var capacity int
	fmt.Fscan(reader, &capacity)

	var n int
	fmt.Fscan(reader, &n)

	heaps := make([]Heap, n)
	for i := 0; i < n; i++ {
		var price, weight int
		fmt.Fscan(reader, &price, &weight)
		heap := Heap{price: price, weight: weight}
		heaps[i] = heap
	}
	// fmt.Println(heaps)

	sort.Slice(heaps, func(i, j int) bool {
		return Less(heaps[i], heaps[j])
	})

	fmt.Println(heaps)
	fmt.Println(CalcBackpack(heaps, n, capacity))
}

type Heap struct {
	price  int
	weight int
}

func Less(h1, h2 Heap) bool {
	if h1.price == h2.price {
		return h1.weight > h2.weight
	}

	return h1.price > h2.price
}

func CalcBackpack(heaps []Heap, n int, capacity int) int {
	res := 0

	for i := 0; i < n; i++ {
		if capacity-heaps[i].weight >= 0 {
			res += heaps[i].price * heaps[i].weight
			// fmt.Println("res:", res)
		} else {
			res += heaps[i].price * capacity
			return res
		}
		capacity -= heaps[i].weight

	}
	return res
}
