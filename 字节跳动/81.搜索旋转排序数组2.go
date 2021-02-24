package main

/**
假设按照升序排序的数组在预先未知的某个点上进行了旋转。
( 例如，数组[0,0,1,2,2,5,6]可能变为[2,5,6,0,0,1,2])。
编写一个函数来判断给定的目标值是否存在于数组中。若存在返回true，否则返回false。

示例1:
输入: nums = [2,5,6,0,0,1,2], target = 0
输出: true

示例2:
输入: nums = [2,5,6,0,0,1,2], target = 3
输出: false

进阶:
这是 搜索旋转排序数组的延伸题目，本题中的 nums 可能包含重复元素。
这会影响到程序的时间复杂度吗？会有怎样的影响，为什么？
*/

func search2(nums []int, target int) bool {
	length := len(nums)
	if length == 0 {
		return false
	}

	i, j := 0, length-1
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] == target {
			return true
		}
		// 注意 [1,0,1,1,1] 的情况，直接略过
		if nums[i] == nums[mid] {
			i++
			continue
		}
		if nums[i] < nums[mid] {
			if nums[mid] > target && nums[i] <= target {
				j = mid - 1
			} else {
				i = mid + 1
			}
		} else {
			if nums[mid] < target && nums[j] >= target {
				i = mid + 1
			} else {
				j = mid - 1
			}
		}
	}

	return false
}

func main() {
	search2([]int{3, 1}, 1)
}
