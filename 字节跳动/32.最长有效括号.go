package main

/**
给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。

示例 1：
输入：s = "(()"
输出：2
解释：最长有效括号子串是 "()"

示例 2：
输入：s = ")()())"
输出：4
解释：最长有效括号子串是 "()()"

示例 3：
输入：s = ""
输出：0

提示：
0 <= s.length <= 3 * 104
s[i] 为 '(' 或 ')'
*/

/**
思路1：动态规划
	dp[i] 表示字符串 0 到 i 的最长有效括号
	分类讨论：
		1. 考虑 ".....()" 的形式，则 dp[i] = dp[i-2] + 2
		2. 考虑 "..((....))" 的形式，则 dp[i] = dp[i-1] + dp[i-dp[i-1]-2]+2
			内层的括号为 dp[i-1]，而 dp[i-1] 往左两位为左边的有效括号，需要加上，再加上新增的合法2个括号
			需要注意的是，防止越界 i-dp[i-1]-1 >= 0
	故递推式：
		dp[i] = dp[i-2] + 2 (s[i] == ')' && s[i-1] == '(')
		dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2 (s[i] == ')' && s[i-1] == ')' && s[i-dp[i-1]-1] == '(')
	边界条件：
		dp[0] = 0
*/

func longestValidParentheses(s string) int {
	length := len(s)
	if length < 2 {
		return 0
	}
	var result int
	dp := make([]int, length)
	for i := 1; i < length; i++ {
		var temp int
		if s[i] == ')' {
			if s[i-1] == '(' {
				// dp[i] = dp[i-2] + 2
				if i-2 >= 0 {
					temp = dp[i-2]
				}
				dp[i] = temp + 2
			} else if s[i-1] == ')' {
				// dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
				left := i - dp[i-1] - 1
				if left >= 0 && s[left] == '(' {
					if left >= 1 {
						temp = dp[left-1]
					}
					dp[i] = dp[i-1] + temp + 2
				}
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

func main() {

}
