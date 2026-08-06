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
		var from, to int
		fmt.Fscan(reader, &from, &to)
		// т.к. граф неориентированный добавляем ребро два раза
		adjacencyList[from-1] = append(adjacencyList[from-1], to)
		adjacencyList[to-1] = append(adjacencyList[to-1], from)
	}

	color := make([]int, vertices)

	for i := 0; i < vertices; i++ {
		color[i] = -1
	}

	components := make([][]int, 0)
	component := 1

	for i := 0; i < vertices; i++ {
		if color[i] == -1 {
			current := make([]int, 0)
			DFS(i, component, color, adjacencyList, &current)
			components = append(components, current)
			component++
		}
	}

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fprintln(writer, len(components))

	for _, comp := range components {
		sort.Ints(comp)
		for _, v := range comp {
			fmt.Fprint(writer, v, " ")
		}
		fmt.Fprintln(writer)
	}
}

// DFS принимает слайс текущей компоненты vertices *[]int
func DFS(v int, component int, color []int, adjacencyList [][]int, vertices *[]int) {
	color[v] = component
	*vertices = append(*vertices, v+1)

	// Получим список исходящих ребер в зависимости от способа хранения графа
	for _, val := range adjacencyList[v] {
		if color[val-1] == -1 { // Если вершина не посещена, то
			// запустим обход от найденной смежной вершины.
			DFS(val-1, component, color, adjacencyList, vertices)
		}
	}
}
