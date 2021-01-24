package main

import (
	"fmt"
	"math"
	"sort"
)

/**
给定一个包括n 个整数的数组nums和 一个目标值target。找出nums中的三个整数，使得它们的和与target最接近。返回这三个数的和。假定每组输入只存在唯一答案。

示例：
输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。

提示：
3 <= nums.length <= 10^3
-10^3<= nums[i]<= 10^3
-10^4<= target<= 10^4
*/

type array2 []int

func (n array2) Len() int {
	return len(n)
}
func (n array2) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}
func (n array2) Less(i, j int) bool {
	return n[i] < n[j]
}

func main() {
	fmt.Println(threeSum2([]int{0, 2, 1, -3}, 1))
}

func closeTargetVal(a, b, c int) int {
	a1 := math.Abs(float64(a - c))
	b1 := math.Abs(float64(b - c))
	if a1 > b1 {
		return b
	}
	return a
}

func threeSum2(nums []int, target int) int {
	// 排序保证 (a, b, c) 的三元组 a < b < c
	sort.Sort(array2(nums))
	length := len(nums)
	result := nums[0] + nums[1] + nums[2]
	if length == 3 {
		return result
	}
	// 双指针法
	for a := 0; a < length; a++ {
		// 只有和数字不同，才进入迭代
		if a == 0 || nums[a] != nums[a-1] {
			// 双指针法加速，[a+1, n-1] 做中间收缩
			b := a + 1
			c := length - 1
			for b < c {
				if nums[a]+nums[b]+nums[c] == target {
					return target
				}
				result = closeTargetVal(nums[a]+nums[b]+nums[c], result, target)
				if nums[a]+nums[b]+nums[c] < target {
					b++
				} else {
					c--
				}
			}
		}
	}
	return result
}
