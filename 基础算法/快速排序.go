package main

import "fmt"

func quickSort(s []int, i, j int) {
	if i < j {
		index := partition(s, i, j)
		quickSort(s, i, index)
		quickSort(s, index+1, j)
	}
}

func partition(s []int, i, j int) int {
	p := s[i]
	for i < j {
		for i < j && s[j] > p {
			j--
		}
		if i < j {
			s[i] = s[j]
		}
		for i < j && s[i] < p {
			i++
		}
		if i < j {
			s[j] = s[i]
		}
	}

	s[i] = p
	return i
}

func main() {
	s := []int{3, 5, 7, 1, 2, 9, 8, 4}
	quickSort(s, 0, len(s)-1)
	fmt.Println(s)
}
