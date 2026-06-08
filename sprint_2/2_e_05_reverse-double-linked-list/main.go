package main

import "fmt"

type ListNode struct {
	data string
	next *ListNode
	prev *ListNode
}

func Solution(head *ListNode) *ListNode {
	if head == nil || head.next == nil {
		return head
	}

	current := head
	var temp *ListNode
	var newHead *ListNode

	for current != nil {
		if current.next == nil {
			newHead = current
		}
		temp = current.prev
		current.prev = current.next
		current.next = temp

		current = current.prev
	}

	return newHead
}

func main() {
	node3 := ListNode{"node3", nil, nil}
	node2 := ListNode{"node2", &node3, nil}
	node1 := ListNode{"node1", &node2, nil}
	node0 := ListNode{"node0", &node1, nil}
	node3.prev = &node2
	node2.prev = &node1
	node1.prev = &node0

	printDoubleLinkedList(&node0)
	fmt.Println(" ---- ")

	newHead := Solution(&node0) //

	printDoubleLinkedList(newHead)
}

func printDoubleLinkedList(head *ListNode) {
	node := head

	for node != nil {
		fmt.Print(node.data, "<->")
		node = node.next
	}
}

func test() {
	node3 := ListNode{"node3", nil, nil}
	node2 := ListNode{"node2", &node3, nil}
	node1 := ListNode{"node1", &node2, nil}
	node0 := ListNode{"node0", &node1, nil}
	node3.prev = &node2
	node2.prev = &node1
	node1.prev = &node0
	/*newHead :=*/ Solution(&node0)
	// result is : newHead == node3
	// node3.next == node2
	// node2.next == node1
	// node2.prev = node3
	// node1.next == node0
	// node1.prev == node2
	// node0.prev == node1
}
