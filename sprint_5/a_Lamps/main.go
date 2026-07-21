package main

import (
	"fmt"
	"math"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func Solution(root *Node) int {
	// если дерево пустое, возвращаем минимально возможное число
	if root == nil {
		return math.MinInt
	}

	max := root.value

	leftMax := Solution(root.left)
	if leftMax > max {
		max = leftMax
	}

	rightMax := Solution(root.right)
	if rightMax > max {
		max = rightMax
	}

	return max
}

// Пример дерева:
//
//	  10
//	 /  \
//	5    25
//	    /  \
//	   12   27
func main() {
	root := &Node{value: 10}
	root.left = &Node{value: 5}
	root.right = &Node{value: 25}
	root.right.left = &Node{value: 12}
	root.right.right = &Node{value: 27}

	fmt.Println(Solution(root)) // 25
}
