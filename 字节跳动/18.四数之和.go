package main

import "sort"

/**
给定一个包含 n 个整数的数组 nums 和一个目标值 target，
判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。
注意：答案中不可以包含重复的四元组。

示例 1：
输入：nums = [1,0,-1,0,-2,2], target = 0
输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]

示例 2：
输入：nums = [], target = 0
输出：[]

提示：
0 <= nums.length <= 200
-109 <= nums[i] <= 109
-109 <= target <= 109
*/

/**
思路：排序 + 三数之和 O(n^3)
*/
func fourSum(nums []int, target int) [][]int {
	length := len(nums)
	var result [][]int
	if length < 4 {
		return result
	}

	// 排序
	sort.Sort(sortNums(nums))

	// 外层两层循环，内层双指针，需要去重三次
	for i := 0; i < length-3; {
		for j := i + 1; j < length-2; {
			l, r := j+1, length-1
			for l < r {
				sum := nums[i] + nums[j] + nums[l] + nums[r]
				if sum == target {
					result = append(result, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					// 去重
					for l < r && nums[l] == nums[l-1] {
						l++
					}
				} else if sum < target {
					l++
				} else {
					r--
				}
			}
			j++
			// 去重
			for j < length-2 && nums[j] == nums[j-1] {
				j++
			}
		}
		i++
		// 去重
		for i < length-3 && nums[i] == nums[i-1] {
			i++
		}
	}
	return result
}

type sortNums []int

func (s sortNums) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortNums) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortNums) Len() int {
	return len(s)
}

func main() {

}
