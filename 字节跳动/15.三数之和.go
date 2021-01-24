package main

import (
	"fmt"
	"sort"
)

/**
给你一个包含 n 个整数的数组nums，判断nums中是否存在三个元素 a，b，c ，使得a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
注意：答案中不可以包含重复的三元组。
示例 1：
输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]

示例 2：
输入：nums = []
输出：[]

示例 3：
输入：nums = [0]
输出：[]

提示：
0 <= nums.length <= 3000
-105 <= nums[i] <= 105
*/

type array struct {
	Nums []int
}

func (a *array) Len() int {
	return len(a.Nums)
}
func (a *array) Less(i int, j int) bool {
	return a.Nums[i] < a.Nums[j]
}
func (a *array) Swap(i int, j int) {
	a.Nums[i], a.Nums[j] = a.Nums[j], a.Nums[i]
}

func threeSum(nums []int) [][]int {
	// 排序数组
	a := &array{Nums: nums}
	sort.Sort(a)

	le := len(nums)
	var result [][]int
	// 三指针法，一个指针从左到右扫描，另两个指针做范围收敛
	for i := 0; i < le; i++ {
		if i == 0 || nums[i] != nums[i-1] {
			j := i + 1
			k := le - 1
			for j < k {
				addResult := nums[i] + nums[j] + nums[k]
				if addResult == 0 {
					temp := []int{nums[i], nums[j], nums[k]}
					result = append(result, temp)
					// 需要把重复的去除
					k--
					for k > 0 && nums[k] == nums[k+1] {
						k--
					}
				} else if addResult > 0 {
					k--
				} else {
					j++
				}
			}
		}
	}

	return result
}

func main() {
	fmt.Println(threeSum([]int{-2, 0, 1, 1, 2}))
}
