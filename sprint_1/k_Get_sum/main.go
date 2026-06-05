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
	readInt(scanner)
	bigNumber := readArray(scanner)
	smallNumber := readInt(scanner)

	res := getSum(bigNumber, smallNumber)
	for i := 0; i < len(res); i++ {
		fmt.Print(res[i], " ")
	}
}

func getSum(bigNumber []int, smallNumber int) []int {
	lenBig := len(bigNumber)

	secondNum := []int{}
	strSecondNum := strconv.Itoa(smallNumber)
	lenSecondNum := len(strSecondNum)
	for i := 0; i < lenSecondNum; i++ {
		x, _ := strconv.Atoi(string(strSecondNum[i]))
		secondNum = append(secondNum, x)
	}

	lenRes := 0
	if lenBig > lenSecondNum {
		lenRes = lenBig + 1
	} else {
		lenRes = lenSecondNum + 1
		bigNumber, secondNum = secondNum, bigNumber
	}

	res := make([]int, lenRes)

	startPos := lenRes - len(secondNum)
	_ = copy(res[startPos:], secondNum)

	hold := 0
	for i := lenRes - 2; i >= 0; i-- {
		x := bigNumber[i] + res[i+1] + hold
		newNum := x % 10
		hold = x / 10

		res[i+1] = newNum
	}
	res[0] = hold

	if res[0] == 0 {
		return res[1:]
	}
	return res
}

func readInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	strInt := scanner.Text()
	x, _ := strconv.Atoi(strInt)
	return x
}

func readArray(scanner *bufio.Scanner) []int {
	scanner.Scan()
	listStr := strings.Split(scanner.Text(), " ")
	arr := make([]int, len(listStr))

	for i := 0; i < len(listStr); i++ {
		arr[i], _ = strconv.Atoi(listStr[i])
	}

	return arr
}
