package main

/**
给你一个整数数组 nums ，找到其中最长严格递增子序列的长度。
子序列是由数组派生而来的序列，删除（或不删除）数组中的元素而不改变其余元素的顺序。例如，[3,6,2,7] 是数组 [0,3,1,6,2,2,7] 的子序列。

示例 1：
输入：nums = [10,9,2,5,3,7,101,18]
输出：4
解释：最长递增子序列是 [2,3,7,101]，因此长度为 4 。

示例 2：
输入：nums = [0,1,0,3,2,3]
输出：4

示例 3：
输入：nums = [7,7,7,7,7,7,7]
输出：1

提示：
1 <= nums.length <= 2500
-104 <= nums[i] <= 104

进阶：
你可以设计时间复杂度为 O(n2) 的解决方案吗？
你能将算法的时间复杂度降低到O(n log(n)) 吗?
*/

/**
思路：动规，dp[i] 表示 0-i 的最长递增子序列长
	边界条件： dp[0] = 1
	递推式：	dp[i] = max(dp[i], dp[0:i-1]+1) (if nums[i] > nums[0:i-1])
*/
func lengthOfLIS(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	dp := make([]int, length)
	for i := 0; i < length; i++ {
		dp[i] = 1
	}
	result := dp[0]
	for i := 1; i < length; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		result = max(result, dp[i])
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/**
最优解：贪心+二分
	维护一个数组 d[i] ，表示长度为 i 的最长上升子序列的末尾元素的最小值，
	用 len 记录目前最长上升子序列的长度，起始时 len 为 1，d[1]=nums[0]。
*/
func lengthOfLIS(nums []int) int {
	length := len(nums)
	var d []int
	d = append(d, nums[0])

	for i := 1; i < length; i++ {
		l, r := 0, len(d)-1
		if d[r] < nums[i] {
			d = append(d, nums[i])
		} else {
			// 这里用二分查找，找出第一个 >= 当前数的下标
			for l < r {
				mid := l + (r-l)/2
				// 严格递增，不等于
				if d[mid] < nums[i] {
					l = mid + 1
				} else {
					r = mid
				}
			}
			d[l] = nums[i]
		}
	}
	return len(d)
}

func main() {

}
