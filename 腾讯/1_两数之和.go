package main

import "fmt"
/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
*/
func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	fmt.Print(result)
}

// 两编 hash 表
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	var result []int
	for i,v := range nums {
		m[v] = i
		if x,ok := m[target - v];ok && i != x {
			result = append(result,i,x)
			break
		}
	}
	return result
}

// 一遍 hash 表
func twoSum_2(nums []int, target int)  {
	m := make(map[int]int)
	var result []int
	for i,v := range nums {
		if x,ok := m[target - v];ok && i != x {
			result = append(result,i,x)
			break
		}
		m[v] = i
	}
	return result
}
