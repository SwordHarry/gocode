package main

/**
假设按照升序排序的数组在预先未知的某个点上进行了旋转。例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] 。
请找出其中最小的元素。

示例 1：
输入：nums = [3,4,5,1,2]
输出：1

示例 2：
输入：nums = [4,5,6,7,0,1,2]
输出：0

示例 3：
输入：nums = [1]
输出：1

提示：
1 <= nums.length <= 5000
-5000 <= nums[i] <= 5000
nums 中的所有整数都是 唯一 的
nums 原来是一个升序排序的数组，但在预先未知的某个点上进行了旋转
*/

/**
感觉虽然也是二分，但是和 在旋转排序数组中寻找目标值 的细节不同
*/
func findMin(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}
	l, r := 0, length-1
	// 可能没有旋转
	if nums[l] < nums[r] {
		return nums[l]
	}
	for l < r {
		mid := (l + r) / 2
		// 先直接找变化点
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}
		if nums[mid-1] > nums[mid] {
			return nums[mid]
		}
		// 再二分
		if nums[l] < nums[mid] {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return nums[l]
}

func main() {

}
