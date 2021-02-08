package main

/**
（这是一个 交互式问题）
给你一个 山脉数组mountainArr，请你返回能够使得mountainArr.get(index)等于target最小的下标 index值。
如果不存在这样的下标 index，就请返回-1。
何为山脉数组？如果数组A 是一个山脉数组的话，那它满足如下条件：
首先，A.length >= 3
其次，在0 < i< A.length - 1条件下，存在 i 使得：
A[0] < A[1] < ... A[i-1] < A[i]
A[i] > A[i+1] > ... > A[A.length - 1]

你将不能直接访问该山脉数组，必须通过MountainArray接口来获取数据：
MountainArray.get(k)- 会返回数组中索引为k的元素（下标从 0 开始）
MountainArray.length()- 会返回该数组的长度

注意：
对MountainArray.get发起超过 100 次调用的提交将被视为错误答案。此外，任何试图规避判题系统的解决方案都将会导致比赛资格被取消。
为了帮助大家更好地理解交互式问题，我们准备了一个样例 “答案”：https://leetcode-cn.com/playground/RKhe3ave，请注意这 不是一个正确答案。
示例 1：
输入：array = [1,2,3,4,5,3,1], target = 3
输出：2
解释：3 在数组中出现了两次，下标分别为 2 和 5，我们返回最小的下标 2。

示例 2：
输入：array = [0,1,2,4,2,1], target = 3
输出：-1
解释：3 在数组中没有出现，返回 -1。

提示：
3 <= mountain_arr.length() <= 10000
0 <= target <= 10^9
0 <= mountain_arr.get(index) <=10^9
*/

/**
 * // This is the MountainArray's API interface.
 * // You should not implement it, or speculate about its implementation
 */

type MountainArray interface {
	get(index int) int
	length() int
}

// 思路： 用二分查找找 [mid] 和 [mid+1] 的关系，直到找到山顶，右半边数组取负数则为递增数组
func findInMountainArray(target int, mountainArr MountainArray) int {
	l, r := 0, mountainArr.length()-1
	var mid int
	for l < r {
		mid = (l + r) / 2
		if mountainArr.get(mid) < mountainArr.get(mid+1) {
			l = mid + 1
		} else {
			r = mid
		}
	}
	leftIndex := binarySearch(target, 0, l, 1, mountainArr)
	if leftIndex != -1 {
		return leftIndex
	}
	rightIndex := binarySearch(target, l, mountainArr.length(), -1, mountainArr)
	if rightIndex != -1 {
		return rightIndex
	}
	return -1
}

// 二分查找：左闭右开
func binarySearch(target, left, right, flag int, mountainArray MountainArray) int {
	target *= flag
	for left < right {
		mid := (left + right) / 2
		midVal := mountainArray.get(mid) * target
		if midVal == target {
			return mid
		} else if midVal > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func main() {

}
