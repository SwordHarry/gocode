package main

import "fmt"

/*
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"
*/

// 动态规划
func longestPalindrome(s string) string {
	length := len(s)
	if length == 0 {
		return s
	}
	var l [1000][1000]bool
	for i := 0; i < length; i++ {
		l[i][i] = true
	}

	var begin, end, maxLength int
	for i := length - 2; i >= 0; i-- {
		for j := i + 1; j < length; j++ {
			if s[i] == s[j] && j == i+1 {
				l[i][j] = true
				if maxLength < 1 {
					maxLength = 1
					begin = i
					end = j
				}
			} else if s[i] == s[j] {
				l[i][j] = l[i+1][j-1]
				if maxLength < j-i+1 && l[i][j] {
					maxLength = j - i + 1
					begin = i
					end = j
				}
			}
		}
	}
	return s[begin : end+1]
}

// 更好解：中心扩散法
func longestPalindrome2(s string) string {
	if s == "" {
		return ""
	}
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		left1, right1 := expandAroundCenter(s, i, i)
		left2, right2 := expandAroundCenter(s, i, i+1)
		if right1-left1 > end-start {
			start, end = left1, right1
		}
		if right2-left2 > end-start {
			start, end = left2, right2
		}
	}
	return s[start : end+1]
}

func expandAroundCenter(s string, left, right int) (int, int) {
	for ; left >= 0 && right < len(s) && s[left] == s[right]; left, right = left-1, right+1 {
	}
	return left + 1, right - 1
}

func main() {
	fmt.Print(longestPalindrome("abcdca"))
}
