package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func Solution(root *Node) bool {
	if root == nil {
		return true
	}

	return isMirrow(root.left, root.right)
}

func isMirrow(left, right *Node) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	if left.value != right.value {
		return false
	}

	return isMirrow(left.left, right.right) &&
		isMirrow(left.right, right.left)
}

func main() {
	node1 := Node{5, nil, nil}
	node2 := Node{4, nil, nil}
	node3 := Node{4, nil, nil}
	node4 := Node{3, nil, nil}
	node5 := Node{2, &node1, &node2}
	node6 := Node{2, &node3, &node4}
	node7 := Node{1, &node5, &node6}

	// if !Solution(&node7) {
	// 	panic("WA")
	// }

	if !Solution(&node7) {
		fmt.Println("False")
	} else {
		fmt.Println("True")
	}
}

// Пример дерева:
//
//         1(7)
//	    /       \
//	  2(5)      2(6)
//    /  \      /   \
//	3(1) 4(2) 4(3)   3(4)
