package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var vertices, edges int
	fmt.Fscan(reader, &vertices, &edges)

	// graph[v] содержит список ребер из вершины v
	graph := make([][]Edge, vertices)
	for i := 0; i < vertices; i++ {
		graph[i] = []Edge{}
	}

	for i := 0; i < edges; i++ {
		var from, to, weight int
		fmt.Fscan(reader, &from, &to, &weight)
		edge1 := Edge{NextNodeIdx: to, Dist: weight}
		edge2 := Edge{NextNodeIdx: from, Dist: weight}
		// граф неориентированный, добавляем туда и обратно
		// нумерация вершин по заданию начинается с единицы, а в списке с 0
		graph[from-1] = append(graph[from-1], edge1)
		graph[to-1] = append(graph[to-1], edge2)
	}

	// fmt.Println(adjacencyList)

	for i := 0; i < vertices; i++ {
		Dijkstra(i, vertices, graph)
	}

}

// началом ребра является позиция в массиве nodes[i+1]
type Edge struct {
	NextNodeIdx int
	Dist        int
}

type WayNode struct {
	NodeIdx  int
	Visited  bool
	Dist     int
	Previous int
}

// нумерация вершин по заданию начинается с единицы
func Dijkstra(vtx int, vertices int, graph [][]Edge) {
	way := make([]WayNode, vertices)

	for i := range vertices {
		dist := math.MaxInt32 // Задаём расстояние по умолчанию
		previous := -1        // Задаём предшественника для восстановления SPT
		visited := false      // Список статусов посещённости вершин

		wNode := WayNode{
			NodeIdx:  i + 1,
			Visited:  visited,
			Dist:     dist,
			Previous: previous,
		}

		way[i] = wNode
	}

	way[vtx].Dist = 0

	for {
		newMin := FindMinUnvizited(way)
		if newMin == nil {
			break
		}
		newMin.Visited = true
		// получим список исходящих вершин из вершины с индексом newMin.NodeIdx-1
		// из списка смежности
		outs := graph[newMin.NodeIdx-1]
		for _, out := range outs {
			// чтобы заменить изначальные бесконечные расстояния
			if way[out.NextNodeIdx-1].Dist > newMin.Dist+out.Dist {
				way[out.NextNodeIdx-1].Dist = newMin.Dist + out.Dist
				way[out.NextNodeIdx-1].Previous = newMin.NodeIdx
			}
		}
	}
	// fmt.Println(way)
	for i := 0; i < len(way); i++ {
		if way[i].Dist == math.MaxInt32 {
			fmt.Print(-1, " ")
		} else {
			fmt.Print(way[i].Dist, " ")
		}
	}
	fmt.Println()
}

// найдём из ещё не посещённых вершин ту, расстояние до которой минимально
func FindMinUnvizited(arr []WayNode) *WayNode {
	var currentMin *WayNode

	for i := 0; i < len(arr); i++ {
		if !arr[i].Visited {
			if currentMin == nil {
				currentMin = &arr[i]
			} else if arr[i].Dist < currentMin.Dist {
				currentMin = &arr[i]
			}
		}
	}

	return currentMin
}
