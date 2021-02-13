package main

import "fmt"

func mergeSort(s, temp []int, i, j int) {
	if i < j {
		mid := (i + j) / 2
		mergeSort(s, temp, i, mid)
		mergeSort(s, temp, mid+1, j)
		merge(s, temp, i, mid, j)
	}
}

func merge(s, temp []int, left, mid, right int) {
	i, j := left, mid+1
	var index int
	for i <= mid && j <= right {
		if s[i] < s[j] {
			temp[index] = s[i]
			i++
		} else {
			temp[index] = s[j]
			j++
		}
		index++
	}

	for i <= mid {
		temp[index] = s[i]
		i++
		index++
	}

	for j <= right {
		temp[index] = s[j]
		j++
		index++
	}

	index = 0
	for left <= right {
		s[left] = temp[index]
		left++
		index++
	}
}

func main() {
	s := []int{3, 5, 7, 1, 2, 9, 8, 4}
	temp := make([]int, len(s))
	mergeSort(s, temp, 0, len(s)-1)
	fmt.Println(s)
}
