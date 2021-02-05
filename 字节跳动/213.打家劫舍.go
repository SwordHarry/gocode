package main

/**
你是一个专业的小偷，计划偷窃沿街的房屋，每间房内都藏有一定的现金。这个地方所有的房屋都 围成一圈 ，这意味着第一个房屋和最后一个房屋是紧挨着的。
同时，相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警 。
给定一个代表每个房屋存放金额的非负整数数组，计算你 在不触动警报装置的情况下 ，能够偷窃到的最高金额。
示例1：
输入：nums = [2,3,2]
输出：3
解释：你不能先偷窃 1 号房屋（金额 = 2），然后偷窃 3 号房屋（金额 = 2）, 因为他们是相邻的。

示例 2：
输入：nums = [1,2,3,1]
输出：4
解释：你可以先偷窃 1 号房屋（金额 = 1），然后偷窃 3 号房屋（金额 = 3）。
偷窃到的最高金额 = 1 + 3 = 4 。

示例 3：
输入：nums = [0]
输出：0

提示：
1 <= nums.length <= 100
0 <= nums[i] <= 1000
*/

func rob(nums []int) int {
	// 先求 0 -> n-1 的最大值
	// 再求 1 -> n 的最大值
	// 最后求两者最大值

	length := len(nums)
	if length == 0 {
		return 0
	}
	if length == 1 {
		return nums[0]
	}
	if length == 2 {
		return max(nums[0], nums[1])
	}

	nums1 := nums[:length-1]
	nums2 := nums[1:]
	s1 := make([]int, length-1)
	s2 := make([]int, length-1)
	s1[0] = nums1[0]
	s1[1] = max(nums1[0], nums1[1])
	s2[0] = nums2[0]
	s2[1] = max(nums2[0], nums2[1])
	for i := 2; i < length-1; i++ {
		s1[i] = max(s1[i-2]+nums1[i], s1[i-1])
		s2[i] = max(s2[i-2]+nums2[i], s2[i-1])
	}

	return max(s1[length-2], s2[length-2])
}

func main() {

}
