package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := makeScanner()
	scanner.Scan()
	base, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	module, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	s := scanner.Text()

	// fmt.Println(base, module)
	// fmt.Println()
	fmt.Println(getHash(s, int64(base), int64(module)))
}

// x * base^(n-i-1): при base = 709 и длинной строке значение 709^k
// растёт астрономически быстро. Уже для небольших k оно намного превышает
// максимальное значение int64 (≈ 9·10^18)
func getHash(line string, base int64, module int64) int64 {
	var h int64
	// n := len(line)
	for i := 0; i < len(line); i++ {
		// x := int64(line[i])
		// float64 начинает терять точность для больших чисел
		// и после преобразования обратно в int
		// результат вообще становится некорректным

		// вместо вычисления модуля итога, модуль берут на каждом шаге
		// sum += x * int64(math.Pow(float64(base), float64(n-i-1)))
		h = (h*base + int64(line[i])) % module
	}
	return h
}

func makeScanner() *bufio.Scanner {
	scanner := bufio.NewScanner(os.Stdin)
	maxCap := 1024 * 1024 * 1000
	buf := make([]byte, maxCap)
	scanner.Buffer(buf, maxCap)
	return scanner
}
