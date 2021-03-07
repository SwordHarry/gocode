package main

/**
给定一个 没有重复 数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]
*/

var res [][]int

func permute(nums []int) [][]int {
	length := len(nums)
	res = nil
	if length == 0 {
		return res
	}

	backtrack(nums, 0, length)
	return res
}

func backtrack(nums []int, first, n int) {
	length := len(nums)
	if first == n {
		// 需要 copy，切片 为引用
		temp := make([]int, length)
		copy(temp, nums)
		res = append(res, temp)
		return
	}
	for i := first; i < length; i++ {
		nums[i], nums[first] = nums[first], nums[i]
		// 开头不变，后面做交换
		backtrack(nums, first+1, n)
		nums[i], nums[first] = nums[first], nums[i]
	}
}

func main() {

}
