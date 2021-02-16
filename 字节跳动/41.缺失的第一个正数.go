package main

/**
给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
进阶：你可以实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案吗？

示例 1：
输入：nums = [1,2,0]
输出：3

示例 2：
输入：nums = [3,4,-1,1]
输出：2

示例 3：
输入：nums = [7,8,9,11,12]
输出：1

提示：
0 <= nums.length <= 300
-231 <= nums[i] <= 231 - 1
*/

// 若不考虑进阶，则直接使用 map 做遍历，时空间复杂度均是 O(n)
func firstMissingPositive(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 1
	}

	m := make(map[int]bool, length)

	for i := 0; i < length; i++ {
		m[nums[i]] = true
	}

	i := 1
	for m[i] {
		i++
	}
	return i
}

/**
进阶版
思想1：遍历一遍数组，将大于0和小于数组长的元素放到对应位置;
	再遍历一遍数组，如果 值和[下标+1] 不对应，则答案为[下标+1]，否则为数组长度+1

思想2：将 原地数组 表示为 哈希表，先将所有负数变为数组长度+1，再将遍历到的元素作对应位置的负数操作，即 nums[nums[i]] = -nums[nums[i]]
	第一个不是负数的下标即为答案
*/
func firstMissingPositive2(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 1
	}

	for i := 0; i < length; i++ {
		if nums[i] <= 0 {
			nums[i] = length + 1
		}
	}

	// 这里利用 数组 做 标记
	for i := 0; i < length; i++ {
		num := abs(nums[i])
		if num > 0 && num < length+1 {
			nums[num-1] = -abs(nums[num-1])
		}
	}

	for i := 0; i < length; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return length + 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {

}
