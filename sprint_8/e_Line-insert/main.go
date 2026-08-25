package main

import (
	"bufio"
	"fmt"
	"os"
)

const lenTi = 100000 // суммарная длина всех строк ti не превосходит 10^5.

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var s string
	fmt.Fscan(reader, &s)

	var n int
	fmt.Fscan(reader, &n)

	// insert[k] — строка, которую нужно вставить
	// после k-го символа исходной строки.
	insert := make([]string, len(s)+1)

	for i := 0; i < n; i++ {
		var word string
		var idx int
		fmt.Fscan(reader, &word, &idx)

		insert[idx] = word
	}

	res := make([]byte, 0, len(s)+lenTi)

	res = append(res, insert[0]...)

	for i := 0; i < len(s); i++ {
		res = append(res, s[i])
		res = append(res, insert[i+1]...)
	}
	fmt.Fprintln(writer, string(res))
}

/*
// time-limit-exceeded
func main() {
	reader := bufio.NewReader(os.Stdin)
	var s string
	fmt.Fscan(reader, &s)

	var n int
	fmt.Fscan(reader, &n)

	data := make([]Input, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &data[i].Word, &data[i].Idx)
	}

	sort.Slice(data, func(i, j int) bool {
		return data[i].Idx > data[j].Idx
	})

	for i := 0; i < n; i++ {
		s = s[:data[i].Idx] + data[i].Word + s[data[i].Idx:]
	}
	fmt.Println(s)
}

type Input struct {
	Word string
	Idx  int
}
*/
