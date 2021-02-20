package main

import "math"

/**
给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
注意：如果 s 中存在这样的子串，我们保证它是唯一的答案。

示例 1：
输入：s = "ADOBECODEBANC", t = "ABC"
输出："BANC"

示例 2：
输入：s = "a", t = "a"
输出："a"

提示：
1 <= s.length, t.length <= 105
s 和 t 由英文字母组成

进阶：你能设计一个在 o(n) 时间内解决此问题的算法吗？
*/

/**
思路：滑动窗口
*/
func minWindow(s string, t string) string {
	lenS := len(s)
	lenT := len(t)
	if lenS < lenT {
		return ""
	}

	ref, m := make(map[byte]int), make(map[byte]int)
	for i := 0; i < lenT; i++ {
		ref[t[i]]++
	}

	minLen := math.MaxInt64
	var ansL, ansR int
	// 滑动窗口，若包含 t 中的内容，r左移；完全包含的情况下，则l右移，直到不完全包含
	for l, r := 0, 0; r < lenS; r++ {
		if ref[s[r]] > 0 {
			m[s[r]]++
		}
		for checkALl(ref, m) && l <= r {
			if r-l+1 < minLen {
				minLen = r - l + 1
				ansL, ansR = l, r+1
			}
			if ref[s[l]] > 0 {
				m[s[l]]--
			}
			l++
		}
	}

	if ansL == ansR {
		return ""
	}
	return s[ansL:ansR]
}

func checkALl(ref, m map[byte]int) bool {
	for k, v := range ref {
		if m[k] < v {
			return false
		}
	}
	return true
}

func main() {

}
