package main

import "fmt"

type ListNode struct {
	data string
	next *ListNode
}

func Solution(head *ListNode, elem string) int {
	if head == nil {
		return -1
	}

	counter := 0

	for node := head; node != nil; node = node.next {
		if node.data == elem {
			return counter
		}
		counter++
	}

	return -1
}

func main() {
	node3 := ListNode{"node3", nil}
	node2 := ListNode{"node2", &node3}
	node1 := ListNode{"node1", &node2}
	node0 := ListNode{"node0", &node1}
	idx := Solution(&node0, "node5")
	fmt.Println(idx)
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
	/*idx :=*/ Solution(&node0, "node2")
	// result is : idx == 2
}
