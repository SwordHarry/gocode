package main

import "fmt"

/**
流程：先将数组构建为一个大顶堆
	然后将堆顶元素放到数组末尾，并将剩余元素继续构建大顶堆
	直到剩余元素为0
*/

func heapSort(s []int) []int {
	length := len(s)
	// 1. 构建大顶堆
	buildMaxHeap(s, length)
	// 2. 将堆顶放到最后，再将剩余元素继续构建大顶堆
	for i := length - 1; i >= 0; i-- {
		s[0], s[i] = s[i], s[0]
		length -= 1
		heapify(s, 0, length)
	}
	return s
}

func buildMaxHeap(s []int, length int) {
	for i := length / 2; i >= 0; i-- {
		heapify(s, i, length)
	}
}

// 堆化子堆
func heapify(s []int, index, length int) {
	left := index*2 + 1
	right := index*2 + 2
	largest := index
	if left < length && s[left] > s[largest] {
		largest = left
	}
	if right < length && s[right] > s[largest] {
		largest = right
	}
	if largest != index {
		s[index], s[largest] = s[largest], s[index]
		// 子堆继续构建大顶堆
		heapify(s, largest, length)
	}
}

func main() {
	s := []int{4, 3, 5, 8, 2, 7, 1, 9}
	s = heapSort(s)
	fmt.Println(s)
}
