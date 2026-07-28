package main

import "fmt"

func siftDown(heap []int, idx int) int {
	for {
		left := 2 * idx
		right := 2*idx + 1

		// Проверяем, существует ли левый потомок:
		if left >= len(heap) {
			return idx // Если элемент не двигался, нужно вернуть его индекс
		}

		// right < len(heap) проверяет, что есть оба дочерних узла
		idxLargest := left

		if right < len(heap) && heap[right] > heap[left] {
			idxLargest = right
		}

		if heap[idxLargest] <= heap[idx] {
			return idx
		}

		heap[idx], heap[idxLargest] = heap[idxLargest], heap[idx]
		idx = idxLargest
	}
}

/*
// вариант с рекурсией:

	func siftDown(heap []int, idx int) int {
		left := 2 * idx
		right := 2*idx + 1

		if left >= len(heap) {
			return idx
		}

		largest := left

		if right < len(heap) && heap[right] > heap[left] {
			largest = right
		}

		if heap[largest] <= heap[idx] {
			return idx
		}

		heap[idx], heap[largest] = heap[largest], heap[idx]

		return siftDown(heap, largest)
	}
*/
func main() {
	sample := []int{-1, 12, 1, 8, 3, 4, 7}
	if siftDown(sample, 2) != 5 {
		fmt.Println(sample)
		fmt.Println(siftDown(sample, 2))
	}
}

//        12(1)
//      /      \
//	   1(2)     8(3)
//	  / \       /
//	3(4) 4(5)  7(6)

// arr = [12, 1, 8, 3, 4, 7]
// siftDown(arr, 2)

//        12
//      /    \
//	   4      8
//	  / \    /
//	 3   1  7
// arr = [12, 4, 8, 3, 1, 7]
// newPos = 5
