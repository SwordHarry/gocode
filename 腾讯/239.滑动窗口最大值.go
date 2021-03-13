package main

import "fmt"

/**
给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。
你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。
返回滑动窗口中的最大值。

示例 1：
输入：nums = [1,3,-1,-3,5,3,6,7], k = 3
输出：[3,3,5,5,6,7]
解释：
滑动窗口的位置                最大值
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7

示例 2：
输入：nums = [1], k = 1
输出：[1]
示例 3：
输入：nums = [1,-1], k = 1
输出：[1,-1]
示例 4：
输入：nums = [9,11], k = 2
输出：[11]
示例 5：
输入：nums = [4,-2], k = 2
输出：[4]

提示：
1 <= nums.length <= 105
-104 <= nums[i] <= 104
1 <= k <= nums.length
*/

/**
思路：双向单调队列
	双向队列存元素下标，保证其元素值为单调递减，若原本的元素值不为单调递减，则出队直到队列为空或单调递减
	滑动窗口过程中，同时保证队首的下标在窗口中，有点最小栈的思想，但是是动态维护
*/
func maxSlidingWindow(nums []int, k int) []int {
	length := len(nums)
	if length < 2 {
		return []int{nums[0]}
	}
	var (
		result []int
		dq     []int // 单调双向队列
		push   func(index int)
	)
	// 封装 push 和 pop 操作
	push = func(index int) {
		dqLen := len(dq)
		for dqLen > 0 && nums[dq[dqLen-1]] < nums[index] {
			dq = dq[:dqLen-1]
			dqLen = len(dq)
		}
		dq = append(dq, index)
	}

	// 滑动窗口
	i := 0
	for ; i < k; i++ {
		push(i)
	}
	result = append(result, nums[dq[0]])
	for ; i < length; i++ {
		if dq[0] <= i-k {
			dq = dq[1:]
		}
		push(i)
		result = append(result, nums[dq[0]])
	}
	return result
}

func main() {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}
