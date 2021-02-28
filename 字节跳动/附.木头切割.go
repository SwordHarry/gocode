package main

import "fmt"

/**
给定长度为n的数组，每个元素代表一个木头的长度，木头可以任意截断，从这堆木头中截出至少k个相同长度为m的木块。已知k，求max(m)。
输入两行，第一行n, k，第二行为数组序列。输出最大值。
输入
5 5
4 7 2 10 5
输出
4
解释：如图，最多可以把它分成5段长度为4的木头
ps:数据保证有解，即结果至少是1。
*/

/**
思路：二分查找，切分的木头块为 n := woods[i] / mid，需要保证 n>=k，此时要求 max(mid)
*/
func splitWood(woods []int, k int) int {
	length := len(woods)
	var result int
	if length == 0 {
		return result
	}
	l, r := 1, getMaxOfList(woods...)
	for l < r {
		mid := (l + r) / 2
		var count int
		for i := 0; i < length; i++ {
			count += woods[i] / mid
		}
		if count >= k {
			result = getMaxOfList(result, mid)
			l = mid + 1
		} else {
			r = mid
		}
	}
	return result
}

func getMaxOfList(nums ...int) int {
	result := nums[0]
	for _, v := range nums {
		if v > result {
			result = v
		}
	}
	return result
}

func main() {
	fmt.Println(splitWood([]int{4, 7, 2, 10, 5}, 5))
}
