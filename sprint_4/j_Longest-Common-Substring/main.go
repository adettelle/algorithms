// Итоговая сложность: O((n+m)log(min(n,m)))
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(reader, &n)
	arr1 := make([]int, n)

	for i := range n {
		fmt.Fscan(reader, &arr1[i])
	}

	var m int
	fmt.Fscan(reader, &m)

	arr2 := make([]int, m)

	for i := range m {
		fmt.Fscan(reader, &arr2[i])
	}

	// fmt.Println(arr1, arr2)
	// fmt.Println(findMaxSubArray(arr1, arr2))

	pow := buildPow(min(n, m))
	hashA := buildHash(arr1)
	hashB := buildHash(arr2)

	res := binarySearch(arr1, arr2, hashA, hashB, pow)
	fmt.Println(res)
}

// значения элементов лежат от 0 до 255; 257 больше любого возможного элемента
const P uint64 = 257

// строит массив префиксных хешей
// O(n + m)
func buildHash(a []int) []uint64 {
	hash := make([]uint64, len(a)+1)

	for i := 0; i < len(a); i++ {
		hash[i+1] = hash[i]*P + uint64(a[i])
	}
	return hash
}

/*
a = [3 5 7]
hash[0] = 0
hash[1] = 3
hash[2] = 3*257 + 5
hash[3] = (3*257+5)*257 + 7
*/

/*
Мы используем тип uint64.
Когда число становится слишком большим, происходит переполнение.
Компьютер автоматически оставляет только младшие 64 бита.
Это тоже своего рода вычисление "по модулю", только модуль равен 2⁶⁴
То есть обошлись без %.

или
const MOD = 1_000_000_007
hash = (hash*P + x) % MOD
*/

// строит массив степеней основания
// pow[i] = P^i
// Например, если P = 257: pow = [1, 257, 257², 257³, ...]
// O(min(n, m))
func buildPow(n int) []uint64 {
	pow := make([]uint64, n+1)

	pow[0] = 1

	for i := 1; i <= n; i++ {
		pow[i] = pow[i-1] * P
	}
	return pow
}

/*
при n = 5
pow[0] = 1
pow[1] = 257
pow[2] = 257²
pow[3] = 257³
pow[i] = Pⁱ
*/

// возвращает хеш любого подмассива за O(1)
// полуинтервал [l, r)
func getHash(hash []uint64, pow []uint64, l, r int) uint64 {
	return hash[r] - hash[l]*pow[r-l]
}

// проверяет, существует ли общий подмассив данной длины
// Время работы: O(n+m)
func check(a, b []int, hashA, hashB []uint64, pow []uint64, length int) bool {
	seen := make(map[uint64]struct{})

	for l := 0; l+length <= len(a); l++ {
		h := getHash(hashA, pow, l, l+length)
		seen[h] = struct{}{}
	}

	for l := 0; l+length <= len(b); l++ {
		h := getHash(hashB, pow, l, l+length)
		if _, ok := seen[h]; ok {
			return true
		}
	}

	return false
}

// ищет максимальную длину
// На каждой итерации вызывает check(mid)
// Всего таких вызовов будет примерно log₂(min(n,m)). Для 100000 это около 17.
// Поэтому общая сложность 17 × (n+m)
// func binarySearch(check func(int) bool, maxLen int) int {}
// O(log(min(n, m))) проверок
func binarySearch(a, b []int, hashA, hashB, pow []uint64) int {
	left := 0
	right := min(len(a), len(b))

	for left <= right {
		mid := (left + right) / 2

		if check(a, b, hashA, hashB, pow, mid) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}

/*
func findMaxSubArray(a, b []int) int {
	dp := make([]int, len(b)+1)

	res := 0

	for i := 1; i <= len(a); i++ {
		prev := 0
		for j := 1; j <= len(b); j++ {
			tmp := dp[j]

			if a[i-1] == b[j-1] {
				dp[j] = prev + 1
				if dp[j] > res {
					res = dp[j]
				}
			} else {
				dp[j] = 0
			}
			prev = tmp
		}
	}
	return res
}
*/
/*
func findMaxSubArrayDP(a, b []int) int {
	dp := make([]int, len(b)+1)

	res := 0

	for i := 1; i <= len(a); i++ {
		prev := 0
		for j := 1; j <= len(b); j++ {
			tmp := dp[j]

			if a[i-1] == b[j-1] {
				dp[j] = prev + 1
				if dp[j] > res {
					res = dp[j]
				}
			} else {
				dp[j] = 0
			}
			prev = tmp
		}
	}
	return res
}
*/
/*
func findMaxSubArray_Memory(a, b []int) int {
	dp := make([][]int, len(a)+1)

	for i := range dp {
		dp[i] = make([]int, len(b)+1)
	}

	res := 0

	for i := 1; i <= len(a); i++ {
		for j := 1; j <= len(b); j++ {
			if a[i-1] == b[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > res {
					res = dp[i][j]
				}
			}
		}
	}
	return res
}
*/
/*
func findMaxSubArray_Cubical(arr1, arr2 []int) int { // []int
	maxLen := 0
	// index := 0
	// idx := 0

	for i := range arr1 {
		for j := range arr2 {
			count := 0
			// if arr1[i] == arr2[j] {
			// idx := i
			for i+count < len(arr1) && j+count < len(arr2) && arr1[i+count] == arr2[j+count] {

				// fmt.Println("i+count=", i+count, "; j+count=", j+count, "; arr1[i]=", arr1[i], "; arr2[j]=", arr2[j])
				// idx = i
				count++
			}
			if count > maxLen {
				// fmt.Println("count:", count, "; maxLen:", maxLen)
				maxLen = count
				// index = idx
				count = 0
				// idx = 0
			}
			// }
		}
	}
	// fmt.Println("maxLen:", maxLen)
	return maxLen // arr1[index : index+maxLen]
}
*/
/*
func readArray(scanner *bufio.Scanner, n int) []int {
	scanner.Scan()
	strList := strings.Fields(scanner.Text())

	arr := make([]int, n)

	for i := range n {
		arr[i], _ = strconv.Atoi(strList[i])
	}

	return arr
}
*/
