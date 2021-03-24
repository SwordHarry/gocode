package main

/**
给定整数 n 和 k，找到 1 到 n 中字典序第 k 小的数字。
注意：1 ≤ k ≤ n ≤ 109。

示例 :
输入:
n: 13   k: 2
输出:
10

解释:
字典序的排列是 [1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9]，所以第二小的数字是 10。
*/
func findKthNumber(n int, k int) int {
	index, num := 1, 1
	for index < k {
		// *10 的增幅是小于 +1 的增幅
		count := getCount(num, n)
		if index+count > k {
			num *= 10
			index++
		} else {
			num++
			index += count
		}
	}
	return num
}

// 获取该前缀节点数
func getCount(prefix, n int) int {
	count := 0
	// 注意是 cur <= n
	for cur, next := prefix, prefix+1; cur <= n; cur, next = 10*cur, 10*next {
		count += min(n+1, next) - cur
	}
	return count
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {

}
