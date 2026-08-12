// https://contest.yandex.ru/contest/25070/run-report/164314368/

/*
-- ПРИНЦИП РАБОТЫ --
Считываем граф и представляем его в виде списка смежности.
Поскольку граф неориентированный, каждое ребро добавляется в список смежности
обеих вершин.
Для поиска максимального остовного дерева используется алгоритм Прима.
Сначала все вершины помечаем как недобаленные в остов. Выбираем произвольную
стартовую вершину и добавляем её в остов. Все рёбра, выходящие из неё
в ещё не добавленные вершины, помещаем в максимальную кучу.

Пока в массиве недобавленных вершин есть ещё вершины и пока куча не пустая,
берем из кучи ребро максимального веса. Если его конечная вершина уже принадлежит
остову, пропускаем это ребро. Иначе добавляем ребро в остов,
включаем новую вершину в остов и помещаем в кучу все рёбра,
ведущие из неё в ещё не добавленные вершины.
После завершения алгоритма проверяем, остались ли недобавленные вершины.
Если да, граф несвязный, поэтому остовное дерево не существует
и выводится строка Oops! I did it again. Иначе вычисляется сумма весов рёбер
найденного остовного дерева и выводится ответ.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
На каждом шаге алгоритм поддерживает список уже добавленных в остов вершин
и рассматривает все рёбра, соединяющие эти вершины с остальными вершинами графа.
Из этих рёбер выбирается ребро максимального веса.

Каждое добавленное ребро соединяет уже построенную часть остова с новой вершиной,
поэтому циклы не образуются. После добавления всех достижимых вершин получается
дерево, содержащее все вершины связного графа.

Если после завершения алгоритма остались недобавленные вершины, это значит,
что из уже построенной части графа они недостижимы, следовательно, граф несвязный.
В этом случае остовного дерева не существует.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Пусть V - количество вершин и E - количество ребер.
Каждая вершина добавляется в остов ровно один раз. Каждое ребро не более одного
раза помещается в кучу с каждой стороны, поэтому общее число операций добавления
и извлечения из кучи пропорционально E.

Каждое ребро не более одного раза добавляется в кучу (heap.Push)
и не более одного раза извлекается из неё (heap.Pop). Операции добавления
и извлечения из бинарной кучи выполняются за O(logE), так как в худшем случае
куча содержит O(E) рёбер.

Следовательно, общая временная сложность алгоритма составляет O(ElogE).

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Пусть V - количество вершин и E - количество ребер.
Список смежности хранит каждое ребро дважды (поскольку
граф неориентированный) - O(E).
Множества added и notAdded содержат по одной записи для каждой вершины — O(V).
Массив рёбер остовного дерева содержит не более V−1 рёбер — O(V).
Максимальная куча в худшем случае может содержать O(E) рёбер.

Таким образом пространственная сложность составляет O(V+E).
*/
package main

import (
	"bufio"
	"container/heap"
	"fmt"
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

	// все вершины внутри программы имеют номера 0..n-1
	for i := 0; i < edges; i++ {
		var from, to, weight int
		fmt.Fscan(reader, &from, &to, &weight)
		from--
		to--

		edge1 := Edge{From: from, To: to, Weight: weight}
		edge2 := Edge{From: to, To: from, Weight: weight}
		// граф неориентированный, добавляем туда и обратно
		graph[from] = append(graph[from], edge1)
		graph[to] = append(graph[to], edge2)
	}

	fmt.Println("graph:", graph, "; len:", len(graph))

	if len(graph) == 1 && len(graph[0]) == 0 {
		fmt.Println(0)
		return
	}

	mst := findMSP(vertices, graph)
	if mst == nil {
		fmt.Println("Oops! I did it again")
		return
	}
	fmt.Println(mst)

	totalWeight := 0

	for _, e := range mst {
		totalWeight += e.Weight
	}

	fmt.Println("Weight:", totalWeight)
}

type Edge struct {
	From   int
	To     int
	Weight int
}

type MaxHeap []Edge

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	// максимум сверху
	return h[i].Weight > h[j].Weight
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x any) {
	*h = append(*h, x.(Edge))
}

func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)

	x := old[n-1]
	*h = old[:n-1]

	return x
}

func addVertex(v int, added, notAdded map[int]bool, h *MaxHeap, graph [][]Edge) {
	added[v] = true
	delete(notAdded, v)

	for _, edge := range graph[v] {
		if notAdded[edge.To] {
			heap.Push(h, edge)
		}
	}
}

// maximum spanning tree (MST) Prym algorithm
func findMSP(vertices int, graph [][]Edge) []Edge {
	var maximumSpanningTree []Edge // Рёбра, составляющие MST.

	added := make(map[int]bool)    // Множество вершин, уже добавленных в остов.
	notAdded := make(map[int]bool) // Множество вершины, ещё не добавленных в остов.

	for i := 0; i < vertices; i++ {
		notAdded[i] = true
	}

	h := &MaxHeap{}
	heap.Init(h)
	v := 0
	addVertex(v, added, notAdded, h, graph)

	// Основной цикл алгоритма Прима
	for len(notAdded) > 0 && h.Len() > 0 {
		edge := heap.Pop(h).(Edge)

		if !notAdded[edge.To] {
			continue
		}
		maximumSpanningTree = append(maximumSpanningTree, edge)
		addVertex(edge.To, added, notAdded, h, graph)
	}

	if len(notAdded) != 0 {
		// граф несвязный
		return nil
	}

	return maximumSpanningTree
}
