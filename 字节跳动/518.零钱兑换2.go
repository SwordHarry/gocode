package main

/**
给定不同面额的硬币和一个总金额。写出函数来计算可以凑成总金额的硬币组合数。假设每一种面额的硬币有无限个。

示例 1:
输入: amount = 5, coins = [1, 2, 5]
输出: 4
解释: 有四种方式可以凑成总金额:
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1

示例 2:
输入: amount = 3, coins = [2]
输出: 0
解释: 只用面额2的硬币不能凑成总金额3。

示例 3:
输入: amount = 10, coins = [10]
输出: 1

注意:
你可以假设：
0 <= amount (总金额) <= 5000
1 <= coin (硬币面额) <= 5000
硬币种类不超过 500 种
结果符合 32 位符号整数
*/

/**
思路：动态规划，dp[i]
	该题和 爬楼梯 类似，但是 爬楼梯 为排列数，该题为 组合数
*/
func change(amount int, coins []int) int {
	dp := make([]int, amount+1)

	dp[0] = 1
	length := len(coins)
	// 组合数
	for j := 0; j < length; j++ {
		for i := coins[j]; i <= amount; i++ {
			dp[i] += dp[i-coins[j]]
		}
	}
	// 排列数，错误答案，仅适用于爬楼梯
	// for i := 1;i <= amount;i++ {
	//     for j := 0;j < length;j++ {
	//         if i - coins[j] >= 0 {
	//             dp[i] += dp[i-coins[j]]
	//         }
	//     }
	// }

	return dp[amount]
}

func main() {

}
