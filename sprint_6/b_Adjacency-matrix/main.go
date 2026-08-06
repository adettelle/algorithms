package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var edges, vertices int
	fmt.Fscan(reader, &vertices, &edges)

	adjacencyMatrix := make([][]int, vertices)

	for i := 0; i < vertices; i++ {
		adjacencyMatrix[i] = make([]int, vertices)
	}

	for i := 0; i < edges; i++ {
		var col, row int
		fmt.Fscan(reader, &col, &row)
		adjacencyMatrix[col-1][row-1] = 1
	}
	for _, row := range adjacencyMatrix {
		for _, val := range row {
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}
