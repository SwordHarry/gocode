package main

import "strconv"

/**
给定两个字符串形式的非负整数num1 和num2，计算它们的和。
提示：
num1 和num2 的长度都小于 5100
num1 和num2 都只包含数字0-9
num1 和num2 都不包含任何前导零
你不能使用任何內建 BigInteger 库，也不能直接将输入的字符串转换为整数形式
*/

// 思路：直接每个数字按位做加法，必要的话可以用一个 []int 中间变量存每个值
func addStrings(num1 string, num2 string) string {
	len1 := len(num1)
	len2 := len(num2)
	var result string
	var carry int
	for i, j := len1-1, len2-1; i >= 0 || j >= 0 || carry != 0; i, j = i-1, j-1 {
		var a, b int
		if i >= 0 {
			a = int(num1[i] - '0')
		}
		if j >= 0 {
			b = int(num2[j] - '0')
		}
		n := a + b + carry
		carry = n / 10
		n = n % 10
		result = strconv.Itoa(n) + result
	}

	return result
}

func main() {

}
