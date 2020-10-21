package main

import "fmt"

// 递归
func fibonacciSequence(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	return fibonacciSequence(n-1) + fibonacciSequence(n-2)
}

// 动归
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func fibonacciSequence2(n int) int {
	if n == 1 || n == 2 {
		return 1
	}

	l := []int{0, 1, 1}
	for i := 3; i <= n; i++ {
		l = append(l, l[i-1]+l[i-2])
	}
	return l[n]
}

func main() {
	fmt.Println(fibonacciSequence2(6))
}
