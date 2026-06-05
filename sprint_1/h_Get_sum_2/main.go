package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	firstNumber := readLine(scanner)
	secondNumber := readLine(scanner)

	res := getSum(firstNumber, secondNumber)
	fmt.Println(res)
}

func getSum(firstNumber string, secondNumber string) string {
	if len(firstNumber) < len(secondNumber) {
		firstNumber, secondNumber = secondNumber, firstNumber
	}

	lenRes := len(firstNumber) + 1

	long := []int{}
	for i := 0; i < len(firstNumber); i++ {
		x, _ := strconv.Atoi(string(firstNumber[i]))
		long = append(long, x)
	}

	short := []int{}
	for i := 0; i < len(secondNumber); i++ {
		x, _ := strconv.Atoi(string(secondNumber[i]))
		short = append(short, x)
	}

	res := make([]int, lenRes)
	startPos := lenRes - len(short)
	_ = copy(res[startPos:], short)

	hold := 0
	for i := len(long) - 1; i >= 0; i-- {
		x := long[i] + res[i+1] + hold
		newNum := x % 2
		hold = x / 2

		res[i+1] = newNum
	}

	res[0] = hold

	if res[0] == 0 {
		res = res[1:]
	}

	result := make([]byte, len(res))
	for i, elem := range res {
		result[i] = '0' + byte(elem)
	}

	return string(result)
}

func readLine(scanner *bufio.Scanner) string {
	scanner.Scan()
	line := scanner.Text()

	return line
}
