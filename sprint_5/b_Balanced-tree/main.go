package main

import (
	"fmt"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func Solution(root *Node) bool {
	return check(root) != -1
}

func check(root *Node) int {
	if root == nil {
		return 0
	}

	leftHeight := check(root.left)
	if leftHeight == -1 {
		return -1
	}

	rightHeight := check(root.right)
	if rightHeight == -1 {
		return -1
	}

	if abs(leftHeight-rightHeight) > 1 {
		return -1
	}

	if leftHeight > rightHeight {
		return leftHeight + 1
	}

	return rightHeight + 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func main() {
	node0 := Node{10, nil, nil}   //
	node1 := Node{1, &node0, nil} // node1 := Node{1, nil, nil}
	node2 := Node{-5, nil, nil}
	node3 := Node{3, &node1, &node2}
	node4 := Node{10, nil, nil}
	node5 := Node{2, &node3, &node4}
	if !Solution(&node5) {
		fmt.Println("Tree is not balanced")
	}
}

//         2(5)
//       /   \
//	   3(3)     10(4)
//	  /   \
//	 1(1)  -5(2)
// true

//         2(5)
//       /   \
//	   3(3)     10(4)
//	  /   \
//	 1(1)  -5(2)
//   /
//  10(0)
// flase
