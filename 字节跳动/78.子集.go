package main

/**
给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

示例 1：
输入：nums = [1,2,3]
输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

示例 2：
输入：nums = [0]
输出：[[],[0]]

提示：
1 <= nums.length <= 10
-10 <= nums[i] <= 10
nums 中的所有元素 互不相同
*/

/**
思路：递归回溯，对于 123 来说，当前状态为 state，下一个状态为 state+1
	则当前状态 state 到 n 的数不确认状态，但 0 到 state 的数确定了状态，每个数的状态只有 0 和 1 两种
	则 0 和 1 两种状态都要进行递归回溯
	当前数状态为 1
	dfs(state+1, n)
	当前数状态为 0
	dfs(state+1, n)
*/
func subsets(nums []int) [][]int {
	var (
		result [][]int
		temp   []int
		dfs    func(count int)
	)
	length := len(nums)
	dfs = func(count int) {
		if count == length {
			res := make([]int, len(temp))
			copy(res, temp)
			result = append(result, res)
			return
		}

		temp = append(temp, nums[count])
		dfs(count + 1)
		temp = temp[:len(temp)-1]
		dfs(count + 1)
	}

	dfs(0)

	return result
}

func main() {

}
