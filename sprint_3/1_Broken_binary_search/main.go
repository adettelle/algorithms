/*
-- ПРИНЦИП РАБОТЫ --
По заданию на вход получаем сдвинутый циклический строго возрастающий массив.
То есть массив состоит из двух возрастающих частей.
Сначала мы ищем индекс максимального элемента, таким образом получаем
две отсортированные части исходного массива.
Далее производим бинарный поиск элемента в этих частях и возвращаем его индекс.

-- ДОКАЗАТЕЛЬСТВО КОРРЕКТНОСТИ --
После нахождения индекса максимума массив разбивается на две части.
Обе части строго возрастают.
Следовательно для каждой из них корректно работает обычный бинарный поиск binarySearch.
Если искомый элемент принадлежит массиву, то он находится либо в arrLeft, либо в arrRight.
Функция выполняет бинарный поиск сначала в левой части, затем в правой.
При нахождении элемента в правой части локальный индекс корректно переводится в глобальный.

-- ВРЕМЕННАЯ СЛОЖНОСТЬ --
Временная сложность бинарного поиска равна O(logn), потому что на каждом шаге алгоритм
делит область поиска пополам.

-- ПРОСТРАНСТВЕННАЯ СЛОЖНОСТЬ --
Память, затрачиваемая на выполнение бинарного поиска, составляет O(1),
так как он не требует дополнительной памяти, кроме нескольких переменных
для хранения индексов.
*/
package main

import (
	"fmt"
)

func main() {
	/*
		reader := bufio.NewReader(os.Stdin)

		var n int
		fmt.Fscan(reader, &n)

		var k int
		fmt.Fscan(reader, &k)

		arr := make([]int, n)
		for i := range n {
			fmt.Fscan(reader, &arr[i])
		}

		fmt.Println("n =", n, "; k =", k)
		fmt.Println(arr)

		fmt.Println("result = ", brokenSearch(arr, k))
	*/

	a := []int{8, 10, 0, 2, 4}
	fmt.Println(" array = {8, 10, 0, 2, 4};", "max idx =", binarySearchMax(a))

	arr4 := []int{99, 100, 0, 1, 4, 5, 7, 12, 15, 19}
	// fmt.Println(binarySearchMax(arr4))
	fmt.Println(partArray(arr4, binarySearchMax(arr4)))
	fmt.Println(brokenSearch(arr4, 12)) // 7

	// fmt.Println(brokenSearch(arr4, 100)) // 2
	// fmt.Println(brokenSearch(arr4, 1))   // 3

}

// Binary Search of max element of array which consist of two sorted parts
// returns index of max element
func binarySearchMax(arr []int) int {
	fmt.Println("binarySearchMax: ", arr, "len = ", len(arr))

	if len(arr) == 1 {
		return 0
	}

	left := 0
	right := len(arr) - 1

	if arr[right] >= arr[left] {
		fmt.Println(" 111 idx =", right)
		return right
	}

	if arr[left] > arr[left+1] {
		fmt.Println(" 222 idx =", left)
		return left
	}

	for left <= right {
		if left == right {
			return left
		}
		middle := left + (right-left)/2

		if left == 0 && right == 1 {
			if arr[right] > arr[left] {
				return right
			} else {
				return left
			}
		}

		// fmt.Println("mid-1 =", middle-1, "; mid =", middle, "; mid+1 =", middle+1)
		// fmt.Println("mid-1 =", arr[middle-1], "; mid =", arr[middle], "; mid+1 =", arr[middle+1])
		// fmt.Println("left =", left, "right =", right)

		if arr[middle] >= arr[middle-1] && arr[middle] > arr[middle+1] {
			// fmt.Println(" 333 idx =", middle)
			return middle
		} else if arr[middle] >= arr[left] {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}
	return -1
}

func partArray(arr []int, idx int) ([]int, []int) {
	fmt.Println("index =", idx)
	return arr[:idx+1], arr[idx+1:]
}

func brokenSearch(arr []int, k int) int {
	// fmt.Println(arr, "; k =", k)
	idx := 0

	arrLeft, arrRight := partArray(arr, binarySearchMax(arr))
	fmt.Println(arrLeft, arrRight)

	if len(arrLeft) != 0 {
		idx = binarySearch(arrLeft, k)
		if idx != -1 {
			return idx
		}
	}
	if len(arrRight) != 0 {
		// fmt.Println("len(arrRight) =", len(arrRight))
		// fmt.Println("len(arrLeft) =", len(arrLeft))
		idx = binarySearch(arrRight, k)
		// fmt.Println("idx =", idx)
		if idx != -1 {
			idx = len(arrLeft) + idx
			return idx
		}
	}
	return idx
}

// Binary Search
func binarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		middle := left + (right-left)/2

		if target == arr[middle] {
			return middle
		}
		if target < arr[middle] {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return -1
}
