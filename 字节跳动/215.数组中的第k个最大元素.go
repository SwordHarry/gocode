package main

// 用快排思路实现，不需要申请额外空间
func findKthLargest(nums []int, k int) int {
	p := nums[0]
	index := partition(nums)
	if index == k-1 {
		return p
	} else if index >= k {
		return findKthLargest(nums[:index], k)
	} else {
		return findKthLargest(nums[index+1:], k-1-index)
	}
}

func partition(nums []int) int {
	p := nums[0]
	length := len(nums)
	i, j := 0, length-1
	for i < j {
		for i < j && nums[j] <= p {
			j--
		}
		nums[i] = nums[j]
		for i < j && nums[i] >= p {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = p
	return i
}

func main() {

}
