package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Stack struct {
	items []int
}

func NewStack() *Stack {
	return &Stack{items: []int{}}
}

func (s *Stack) push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) pop() {
	if len(s.items) == 0 {
		fmt.Println("error")
		return
	}

	lastIndex := len(s.items) - 1
	s.items = s.items[:lastIndex]
}

func (s *Stack) getMax() {
	if len(s.items) == 0 {
		fmt.Println("None")
		return
	}
	maxItem := slices.Max(s.items)
	fmt.Println(maxItem)
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

	stack := NewStack()
	for _, command := range commands {
		switch command {
		case "pop":
			stack.pop()
		case "get_max":
			stack.getMax()
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
