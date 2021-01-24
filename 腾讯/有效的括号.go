package main

import "fmt"

/**
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例2:

输入: "()[]{}"
输出: true
示例3:

输入: "(]"
输出: false
示例4:

输入: "([)]"
输出: false
示例5:

输入: "{[]}"
输出: true
*/

func main() {
	fmt.Println(isValid("([)]"))
}

func isValid(s string) bool {
	length := len(s)
	if length == 0 {
		return true
	}
	m := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	stack := make([]byte, 0, length)
	for _, b := range s {
		if b == '(' || b == '[' || b == '{' {
			stack = append(stack, byte(b))
			continue
		}
		curLen := len(stack)
		if curLen != 0 {
			if byte(b) == m[stack[curLen-1]] {
				stack = stack[:curLen-1]
			} else {
				return false
			}
		} else {
			return false
		}
	}
	return len(stack) == 0
}
