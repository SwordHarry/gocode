package main

import "fmt"

/**
圆环上有10个点，编号为0~9。从0点出发，每次可以逆时针和顺时针走一步，问走n步回到0点共有多少种走法。

输入: 2
输出: 2
解释：有2种方案。分别是0->1->0和0->9->0
*/

/**
思路：动态规划 dp[i][j] 表示 走 i 步到第 j 点，
	则可以变为类似爬楼梯的子问题，走第 i-1 步到 第 j-1 点 或 走第 i+1 步到 第 j+1 点
	则有：
	边界条件：dp[0][0->n] = 0; dp[1][1] = 1; dp[1][n-1] = 1
	递推式：dp[i][j] = dp[i-1][(j-1+len)%len] + dp[i-1][(j+1)%len]
*/
// k 个点，n 步
func returnToPoint(k, n int) int {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, k)
	}
	dp[0][0] = 1
	dp[1][1] = 1
	dp[1][k-1] = 1

	for i := 2; i <= n; i++ {
		for j := 0; j < k; j++ {
			dp[i][j] = dp[i-1][(j-1+k)%k] + dp[i-1][(j+1)%k]
		}
	}
	return dp[n][0]
}

func main() {
	fmt.Println(returnToPoint(4, 4))
}
