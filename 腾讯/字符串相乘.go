package main

import (
	"fmt"
	"strconv"
)

/**
字符串相乘
给定两个以字符串形式表示的非负整数num1和num2，返回num1和num2的乘积，它们的乘积也表示为字符串形式。

示例 1:

输入: num1 = "2", num2 = "3"
输出: "6"
示例2:

输入: num1 = "123", num2 = "456"
输出: "56088"
说明：

num1和num2的长度小于110。
num1 和num2 只包含数字0-9。
num1 和num2均不以零开头，除非是数字 0 本身。
不能使用任何标准库的大数类型（比如 BigInteger）或直接将输入转换为整数来处理。
*/

/**
主要思路是：
	最好是先将 string 转化为 []int 的操作
	然后在数组上进行 数位相乘 和 进位 的操作
	注意 一切的进位 ，如 相乘进位 和 相加进位
*/

func multiply(num1 string, num2 string) string {
	// 处理特殊情况
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	if num1 == "1" {
		return num2
	}
	if num2 == "1" {
		return num1
	}

	len1 := len(num1)
	len2 := len(num2)
	// 用数组存储最终结果
	var result [220]int
	var index int
	// 进位
	var carry int
	// 余数位
	var remainder int
	for i := len1 - 1; i >= 0; i-- {
		a, _ := strconv.Atoi(string(num1[i]))
		index = len1 - 1 - i
		for j := len2 - 1; j >= 0; j-- {
			b, _ := strconv.Atoi(string(num2[j]))
			c := a*b + carry
			carry = c / 10
			remainder = c % 10
			result[index] += remainder
			// 注意相加进位
			carry += result[index] / 10
			result[index] = result[index] % 10
			index++
		}
		if carry > 0 {
			result[index] = carry
			index++
			carry = 0
		}
	}
	var resultStr string
	for index = index - 1; index >= 0; index-- {
		resultStr += strconv.Itoa(result[index])
	}
	return resultStr
}

func main() {
	fmt.Println(multiply("123", "456"))
}
