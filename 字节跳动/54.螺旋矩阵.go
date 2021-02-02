package main

/**
给你一个 m 行 n 列的矩阵matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
示例 1：
输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
输出：[1,2,3,6,9,8,7,4,5]

示例 2：
输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
输出：[1,2,3,4,8,12,11,10,9,5,6,7]

提示：
m == matrix.length
n == matrix[i].length
1 <= m, n <= 10
-100 <= matrix[i][j] <= 100
*/

// 思路：直接遍历 时间复杂度：O(1)，空间复杂度：O(m*n)
// 注意四个方向的边界问题

const (
	right = iota
	down
	left
	up
)

func spiralOrder(matrix [][]int) []int {
	row := len(matrix)
	column := len(matrix[0])
	var i, j int
	var result []int
	var direction int
	// 四个边界
	lb, ub := 0, 0
	rb := column - 1
	db := row - 1
	for len(result) != row*column {
		switch direction {
		case right:
			result = append(result, matrix[i][j])
			j++
			if j == rb+1 {
				direction = down
				j--
				i++
				ub++
			}
		case down:
			result = append(result, matrix[i][j])
			i++
			if i == db+1 {
				direction = left
				i--
				j--
				rb--
			}
		case left:
			result = append(result, matrix[i][j])
			j--
			if j == lb-1 {
				direction = up
				j++
				i--
				db--
			}
		case up:
			result = append(result, matrix[i][j])
			i--
			if i == ub-1 {
				direction = right
				i++
				j++
				lb++
			}
		}
	}

	return result
}

func main() {

}
