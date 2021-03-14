package main

/**
给定一个无重复元素的数组 candidates 和一个目标数 target，找出 candidates 中所有可以使数字和为 target 的组合。
candidates 中的数字可以无限制重复被选取。
说明：
所有数字（包括 target）都是正整数。
解集不能包含重复的组合。

示例 1：
输入：candidates = [2,3,6,7], target = 7,
所求解集为：
[
  [7],
  [2,2,3]
]

示例 2：
输入：candidates = [2,3,5], target = 8,
所求解集为：
[
 [2,2,2,2],
 [2,3,3],
 [3,5]
]

提示：
1 <= candidates.length <= 30
1 <= candidates[i] <= 200
candidate 中的每个元素都是独一无二的。
1 <= target <= 500
*/

/**
思路：回溯法，并且限定边界
*/

func combinationSum(candidates []int, target int) [][]int {
	length := len(candidates)
	var (
		result [][]int
		temp   []int
		check  func(tar, index int)
	)
	// index 用于指定最小值边界
	check = func(tar, index int) {
		if tar == 0 {
			ans := make([]int, len(temp))
			copy(ans, temp)
			result = append(result, ans)
			return
		}

		for i := index; i < length; i++ {
			// 剪枝
			if tar >= candidates[i] {
				temp = append(temp, candidates[i])
				check(tar-candidates[i], i)
				temp = temp[:len(temp)-1]
			}
		}
	}

	check(target, 0)

	return result
}

func main() {

}
