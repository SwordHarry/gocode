package main

import (
	"fmt"
	"strconv"
)

/**
将正整数 n 表示成一系列正整数之和
求正整数的不同划分个数
*/

// 全局变量存储记录数和组合数
var total, count int
var combination []int

// 回溯法-剪枝
func partition(n, m int) {
	if n <= 0 {
		total++
		for i := 0; i < count; i++ {
			fmt.Print(strconv.Itoa(combination[i]) + " ")
		}
		fmt.Println()
	} else {
		for i := m; i <= n; i++ {
			if len(combination) <= count {
				combination = append(combination, i)
			} else {
				combination[count] = i
			}
			count++
			partition(n-i, i)
			count--
		}
	}
}

func main() {
	fmt.Println("划分 6 的组合为：")
	partition(6, 1)
	fmt.Println("划分数量为", total)
	total = 0
	fmt.Println("划分 5 的组合为：")
	partition(5, 1)
	fmt.Println("划分数量为", total)
}
