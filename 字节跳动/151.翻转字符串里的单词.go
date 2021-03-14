package main

import "strings"

/**
给定一个字符串，逐个翻转字符串中的每个单词。
说明：
无空格字符构成一个 单词 。
输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。
如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

示例 1：
输入："the sky is blue"
输出："blue is sky the"

示例 2：
输入：" hello world! "
输出："world! hello"
解释：输入字符串可以在前面或者后面包含多余的空格，但是反转后的字符不能包括。

示例 3：
输入："a good  example"
输出："example good a"
解释：如果两个单词间有多余的空格，将反转后单词间的空格减少到只含一个。

示例 4：
输入：s = "  Bob    Loves  Alice   "
输出："Alice Loves Bob"

示例 5：
输入：s = "Alice does not even like bob"
输出："bob like even not does Alice"

提示：
1 <= s.length <= 104
s 包含英文大小写字母、数字和空格 ' '
s 中 至少存在一个 单词

进阶：
请尝试使用 O(1) 额外空间复杂度的原地解法。
*/

/**
思路：1.使用 strings 包的 api TrimSpace 和 Split
	2. 若不允许使用 api
*/

// 用 api
func reverseWords(s string) string {
	s = strings.TrimSpace(s)
	strs := strings.Split(s, " ")
	length := len(strs)
	var result string
	for i := length - 1; i >= 0; i-- {
		if strs[i] != "" {
			if i == 0 {
				result += strs[i]
			} else {
				result += strs[i] + " "
			}
		}

	}
	return result
}

// 不用 api：先去除左右空格，然后将字母按空格切分
func reverseWords2(s string) string {
	length := len(s)
	var result string
	// 去除左右空格
	i := 0
	for ; i < length && s[i] == ' '; i++ {
	}
	s = s[i:]
	length = len(s)
	i = length - 1
	for ; i >= 0 && s[i] == ' '; i-- {
	}
	s = s[:i+1]

	// 将字母按空格切分
	length = len(s)
	i = 0
	for ; i < length; i++ {
		if s[i] != ' ' {
			j := i
			for ; j < length && s[j] != ' '; j++ {
			}
			result = s[i:j] + " " + result
			i = j
		}
	}
	// 最后还有一个空格
	return result[:len(result)-1]
}

func main() {

}
