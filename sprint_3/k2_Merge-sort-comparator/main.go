package main

import "fmt"

func main() {
	nums := []int{38, 27, 43, 3, 9, 82, 10}
	sorted := mergeSortWithComparator(nums)
	fmt.Println(sorted)
}

func less(a, b int) bool {
	return a < b
}

func mergeSortWithComparator(arr []int) []int {
	if len(arr) <= 1 { // базовый случай рекурсии
		return arr
	}

	mid := len(arr) / 2
	// запускаем сортировку рекурсивно на левой половине
	left := mergeSortWithComparator(arr[:mid])
	// запускаем сортировку рекурсивно на правой половине
	right := mergeSortWithComparator(arr[mid:])

	return merge(left, right)
}

// сливаем результаты
func merge(left, right []int) []int {
	// заводим массив для результата сортировки
	res := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		// выбираем, из какого массива забрать минимальный элемент
		if less(left[i], right[j]) {
			res = append(res, left[i])
			i++
		} else {
			res = append(res, right[j])
			j++
		}
	}
	//  Добавляем остатки, если в одном из списков еще есть элементы
	res = append(res, left[i:]...)
	res = append(res, right[j:]...)
	return res
}
