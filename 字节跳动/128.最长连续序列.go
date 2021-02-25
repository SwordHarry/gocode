package main

/**
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
进阶：你可以设计并实现时间复杂度为 O(n) 的解决方案吗？

示例 1：
输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。

示例 2：
输入：nums = [0,3,7,2,5,8,4,6,0,1]
输出：9

提示：
0 <= nums.length <= 104
-109 <= nums[i] <= 109
*/

/**
思路：哈希表，从小到大遍历
*/
func longestConsecutive(nums []int) int {
	length := len(nums)
	var result int
	if length == 0 {
		return result
	}

	m := make(map[int]bool, length)
	for i := 0; i < length; i++ {
		m[i] = true
	}
	for i := 0; i < length; i++ {
		if !m[nums[i]-1] {
			temp := nums[i]
			tempResult := 1
			for m[temp] {
				temp++
				tempResult++
			}
			result = max(result, tempResult)
		}
	}

	return result
}

func main() {

}
