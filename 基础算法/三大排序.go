package main

import "fmt"

// 快排
func QuickSort(nums []int, i, j int) {
	if i < j {
		mid := partition(nums, i, j)
		QuickSort(nums, i, mid)
		QuickSort(nums, mid+1, j)
	}
}

func partition(nums []int, i, j int) int {
	p := nums[i]
	for i < j {
		for i < j && nums[j] >= p {
			j--
		}
		nums[i] = nums[j]
		for i < j && nums[i] <= p {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = p
	return i
}

// 堆排序
func HeapSort(nums []int) {
	le := len(nums)
	buildBigHeap(nums, le)
	for i := le - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		le--
		heapify(nums, 0, le)
	}
}

func buildBigHeap(nums []int, le int) {
	for i := le / 2; i >= 0; i-- {
		heapify(nums, i, le)
	}
}

func heapify(nums []int, i, le int) {
	largest, left, right := i, 2*i+1, 2*i+2
	if left < le && nums[left] > nums[largest] {
		largest = left
	}
	if right < le && nums[right] > nums[largest] {
		largest = right
	}
	if largest != i {
		nums[i], nums[largest] = nums[largest], nums[i]
		heapify(nums, largest, le)
	}
}

// 归并排序
func MergeSort(nums, temp []int, i, j int) {
	if i < j {
		mid := (i + j) / 2
		MergeSort(nums, temp, i, mid)
		MergeSort(nums, temp, mid+1, j)
		merge(nums, temp, i, j, mid)
	}
}

func merge(nums, temp []int, i, j, m int) {
	l, r, index := i, m+1, 0
	for l <= m && r <= j {
		if nums[l] < nums[r] {
			temp[index] = nums[l]
			l++
		} else {
			temp[index] = nums[r]
			r++
		}
		index++
	}

	for l <= m {
		temp[index] = nums[l]
		l++
		index++
	}
	for r <= j {
		temp[index] = nums[r]
		r++
		index++
	}

	for k := 0; k < index; k++ {
		nums[i+k] = temp[k]
	}
}

func main() {
	s := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	t := make([]int, len(s))
	MergeSort(s, t, 0, len(s)-1)
	fmt.Println(s)
}
