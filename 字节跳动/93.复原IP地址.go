package main

import (
	"strconv"
	"strings"
)

/**
给定一个只包含数字的字符串，用以表示一个 IP 地址，返回所有可能从 s 获得的 有效 IP 地址 。你可以按任何顺序返回答案。
有效 IP 地址 正好由四个整数（每个整数位于 0 到 255 之间组成，且不能含有前导 0），整数之间用 '.' 分隔。
例如："0.1.2.201" 和 "192.168.1.1" 是 有效 IP 地址，但是 "0.011.255.245"、"192.168.1.312" 和 "192.168@1.1" 是 无效 IP 地址。

示例 1：
输入：s = "25525511135"
输出：["255.255.11.135","255.255.111.35"]

示例 2：
输入：s = "0000"
输出：["0.0.0.0"]

示例 3：
输入：s = "1111"
输出：["1.1.1.1"]

示例 4：
输入：s = "010010"
输出：["0.10.0.10","0.100.1.0"]

示例 5：
输入：s = "101023"
输出：["1.0.10.23","1.0.102.3","10.1.0.23","10.10.2.3","101.0.2.3"]

提示：
0 <= s.length <= 3000
s 仅由数字组成
*/

/**
回溯+剪枝，遇到前导零直接返回
需要注意使用了 strconv.Atoi() int,error 和 strings.Join() string
*/

func restoreIpAddresses(s string) []string {
	length := len(s)
	if length < 4 {
		return nil
	}
	var (
		result []string
		temp   []string
		dfs    func(str string, count int)
	)
	dfs = func(str string, count int) {
		sLen := len(str)
		if count == 4 {
			if sLen == 0 {
				result = append(result, strings.Join(temp, "."))
			}
			return
		}
		for i := 0; i < 3 && i < len(str); i++ {
			// 遇到前导零直接返回
			if str[0] == '0' && i > 0 {
				return
			}
			num, _ := strconv.Atoi(str[:i+1])
			nextS := str[i+1:]
			if num >= 0 && num <= 255 {
				temp = append(temp, str[:i+1])
				dfs(nextS, count+1)
				temp = temp[:len(temp)-1]
			}
		}
	}
	dfs(s, 0)
	return result
}

func main() {

}
