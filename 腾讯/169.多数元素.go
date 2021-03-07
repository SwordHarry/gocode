package main

/**
给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
你可以假设数组是非空的，并且给定的数组总是存在多数元素。

示例 1：
输入：[3,2,3]
输出：3

示例 2：
输入：[2,2,1,1,1,2,2]
输出：2

进阶：
尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。
*/

/**
思路：投票算法，出现次数大于一半，故该数的 count 为所有数的 count 总和还大
*/
func majorityElement(nums []int) int {
	length := len(nums)
	if length < 2 {
		return nums[0]
	}

	index, count := 0, 1
	for i := 1; i < length; i++ {
		if nums[i] != nums[index] {
			count--
			if count == 0 {
				index = i
				count = 1
			}
		} else {
			count++
		}
	}
	return nums[index]
}

func main() {

}
