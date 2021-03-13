package main

// 思路：翻转左边，翻转右边，翻转全部
func rotate(nums []int, k int) {
	length := len(nums)
	if length < 2 {
		return
	}

	k = k % length
	if k == 0 {
		return
	}
	reverse(nums[length-k:])
	reverse(nums[:length-k])
	reverse(nums)
}

func reverse(nums []int) {
	length := len(nums)
	for i := 0; i < length/2; i++ {
		j := length - i - 1
		nums[i], nums[j] = nums[j], nums[i]
	}
}

func main() {

}
