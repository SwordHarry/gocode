package main

import "fmt"

/**
给定一个字符串，请你找出其中不含有重复字符的最长子串的长度。
示例1:
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。

示例 2:
输入: s = "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。

示例 3:
输入: s = "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。
    请注意，你的答案必须是 子串 的长度，"pwke"是一个子序列，不是子串。

示例 4:
输入: s = ""
输出: 0

提示：
0 <= s.length <= 5 * 104
s由英文字母、数字、符号和空格组成
*/

/**
思路：双指针滑动窗口
	i 左指针，j 右指针
*/
func lengthOfLongestSubstring(s string) int {
	le := len(s)
	if le < 2 {
		return le
	}
	m := make(map[byte]bool, le)
	var result int
	var j int

	// 解法1：滑动窗口妙解
	for i := 0; i < le && j < le; i++ {
		if i != 0 {
			// 注意是 i-1 不是 i!!
			delete(m, s[i-1])
		}
		for j < le && !m[s[j]] {
			m[s[j]] = true
			j++
		}
		result = max(result, j-i)
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
	fmt.Println("结果为：", lengthOfLongestSubstring("abcabcbb"))
}
