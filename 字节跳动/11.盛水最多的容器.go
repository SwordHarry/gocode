package main

/**
给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点(i,ai) 。在坐标内画 n 条垂直线，垂直线 i的两个端点分别为(i,ai) 和 (i, 0) 。找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。
说明：你不能倾斜容器。
示例 1：
输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为49。

示例 2：
输入：height = [1,1]
输出：1

示例 3：
输入：height = [4,3,2,1,4]
输出：16

示例 4：
输入：height = [1,2,1]
输出：2

提示：
n = height.length
2 <= n <= 3 * 104
0 <= height[i] <= 3 * 104
*/

// 思路：双指针法
func maxArea(height []int) int {
	length := len(height)
	if length == 2 {
		return min(height[0], height[1])
	}
	i, j := 0, length-1
	var result int
	for i < j {
		temp := min(height[i], height[j]) * (j - i)
		if result < temp {
			result = temp
		}

		if height[i] > height[j] {
			j--
		} else if height[i] < height[j] {
			i++
		} else {
			if i+1 < j {
				if height[i+1] < height[j-1] {
					j--
				} else {
					i++
				}
			} else {
				break
			}
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
