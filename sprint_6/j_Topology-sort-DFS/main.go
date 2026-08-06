package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	White = "white" // iota
	Gray  = "gray"
	Black = "balck"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var vertices, edges int

	fmt.Fscan(reader, &vertices, &edges)

	adjacencyList := make([][]int, vertices)
	for i := 0; i < vertices; i++ {
		adjacencyList[i] = []int{}
	}

	for i := 0; i < edges; i++ {
		var from, to int
		fmt.Fscan(reader, &from, &to)
		adjacencyList[from-1] = append(adjacencyList[from-1], to)
	}

	// for i := 0; i < vertices; i++ {
	// 	fmt.Print("vertex:", i+1, "; num of edges:", len(adjacencyList[i]))
	// 	// sort.Ints(adjacencyList[i])
	// 	for _, vertex := range adjacencyList[i] {
	// 		fmt.Print(" ", vertex)
	// 	}
	// 	fmt.Println()
	// }

	color := make([]string, vertices)

	for i := 0; i < vertices; i++ {
		color[i] = White
	}

	order := make([]int, 0, vertices)

	for i := 0; i < len(color); i++ {
		if color[i] == White {
			topSort(i, color, adjacencyList, &order)
		}
	}

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for i := len(order) - 1; i >= 0; i-- {
		fmt.Fprint(writer, order[i]+1, " ")
	}
}

func topSort(v int, color []string, adjacencyList [][]int, order *[]int) {
	color[v] = Gray

	outGoingEdges := adjacencyList[v]
	// fmt.Println("outGoingEdges:", outGoingEdges)

	for _, to := range outGoingEdges {
		// fmt.Println("color[val-1]:", color[to-1])
		if color[to-1] == White {
			topSort(to-1, color, adjacencyList, order)
		}
	}

	color[v] = Black
	*order = append(*order, v)
}
