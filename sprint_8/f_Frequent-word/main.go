package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)

	words := make(map[string]int)

	for i := 0; i < n; i++ {
		var word string
		fmt.Fscan(reader, &word)
		// if _, ok := words[word]; !ok {
		// }
		words[word]++
	}

	var sortedStruct []KeyValue

	for key, val := range words {
		sortedStruct = append(sortedStruct, KeyValue{Key: key, Value: val})
	}

	// sort.Slice(sortedStruct, func(i, j int) bool {
	// 	return sortedStruct[i].Value > sortedStruct[j].Value
	// })

	sort.Slice(sortedStruct, func(i, j int) bool {
		if sortedStruct[i].Value == sortedStruct[j].Value {
			return sortedStruct[i].Key < sortedStruct[j].Key
		}
		return sortedStruct[i].Value > sortedStruct[j].Value
	})
	fmt.Println(sortedStruct[0].Key)
}

type KeyValue struct {
	Key   string
	Value int
}
