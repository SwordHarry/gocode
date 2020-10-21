package main

import "fmt"

/**
算法设计第四次作业
题目:输出两个字符串的最长公共子序列
要求1:不使用辅助数组
要求2(选作): 不使用m*n数组，使用2*n数组
*/

/**
解题思路：
递推式：
	c(i,j) = c(i-1,j-1) + 1 {xi == yj}
	c(i,j) = max(c(i-1,j),c(i,j-1)) {xi != yj}
边界条件：
	c(i,0) = c(0,j) = 0
*/

// 要求1：不使用辅助数组
func LongestSubsequence(x, y string) string {
	n, m := len(x), len(y)
	// 初始化数组切片
	// go 自动初始化数组元素为 0
	c := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		c[i] = make([]int, m+1)
	}

	// 递推式
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if x[i-1] == y[j-1] {
				c[i][j] = c[i-1][j-1] + 1
				print(x[i])
			} else {
				c[i][j] = max(c[i-1][j], c[i][j-1])
			}
		}
	}

	// 再次遍历 c 构造字符串
	var s string

	for i, j := n, m; i >= 1 && j >= 1; {
		if x[i-1] == y[j-1] {
			s = string(x[i-1]) + s
			i--
			j--
		} else {
			if c[i-1][j] >= c[i][j-1] {
				i--
			} else {
				j--
			}
		}
	}

	return s
}

// 要求2：不使用m*n数组，使用2*n数组
func LongestSubsequence2(x, y string) int {
	n, m := len(x), len(y)
	// 初始化数组切片，只能使用 2 * m 个空间
	// go 自动初始化元素为0
	c := make([][]int, 2)
	for i := 0; i <= 1; i++ {
		c[i] = make([]int, m)
	}

	// 递推式
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			// 需要注意下标越界
			if x[i] == y[j] {
				if i*j == 0 {
					c[i][j] = 1
				} else {
					c[1][j] = c[0][j-1] + 1
				}
			} else {
				if j == 0 {
					c[1][j] = 1
				} else {
					c[1][j] = max(c[0][j], c[1][j-1])
				}

			}
		}
		// 滚动数组
		for j := 0; j < m; j++ {
			c[0][j] = c[1][j]
		}
	}

	return c[1][m-1]
}

func max(a int, b int) int {
	if a >= b {
		return a
	}
	return b
}
func main() {
	x, y := "acdefggh", "cegdhfghb"
	fmt.Println("最长公共子序列为：", LongestSubsequence(x, y))
	fmt.Println("最长公共子序列长为：", LongestSubsequence2(x, y))
}
