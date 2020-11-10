package main

import "fmt"

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
}

func longestCommonPrefix(strs []string) string {
	length := len(strs)
	if length == 0 {
		return ""
	}
	minLen := 10000
	result := ""

	for _, s := range strs {
		if minLen > len(s) {
			minLen = len(s)
		}
	}

outer:
	for j := 0; j < minLen; j++ {
		for i := 1; i < length; i++ {
			if strs[0][j] != strs[i][j] {
				break outer
			}
		}
		result += string(strs[0][j])
	}
	return result
}
