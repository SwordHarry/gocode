package main

/**
给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
示例 1：
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。

示例 2：
输入：height = [4,2,0,3,2,5]
输出：9

提示：
n == height.length
0 <= n <= 3 * 104
0 <= height[i] <= 105
*/

// 最优解：左右指针夹逼原理
func trap(height []int) int {
	length := len(height)
	left, leftMax, right, rightMax := 0, 0, length-1, 0
	// 双指针
	var result int
	for left < right {
		if height[left] < height[right] {
			if height[left] > leftMax {
				leftMax = height[left]
			} else {
				result += leftMax - height[left]
			}
			left++
		} else {
			if height[right] > rightMax {
				rightMax = height[right]
			} else {
				result += rightMax - height[right]
			}
			right--
		}
	}
	return result
}

// 思路2：动规，备忘录 O(n) 4ms 3.1MB
func trap2(height []int) int {
	length := len(height)
	var result int
	if length < 3 {
		return result
	}

	left, right := make([]int, length), make([]int, length)
	left[0] = height[0]
	right[length-1] = height[length-1]
	for i := 1; i < length; i++ {
		left[i] = max(left[i-1], height[i])
		right[length-i-1] = max(right[length-i], height[length-i-1])
	}

	for i := 1; i < length-1; i++ {
		minWall := min(left[i], right[i])
		if minWall > height[i] {
			result += minWall - height[i]
		}
	}
	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {

}
