package main

/**
给定一个非负整数，你至多可以交换一次数字中的任意两位。返回你能得到的最大值。

示例 1 :
输入: 2736
输出: 7236
解释: 交换数字2和数字7。

示例 2 :
输入: 9973
输出: 9973
解释: 不需要交换。

注意:
给定数字的范围是 [0, 108]
*/

// 贪心，最高位小数 和 最大数最低位 做交换
func maximumSwap(num int) int {
	nums := getInts(num)

	// 记录数字n的最后一次出现的位置
	last := make([]int, 10)
	length := len(nums)
	for i := 0; i < length; i++ {
		last[nums[i]] = i
	}

	// 从后往前找最大数
	for i := 0; i < length; i++ {
		for j := 9; j > nums[i]; j-- {
			// 在当前位置后面并且比当前数字大
			if last[j] > i {
				nums[i], nums[last[j]] = nums[last[j]], nums[i]
				return getInt(nums)
			}
		}
	}
	return num
}

func getInt(nums []int) int {
	var result int
	for _, v := range nums {
		result = result*10 + v
	}
	return result
}

func getInts(num int) []int {
	var result []int
	for num > 0 {
		result = append(result, num%10)
		num /= 10
	}
	length := len(result)
	for i := 0; i < length/2; i++ {
		result[i], result[length-i-1] = result[length-i-1], result[i]
	}
	return result
}

func main() {

}
