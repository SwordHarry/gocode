package main

/**
请根据每日 气温 列表，重新生成一个列表。
对应位置的输出为：要想观测到更高的气温，至少需要等待的天数。
如果气温在这之后都不会升高，请在该位置用 0 来代替。
例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4, 2, 1, 1, 0, 0]。
提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，都是在 [30, 100] 范围内的整数。
*/

/**
思路：单调栈，弹出时计算结果数组
*/
func dailyTemperatures(T []int) []int {
	length := len(T)
	var (
		result = make([]int, length)
		// 存下标
		stack = make([]int, 0, length)
	)
	stack = append(stack, 0)
	// 单调栈，探出时需要将结果赋值
	for i := 1; i < length; i++ {
		for len(stack) > 0 && T[i] > T[stack[len(stack)-1]] {
			last := len(stack) - 1
			result[stack[last]] = i - stack[last]
			stack = stack[:last]
		}
		stack = append(stack, i)
	}

	return result
}

func main() {

}
