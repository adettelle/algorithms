package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var edges, vertices int
	fmt.Fscan(reader, &vertices, &edges)

	adjacencyList := make([][]int, vertices)
	for i := 0; i < vertices; i++ {
		adjacencyList[i] = []int{}
	}

	for i := 0; i < edges; i++ {
		var vertexFrom, vertexTo int
		fmt.Fscan(reader, &vertexFrom, &vertexTo)

		adjacencyList[vertexFrom-1] = append(adjacencyList[vertexFrom-1], vertexTo)
	}

	for i := 0; i < vertices; i++ {
		fmt.Print(len(adjacencyList[i]))
		sort.Ints(adjacencyList[i])
		for _, vertex := range adjacencyList[i] {
			fmt.Print(" ", vertex)
		}
		fmt.Println()
	}
}
