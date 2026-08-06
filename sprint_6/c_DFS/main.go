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
	fmt.Println()
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

	// после чтения графа сортируем каждый список смежности, т.к. по заданию
	// соседи будут рассматриваться в порядке возрастания
	// in [[4 2] [3 1] [2 4] [3 1]]
	// out [[2 4] [1 3] [2 4] [1 3]]
	for i := 0; i < vertices; i++ {
		sort.Ints(adjacencyList[i])
	}

	var startVtx int
	fmt.Fscan(reader, &startVtx)

	// fmt.Println(adjacencyList)
	// fmt.Println("start vertex:", startVtx)

	color := make([]string, vertices)
	for i := 0; i < vertices; i++ {
		color[i] = "white"
	}

	DFS(startVtx-1, vertices, adjacencyList, color)
}

// v - номер вершины
func DFS(v int, vertices int, adjacencyList [][]int, color []string) {
	color[v] = "gray"
	fmt.Print(v+1, " ")

	// Получим список исходящих ребер в зависимости от способа хранения графа
	for _, next := range adjacencyList[v] {
		if color[next-1] == "white" { // Если вершина не посещена, то
			// запустим обход от найденной смежной вершины.
			DFS(next-1, vertices, adjacencyList, color)
		}
	}
	color[v] = "black" // Теперь вершина обработана.
}
