package main

import "fmt"

/**
题目：
给定二维数组，输出按 aii 翻转的结果
示例：
3 3
1 2 3
4 5 6
7 8 9

结果：
1 4 7
2 5 8
3 6 9
*/

func getFinalNums(nums [][]int) [][]int {
	row := len(nums)
	column := len(nums[0])

	result := make([][]int, column)
	for i := 0; i < column; i++ {
		result[i] = make([]int, row)
	}

	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			result[j][i] = nums[i][j]
		}
	}
	return result
}

func main() {
	fmt.Println(getFinalNums([][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}))
}
