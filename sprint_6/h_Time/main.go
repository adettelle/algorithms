package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
		var vertexFrom, vertexTo int
		fmt.Fscan(reader, &vertexFrom, &vertexTo)

		adjacencyList[vertexFrom-1] = append(adjacencyList[vertexFrom-1], vertexTo-1)
	}

	for i := 0; i < vertices; i++ {
		sort.Ints(adjacencyList[i])
	}

	color := make([]string, vertices)

	for i := 0; i < vertices; i++ {
		color[i] = "white"
	}

	var timer int
	entry := make([]int, vertices)
	leave := make([]int, vertices)

	fmt.Println()
	DFS(0, &timer, entry, leave, color, adjacencyList)

	for i := 0; i < len(entry); i++ {
		fmt.Println(entry[i], leave[i])
	}
}

// v - номер вершины
func DFS(v int, timer *int, entry, leave []int, color []string, adjacencyList [][]int) {
	entry[v] = *timer
	*timer++
	color[v] = "gray"

	for _, w := range adjacencyList[v] {
		if color[w] == "white" {
			DFS(w, timer, entry, leave, color, adjacencyList)
		}
	}

	leave[v] = *timer
	*timer++
	color[v] = "black"
}
