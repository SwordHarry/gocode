package main

/**
在一个由 '0' 和 '1' 组成的二维矩阵内，找到只包含 '1' 的最大正方形，并返回其面积。

示例 1：
输入：matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
输出：4

示例 2：
输入：matrix = [["0","1"],["1","0"]]
输出：1

示例 3：
输入：matrix = [["0"]]
输出：0

提示：
m == matrix.length
n == matrix[i].length
1 <= m, n <= 300
matrix[i][j] 为 '0' 或 '1'
*/

// 动规：dp[i][j] = min(dp[i-1][j-1], dp[i][j-1], dp[i-1][j]) + 1
func maximalSquare(matrix [][]byte) int {
	row := len(matrix)
	column := len(matrix[0])

	var resultByte byte
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if i > 0 && j > 0 && matrix[i][j] >= '1' && matrix[i-1][j] >= '1' && matrix[i][j-1] >= '1' {
				matrix[i][j] = min(matrix[i-1][j-1], min(matrix[i-1][j], matrix[i][j-1])) + 1
				resultByte = max(resultByte, matrix[i][j])
			}
			// '1' 也要单独算
			resultByte = max(resultByte, matrix[i][j])
		}
	}
	result := int(resultByte - '0')
	return result * result
}

func max(a, b byte) byte {
	if a > b {
		return a
	}
	return b
}

func min(a, b byte) byte {
	if a > b {
		return b
	}
	return a
}

func main() {

}
