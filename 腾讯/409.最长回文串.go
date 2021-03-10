package main

/**
给定一个包含大写字母和小写字母的字符串，找到通过这些字母构造成的最长的回文串。
在构造过程中，请注意区分大小写。比如 "Aa" 不能当做一个回文字符串。
注意:
假设字符串的长度不会超过 1010。

示例 1:
输入:
"abccccdd"
输出:
7
解释:
我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。
*/

func longestPalindrome(s string) int {
	length := len(s)
	var result int
	if length == 0 {
		return result
	}
	// 这里是构造最长回文串，而不是找出原本的最长回文串
	m := make(map[byte]int)
	var flag bool

	for i := 0; i < length; i++ {
		m[s[i]]++
	}
	for _, v := range m {
		if v%2 == 1 {
			if flag {
				if v > 2 {
					result += (v - 1)
				}
			} else {
				result += v
				flag = true
			}
		} else {
			result += v
		}
	}

	return result
}

func main() {

}
