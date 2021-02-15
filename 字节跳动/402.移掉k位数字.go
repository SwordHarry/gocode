package main

/**
给定一个以字符串表示的非负整数num，移除这个数中的 k 位数字，使得剩下的数字最小。
注意:
num 的长度小于 10002 且 ≥ k。
num 不会包含任何前导零。

示例 1 :
输入: num = "1432219", k = 3
输出: "1219"
解释: 移除掉三个数字 4, 3, 和 2 形成一个新的最小的数字 1219。

示例 2 :
输入: num = "10200", k = 1
输出: "200"
解释: 移掉首位的 1 剩下的数字为 200. 注意输出不能有任何前导零。

示例 3 :
输入: num = "10", k = 2
输出: "0"
解释: 从原数字移除所有的数字，剩余为空就是0。
*/

// 思路：贪心 + 单调栈
func removeKdigits(num string, k int) string {
	length := len(num)
	if length == k || length == 0 {
		return "0"
	}

	var s []uint8
	s = append(s, num[0])
	for i := 1; i < length; i++ {
		top := s[len(s)-1]
		for k > 0 && top > num[i] {
			s = s[:len(s)-1]
			k--
			if len(s) == 0 {
				break
			} else {
				top = s[len(s)-1]
			}
		}
		s = append(s, num[i])
	}

	// 注意 0200 的情况
	for len(s) > 0 && s[0] == '0' {
		s = s[1:]
		if len(s) == 0 {
			return "0"
		}
	}

	// 这里直接 return string() 比拼接快得多
	return string(s[0 : len(s)-k])
}

func main() {

}
