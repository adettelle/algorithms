package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

func Solution(root *Node) bool {
	return validate(root, nil, nil)
}

func validate(node *Node, minNode *Node, maxNode *Node) bool {
	if node == nil {
		return true
	}

	// Значение текущего узла должно быть больше minNode.Val
	if minNode != nil && node.value <= minNode.value {
		return false
	}

	// Значение текущего узла должно быть меньше maxNode.Val
	if maxNode != nil && node.value >= maxNode.value {
		return false
	}

	// Рекурсивно проверяем левое и правое поддерево
	return validate(node.left, minNode, node) && validate(node.right, node, maxNode)
}

// вспомогательная структура для хранения id потомков
type Children struct {
	leftID  int
	rightID int
}

func main() {

	node1 := Node{1, nil, nil}
	node2 := Node{4, nil, nil}
	node3 := Node{3, &node1, &node2}
	node4 := Node{8, nil, nil}
	node5 := Node{5, &node3, &node4}

	if !Solution(&node5) {
		fmt.Println("False")
	} else {
		fmt.Println("True")
	}

	node2.value = 5
	if !Solution(&node5) {
		fmt.Println("False")
	} else {
		fmt.Println("True")
	}

	reader := bufio.NewReader(os.Stdin)

	var n int
	fmt.Fscan(reader, &n)

	tree := createTree(reader, n)
	// корень имеет id = 0
	root := &tree[0]

	fmt.Println("Корень:", root.value)

	if root.left != nil {
		fmt.Println("Левый потомок:", root.left.value)
	}

	if root.right != nil {
		fmt.Println("Правый потомок:", root.right.value)
	}

	if !Solution(root) {
		fmt.Println("False")
	} else {
		fmt.Println("True")
	}
}

func createTree(reader *bufio.Reader, n int) []Node {
	tree := make([]Node, n)

	// здесь будем хранить номера потомков
	children := make([]Children, n)

	for i := 0; i < n; i++ {
		var (
			id      int
			value   int
			leftID  int
			rightID int
		)

		fmt.Fscan(
			reader,
			&id,
			&value,
			&leftID,
			&rightID,
		)

		tree[id].value = value
		children[id] = Children{leftID, rightID}
	}

	// связываем вершины указателями
	for i := 0; i < n; i++ {
		if children[i].leftID != -1 {
			tree[i].left = &tree[children[i].leftID]
		}

		if children[i].rightID != -1 {
			tree[i].right = &tree[children[i].rightID]
		}
	}

	return tree
}
