package main

import (
	"fmt"
	"sort"
)

/**
三数之和
给你一个包含 n 个整数的数组，判断数组中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]

*/

type array []int

func (n array) Len() int {
	return len(n)
}
func (n array) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
func (n array) Less(i, j int) bool {
	return n[i] < n[j]
}

func main() {
	fmt.Println(threeSum([]int{1, 2, -2, -1}))
}

func threeSum(nums []int) [][]int {
	length := len(nums)
	if length == 0 {
		return nil
	}
	// 排序保证 (a, b, c) 的三元组 a < b < c
	sort.Sort(array(nums))
	var result [][]int
	// 双指针法
	for a := 0; a < length; a++ {
		c := length - 1
		// 只有和数字不同，才进入迭代
		if a == 0 || nums[a] != nums[a-1] {
			for b := a + 1; b < length-1; b++ {
				if b == a+1 || nums[b] != nums[b-1] {
					for c > b && nums[a]+nums[b]+nums[c] > 0 {
						c--
					}
					// 指针重合则退出
					if c == b {
						break
					}
					// 判断是否满足条件
					if nums[a]+nums[b]+nums[c] == 0 {
						result = append(result, []int{nums[a], nums[b], nums[c]})
					}
				}
			}
		}
	}
	return result
}
