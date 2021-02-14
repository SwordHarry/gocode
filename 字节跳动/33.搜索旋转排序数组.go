package main

/**
升序排列的整数数组 nums 在预先未知的某个点上进行了旋转（例如， [0,1,2,4,5,6,7] 经旋转后可能变为[4,5,6,7,0,1,2] ）。
请你在数组中搜索target ，如果数组中存在这个目标值，则返回它的索引，否则返回-1。
示例 1：
输入：nums = [4,5,6,7,0,1,2], target = 0
输出：4

示例2：
输入：nums = [4,5,6,7,0,1,2], target = 3
输出：-1

示例 3：
输入：nums = [1], target = 0
输出：-1

提示：
1 <= nums.length <= 5000
-10^4 <= nums[i] <= 10^4
nums 中的每个值都 独一无二
nums 肯定会在某个点上旋转
-10^4 <= target <= 10^4
*/

/**
思路：二分查找
*/
func search(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	if length == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	i, j := 0, length-1
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			// 左半边有序
			if nums[mid] >= target && target >= nums[i] {
				j = mid - 1
			} else {
				// 不在左半边，右半边可以看成一个子问题
				i = mid + 1
			}
		} else {
			// 右半边有序
			if nums[mid] <= target && target <= nums[j] {
				i = mid + 1
			} else {
				j = mid - 1
			}
		}
	}

	return -1
}

func main() {

}
