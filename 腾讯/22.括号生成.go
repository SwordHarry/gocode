package main

/**
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

示例 1：
输入：n = 3
输出：["((()))","(()())","(())()","()(())","()()()"]

示例 2：
输入：n = 1
输出：["()"]

提示：
1 <= n <= 8
*/

// 思路：回溯法
var result []string

func generateParenthesis(n int) []string {
	result = nil
	gen("", 0, 0, n)
	return result
}

func gen(s string, open, close, max int) {
	if len(s) == 2*max {
		result = append(result, s)
		return
	}
	// 左括号小于 n 则放置左括号，否则右括号小于左括号就放右括号
	// 左括号必定先于右括号
	if open < max {
		s += "("
		gen(s, open+1, close, max)
		s = s[:len(s)-1]
	}
	if close < open {
		s += ")"
		gen(s, open, close+1, max)
		s = s[:len(s)-1]
	}
}

func main() {

}
