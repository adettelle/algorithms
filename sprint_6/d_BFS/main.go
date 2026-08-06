package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const (
	white = iota
	gray
	black
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
		var from, to int
		fmt.Fscan(reader, &from, &to)
		// т.к. граф неориентированный добавляем ребро два раза
		adjacencyList[from-1] = append(adjacencyList[from-1], to)
		adjacencyList[to-1] = append(adjacencyList[to-1], from)
	}

	for i := 0; i < vertices; i++ {
		sort.Ints(adjacencyList[i])
	}

	var startVtx int
	fmt.Fscan(reader, &startVtx)

	color := make([]int, vertices)
	for i := 0; i < vertices; i++ {
		color[i] = white
	}

	res := BFS(startVtx-1, color, adjacencyList)
	for i := 0; i < len(res); i++ {
		fmt.Print(res[i]+1, " ")
	}
}

// Функция BFS обходит граф по уровням
func BFS(start int, color []int, adjacencyList [][]int) []int {
	visited := make(map[int]bool)
	// Создадим очередь вершин и положим туда стартовую вершину.
	queue := []int{start}
	visited[start] = true

	var result []int

	for len(queue) > 0 {
		// Извлекаем первый элемент из очереди
		node := queue[0]
		queue = queue[1:]
		result = append(result, node)

		// Проходим по всем соседям вершины
		for _, neighbor := range adjacencyList[node] {
			if !visited[neighbor-1] {
				visited[neighbor-1] = true
				queue = append(queue, neighbor-1)
			}
		}
	}
	return result
}
