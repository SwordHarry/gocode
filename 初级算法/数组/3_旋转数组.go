package main

import "fmt"

/*
struct {
 ptr -> array
 len int
 cap int
}
*/

func rotate(nums []int, k int) {
	length := len(nums)
	k = k % length
	left := nums[length-k:]
	right := nums[:length-k]
	reverse(left)
	reverse(right)
	reverse(nums)
	//nums = append(nums[length-k:], nums[:length-k]...)
	//fmt.Println(nums) // [5 6 7 1 2 3 4]
}

func reverse(nums []int) {
	length := len(nums)
	if length == 0 {
		return
	}

	for i := 0; i < length/2; i++ {
		nums[i], nums[length-i-1] = nums[length-i-1], nums[i]
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(nums, 3)
	fmt.Println(nums) // [1 2 3 4 5 6 7]

}
