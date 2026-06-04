package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	strInt := scanner.Text()
	n, _ := strconv.Atoi(strInt)

	res := getBinaryNumber(n)
	resStr := []string{}
	for _, elem := range res {
		x := strconv.Itoa(elem)
		resStr = append(resStr, x)
	}
	s := strings.Join(resStr, "")

	fmt.Println(s)
}

func getBinaryNumber(n int) []int {
	if n/2 == 0 {
		return []int{n % 2}
	}

	return append(getBinaryNumber(n/2), n%2)
}

func getBinaryNumber2(n int) []int {
	if n == 0 {
		return []int{0}
	}

	res := []int{}

	for curr := n; curr > 0; curr /= 2 {
		res = append(res, curr%2)
	}
	slices.Reverse(res)

	return res
}

func getBinaryNumber3(n int) []int {
	if n == 0 {
		return []int{0}
	}

	res := []int{}

	// curr&1 - побитовый and с единицей
	// 1 0 0 1 1
	//         1

	// curr >>= 1 деление на 2 нацело (n/=2)
	// 19         [1 0 0 1 1]
	// 19 / 2 = 9 [1 0 0 1]  (1)
	// 9  / 2 = 4  [1 0 0]   (1)
	// 4  / 2 = 2  [1 0]     (0)
	// 2  / 2 = 1  [1]		 (0)
	// 1  / 2 = 0  [0]		 (1)
	for curr := n; curr > 0; curr >>= 1 { // curr >>= 1 деление на 2 нацело (n/=2)
		res = append(res, curr%2) // или curr&1 - побитовый and с единицей
	}
	slices.Reverse(res)

	return res
}
