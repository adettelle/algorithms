package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	clubs := make(map[string]int)

	for i := range n {
		scanner.Scan()
		line := scanner.Text()
		if _, ok := clubs[line]; !ok {
			clubs[line] = i
		}
	}

	// fmt.Println(clubs)

	type keyValue struct {
		Key   string
		Value int
	}
	var sortedStruct []keyValue

	for key, val := range clubs {
		sortedStruct = append(sortedStruct, keyValue{Key: key, Value: val})
	}

	sort.Slice(sortedStruct, func(i, j int) bool {
		return sortedStruct[i].Value < sortedStruct[j].Value
	})

	for _, elem := range sortedStruct {
		fmt.Println(elem.Key)
	}
}
