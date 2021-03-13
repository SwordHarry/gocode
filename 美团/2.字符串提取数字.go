package main

import (
	"fmt"
	"sort"
)

/**
从英文和数字混杂的字符串中提取出数字，注意删除前导零，并按序输出
示例：
he12be0c002ae7

输出：
0 2 7 12
*/

func getNumFromStr(s string) []int {
	length := len(s)
	var result []int
	for i := 0; i < length; i++ {
		if isNum(s[i]) {
			var num int
			j := i
			for ; j < length && isNum(s[j]); j++ {
				num = num*10 + int(s[j]-'0')
			}
			i = j
			result = append(result, num)
		}
	}
	sort.Ints(result)
	return result
}

func isNum(b byte) bool {
	return b >= '0' && b <= '9'
}

func main() {
	result := getNumFromStr("127sfjiew2fj09nfew10")
	// 存在全部为字母
	if len(result) > 0 {
		fmt.Println(result)
	} else {
		fmt.Println()
	}
}
