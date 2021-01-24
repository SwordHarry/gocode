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

func lengthOfLongestSubstring(s string) int {
	le := len(s)
	if le < 2 {
		return le
	}
	m := make(map[byte]int, le)
	var result, temp int
	var j int

	// 解法1：滑动窗口妙解
	//for i := 0; i < le; i++ {
	//	if i != 0 {
	//		delete(m, s[i-1])
	//	}
	//	for j < le && m[s[j]] == 0 {
	//		m[s[j]]++
	//		temp = j - i + 1
	//		j++
	//	}
	//	result = max(result, temp)
	//}

	// 解法2：从左往右遍历，遇到重复的则 j 指针右移
	for i := 0; i < le; i++ {
		if m[s[i]] > 0 {
			// 指针移动到重复元素的后一个，并且清除 set
			for s[j] != s[i] {
				delete(m, s[j])
				j++
			}
			j++
		} else {
			m[s[i]]++
		}
		result = max(result, i-j+1)
	}
	result = max(result, temp)

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
