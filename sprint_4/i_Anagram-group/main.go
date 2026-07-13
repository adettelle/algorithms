package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n int
	fmt.Fscan(reader, &n)

	words := make([]string, n)
	for i := range n {
		fmt.Fscan(reader, &words[i])
	}

	/*
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		_, _ = strconv.Atoi(scanner.Text()) // n

		scanner.Scan()
		words := strings.Split(scanner.Text(), " ")
	*/

	res := findAnagrams(words, len(words)) // n
	for _, sl := range res {
		for _, x := range sl { // j := 0; j < len(sl); j++
			// fmt.Print(x, " ") // sl[j]
			fmt.Fprint(writer, x, " ")
		}
		fmt.Fprintln(writer)
	}
}

func findAnagrams(words []string, n int) [][]int {
	anagrams := make(map[string][]int)

	for i := 0; i < n; i++ {
		sortedWord := []byte(words[i])

		sort.Slice(sortedWord, func(k, m int) bool {
			return sortedWord[k] < sortedWord[m]
		})

		key := string(sortedWord)
		anagrams[key] = append(anagrams[string(sortedWord)], i)
	}

	res := [][]int{}
	for _, v := range anagrams {
		res = append(res, v)
	}
	slices.SortFunc(res, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})

	// fmt.Println()
	return res
}
