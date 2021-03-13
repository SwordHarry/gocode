package main

/**
实现获取 下一个排列 的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。
如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
必须 原地 修改，只允许使用额外常数空间。

示例 1：
输入：nums = [1,2,3]
输出：[1,3,2]

示例 2：
输入：nums = [3,2,1]
输出：[1,2,3]

示例 3：
输入：nums = [1,1,5]
输出：[1,5,1]

示例 4：
输入：nums = [1]
输出：[1]

提示：
1 <= nums.length <= 100
0 <= nums[i] <= 100
*/

/**
思路：对于 [1,2,3] 的全排列按序如下：
	[1,2,3]
	[1,3,2]
	[2,1,3]
	[2,3,1]
	[3,1,2]
	[3,2,1]
即下一个序列为，找到一个大于当前序列的新序列，且变大的幅度尽可能小
其实际上是找到一个新的数，该数比原数大，且尽可能接近原数
步骤：
1. 从后向前找第一个降序数，即"小数"，索引为 i
2. 继续从后向前找，找到第一个比"小数"大的数，即"大数"，索引为 j
3. 交换两数，此时[i,n)为降序，再反转为升序，此时该数为大于原数的最小数
4. 若整个数已经为最大数，直接反转即可
*/
func nextPermutation(nums []int) {
	length := len(nums)
	if length < 2 {
		return
	}

	i := length - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := length - 1
		for j > i && nums[j] <= nums[i] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(nums, i+1, length-1)
}

func reverse(nums []int, i, j int) {
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

func main() {

}
