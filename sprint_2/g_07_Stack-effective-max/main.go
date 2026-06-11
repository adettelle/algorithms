package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 3 -> 2 -> 1 -> nil
type StackMaxEffective struct {
	head *Node
	// maxValue int
}

type Node struct {
	value    int
	next     *Node
	maxValue int
}

func NewStack() *StackMaxEffective {
	return &StackMaxEffective{head: nil}
}

func (s *StackMaxEffective) push(value int) {
	newNode := Node{value: value, next: s.head}

	if s.head == nil {
		newNode.maxValue = value
	} else if value > s.head.maxValue {
		newNode.maxValue = value
	} else {
		newNode.maxValue = s.head.maxValue
	}

	s.head = &newNode
}

func (s *StackMaxEffective) pop() {
	if s.head == nil {
		fmt.Println("error")
		return
	}
	s.head = s.head.next

}

func (s *StackMaxEffective) getMax() {
	if s.head == nil {
		fmt.Println("None")
		return
	}
	fmt.Println(s.head.maxValue)
}

func (s *StackMaxEffective) top() {
	if s.head == nil {
		fmt.Println("error")
		return
	}
	fmt.Println(s.head.value)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	strInt := scanner.Text()
	n, _ := strconv.Atoi(strInt)

	commands := make([]string, n)

	for i := range n {
		commands[i] = readLine(scanner)
	}

	// var maxValue int
	stack := NewStack()
	for _, command := range commands {
		switch command {
		case "pop":
			stack.pop()
		case "get_max":
			stack.getMax()
		case "top":
			stack.top()
		default:
			args := strings.Split(command, " ")
			argument, _ := strconv.Atoi(args[1])
			stack.push(argument)
		}
	}
}

func readLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	return scanner.Text()
}
