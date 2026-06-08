package main

import "fmt"

type ListNode struct {
	data string
	next *ListNode
}

func Solution(head *ListNode, idx int) *ListNode {
	if head == nil {
		return nil
	}

	if idx == 0 {
		return head.next
	}

	currentNode := head

	for i := 0; currentNode != nil && i < idx-1; i++ {
		currentNode = currentNode.next
	}

	if currentNode == nil || currentNode.next == nil {
		return head
	}

	currentNode.next = currentNode.next.next

	return head
}

func PrintLinkedList(head *ListNode) {
	node := head

	for node.next != nil {
		fmt.Println(node.data)
		node = node.next
	}
	fmt.Println(node.data)
}

func test() {
	node3 := ListNode{"node3", nil}
	node2 := ListNode{"node2", &node3}
	node1 := ListNode{"node1", &node2}
	node0 := ListNode{"node0", &node1}
	/*newHead :=*/ Solution(&node0, 1)
	// result is : node0 -> node2 -> node3
}

func main() {
	node3 := ListNode{"node3", nil}
	node2 := ListNode{"node2", &node3}
	node1 := ListNode{"node1", &node2}
	node0 := ListNode{"node0", &node1}
	newHead := Solution(&node0, 4)
	PrintLinkedList(newHead)
}
