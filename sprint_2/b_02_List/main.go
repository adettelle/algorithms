package main

import "fmt"

type ListNode struct {
	data string
	next *ListNode
}

func main() {
	node3 := ListNode{"node3", nil}
	node2 := ListNode{"node2", &node3}
	node1 := ListNode{"node1", &node2}
	node0 := ListNode{"node0", &node1}
	Solution(&node0)
}

func Solution(head *ListNode) {
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
	Solution(&node0)
	/*
	   Output is:
	   node0
	   node1
	   node2
	   node3
	*/
}
