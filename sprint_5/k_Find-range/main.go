package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func printRange(root *Node, left int, right int) {
	arr := []int{}
	res := FindInRange(root, left, right, &arr)
	for i := 0; i < len(res); i++ {
		fmt.Print(res[i], " ")
	}
}

// FindInRange рекурсивно ищет элементы в интервале [min, max]
func FindInRange(root *Node, left, right int, result *[]int) []int {
	if root == nil {
		return nil
	}

	if root.value >= left {
		FindInRange(root.left, left, right, result)
	}

	// Если значение узла в интервале, добавляем в результат
	if root.value >= left && root.value <= right {
		*result = append(*result, root.value)
	}

	// Если текущий узел < max, исследуем правое поддерево
	if root.value <= right {
		FindInRange(root.right, left, right, result)
	}
	return *result
}

func main() {
	node1 := Node{2, nil, nil}
	node2 := Node{1, nil, &node1}
	node3 := Node{8, nil, nil}
	node4 := Node{8, nil, &node3}
	node5 := Node{9, &node4, nil}
	node6 := Node{10, &node5, nil}
	node7 := Node{5, &node2, &node6}
	printRange(&node7, 2, 8)
	// expected output: 2 5 8 8
}

// Пример дерева:
//         5(7)
//       /      \
//	   1(2)      10(6)
//	    \        /
//	     2(1)   9(5)
//             /
//           8(4)
//            \
//             8(3)
