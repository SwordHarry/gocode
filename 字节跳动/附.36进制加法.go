package main

import "fmt"

/**
36进制由0-9，a-z，共36个字符表示。
要求按照加法规则计算出任意两个36进制正整数的和，如1b + 2x = 48  （解释：47+105=152）
要求：不允许使用先将36进制数字整体转为10进制，相加后再转回为36进制的做法
*/

/**
思路：和十进制字符串的加法一致，只是需要 getInt 和 getByte 转换成36进制，并且求模和除数的时候也要转换
*/
func getInt(b byte) int {
	if b >= '0' && b <= '9' {
		return int(b - '0')
	}
	return int(b-'a') + 10
}

func getByte(n int) byte {
	if n >= 0 && n <= 9 {
		return byte(n + '0')
	}
	return byte(n - 10 + 'a')
}

func add36Strings(num1, num2 string) string {
	len1, len2 := len(num1), len(num2)
	var result []byte
	var carry int

	for i, j := len1-1, len2-1; i >= 0 || j >= 0 || carry > 0; i, j = i-1, j-1 {
		var a, b int
		if i >= 0 {
			a = getInt(num1[i])
		}
		if j >= 0 {
			b = getInt(num2[j])
		}

		n := a + b + carry
		carry = n / 36
		n = n % 36
		result = append(result, getByte(n))
	}
	return reverse(result)
}

func reverse(num []byte) string {
	length := len(num)
	for i := 0; i < length/2; i++ {
		num[i], num[length-1-i] = num[length-1-i], num[i]
	}
	return string(num)
}

func main() {
	fmt.Println(add36Strings("1b", "2x"))
}
