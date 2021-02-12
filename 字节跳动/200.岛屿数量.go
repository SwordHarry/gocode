package main

/**
给你一个由'1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。
岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。
此外，你可以假设该网格的四条边均被水包围。

示例 1：
输入：grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
输出：1

示例 2：
输入：grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
输出：3

提示：
m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] 的值为 '0' 或 '1'
*/

// 思路：dfs，注意判断边界
func numIslands(grid [][]byte) int {
	row := len(grid)
	column := len(grid[0])
	var result int
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if grid[i][j] == '1' {
				result++
				grid = dfs(grid, i, j, row, column)
			}
		}
	}

	return result
}

func dfs(grid [][]byte, i, j, r, c int) [][]byte {
	if grid[i][j] == '0' {
		return grid
	}
	grid[i][j] = '0'
	if i > 0 {
		grid = dfs(grid, i-1, j, r, c)
	}
	if i < r-1 {
		grid = dfs(grid, i+1, j, r, c)
	}
	if j > 0 {
		grid = dfs(grid, i, j-1, r, c)
	}
	if j < c-1 {
		grid = dfs(grid, i, j+1, r, c)
	}
	return grid
}

func main() {

}
