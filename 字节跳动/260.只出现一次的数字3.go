package main

/**
给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。 找出只出现一次的那两个元素。你可以按 任意顺序 返回答案。
进阶：你的算法应该具有线性时间复杂度。你能否仅使用常数空间复杂度来实现？

示例 1：
输入：nums = [1,2,1,3,2,5]
输出：[3,5]
解释：[5, 3] 也是有效的答案。

示例 2：
输入：nums = [-1,0]
输出：[-1,0]

示例 3：
输入：nums = [0,1]
输出：[1,0]

提示：
2 <= nums.length <= 3 * 104
-231 <= nums[i] <= 231 - 1
除两个只出现一次的整数外，nums 中的其他数字都出现两次
*/

/**
思路：全员异或后，将异或数的第一个位为1 的数取出，再次分组异或
*/

func singleNumber(nums []int) []int {

	length := len(nums)
	n := nums[0]
	for i := 1; i < length; i++ {
		n ^= nums[i]
	}
	div := 1
	for div&n == 0 {
		div <<= 1
	}
	var a, b int
	for i := 0; i < length; i++ {
		if nums[i]&div != 0 {
			a ^= nums[i]
		} else {
			b ^= nums[i]
		}
	}
	return []int{a, b}
}

func main() {

}
