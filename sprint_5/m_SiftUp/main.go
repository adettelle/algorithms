package main

import "fmt"

func siftUp(heap []int, idx int) int {
	for {
		if idx == 1 {
			return idx
		}

		parentIdx := idx / 2
		fmt.Println("parentIdx:", parentIdx)

		// если родитель больше, дальше подниматься нельзя
		if heap[parentIdx] >= heap[idx] {
			return idx
		}

		// иначе меняем местами

		heap[parentIdx], heap[idx] = heap[idx], heap[parentIdx]
		fmt.Println("heap[parentIdx]:", heap[parentIdx], "; heap[idx]:", heap[idx])
		idx = parentIdx
	}
}

func main() {
	sample := []int{-1, 12, 6, 8, 3, 15, 7}
	fmt.Println("res:", siftUp(sample, 5))
	if siftUp(sample, 5) != 1 {
		fmt.Println(sample)
		fmt.Println(siftUp(sample, 5))
	}

	sample2 := []int{-1, 12, 6, 8, 3, 4, 7}
	fmt.Println("res:", siftUp(sample2, 6))

	if siftUp(sample2, 6) != 6 {
		fmt.Println(sample2)
		fmt.Println(siftUp(sample2, 6))
	}
}

//        12(1)
//      /      \
//	   6(2)     8(3)
//	  / \       /
//	3(4) 4(5)  7(6)

// arr = [12, 6, 8, 3, 4, 7]
// siftUp(arr, 6)

//        12(1)
//      /    \
//	   6(2)   8(3)
//	  / \      /
//	3(4) 4(5) 7(6)
// arr = [12, 6, 8, 3, 4, 7]
// newPos = 6
