package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	_ = scanner.Text() // n

	scanner.Scan()
	arr := strings.Split(scanner.Text(), " ")
	fmt.Println()
	fmt.Println(arr)

	fmt.Println(bigNumber(arr, isFirstNumBigger))
}

// функция-компаратор
func isFirstNumBigger(num1, num2 string) bool {
	x1 := num1 + num2
	x2 := num2 + num1

	n1, _ := strconv.Atoi(x1)
	n2, _ := strconv.Atoi(x2)
	return n1 >= n2
}

func bigNumber(arr []string, bigger func(string, string) bool) string {
	for i := 1; i < len(arr); i++ {
		itemToInsert := arr[i]
		j := i
		// заменим сравнение itemToInsert < array[j-1] на компаратор bigger
		for j > 0 && bigger(itemToInsert, arr[j-1]) {
			arr[j] = arr[j-1]
			j--
		}
		arr[j] = itemToInsert
	}
	fmt.Println(arr)
	fmt.Println(len(arr))
	res := strings.Join(arr, "")
	return res
}
