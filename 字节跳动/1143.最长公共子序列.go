package main

import "fmt"

/**
给定两个字符串text1 和text2，返回这两个字符串的最长公共子序列的长度。
一个字符串的子序列是指这样一个新的字符串：它是由原字符串在不改变字符的相对顺序的情况下删除某些字符（也可以不删除任何字符）后组成的新字符串。
例如，"ace" 是 "abcde" 的子序列，但 "aec" 不是 "abcde" 的子序列。两个字符串的「公共子序列」是这两个字符串所共同拥有的子序列。
若这两个字符串没有公共子序列，则返回 0。
示例 1:
输入：text1 = "abcde", text2 = "ace"
输出：3
解释：最长公共子序列是 "ace"，它的长度为 3。

示例 2:
输入：text1 = "abc", text2 = "abc"
输出：3
解释：最长公共子序列是 "abc"，它的长度为 3。

示例 3:
输入：text1 = "abc", text2 = "def"
输出：0
解释：两个字符串没有公共子序列，返回 0。

提示:
1 <= text1.length <= 1000
1 <= text2.length <= 1000
输入的字符串只含有小写英文字符。
*/

/**
思路：
	动规：dp[i][j] 表示 text1[0:i-1] 和 text2[0:j-1] 的最长公共子序列
	边界条件：dp[0][0] = text1[0] == text2[0]
	递推式：if text1[i] == text2[j]: dp[i][j] = dp[i-1][j-1] + 1
	      else: dp[i][j] = max(dp[i-1][j], dp[i][j-1])
*/

// 进阶 1：空间优化：只用两行，按位与&1，进行01转换
// 进阶 2：（最终版）空间O(n)，时间O(n^2)
func longestCommonSubsequence(text1 string, text2 string) int {
	len1 := len(text1)
	len2 := len(text2)
	dp := make([]int, len2+1)
	for i := 0; i < len1; i++ {
		// 可以理解为一个 2*2 的滑动窗口
		var leftTop int
		for j := 0; j < len2; j++ {
			top := dp[j+1]
			if text1[i] == text2[j] {
				dp[j+1] = leftTop + 1
			} else {
				// 旧的 dp[i+1] 存的是上次的值，可以理解为 top
				dp[j+1] = max(dp[j], top)
			}
			leftTop = top
		}
	}
	return dp[len2]
}

// 问题进阶：找出最长公共子序列，而不是长度
func longestCommonSubsequence2(text1 string, text2 string) string {
	len1 := len(text1)
	len2 := len(text2)
	dp := make([][]int, len1+1)
	for i := 0; i < len1+1; i++ {
		dp[i] = make([]int, len2+1)
	}
	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {

			if text1[i] == text2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}

	// 利用完整的 dp 构造出最长公共子序列
	var result string
	i, j := len1-1, len2-1
	for i >= 0 && j >= 0 {
		if text1[i] == text2[j] {
			result = string(text1[i]) + result
			i--
			j--
		} else {
			if dp[i][j+1] >= dp[i+1][j] {
				i--
			} else {
				j--
			}
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	//x, y := "acdefggh", "cegdhfghb"
	x, y := "abcde", "ace"
	fmt.Println("最长公共子序列为：", longestCommonSubsequence2(x, y))
}
