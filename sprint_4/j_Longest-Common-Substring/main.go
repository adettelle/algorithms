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
