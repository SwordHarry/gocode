package main

/**
给定一个二维网格和一个单词，找出该单词是否存在于网格中。
单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

示例:
board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]
给定 word = "ABCCED", 返回 true
给定 word = "SEE", 返回 true
给定 word = "ABCB", 返回 false

提示：
board 和 word 中只包含大写和小写英文字母。
1 <= board.length <= 200
1 <= board[i].length <= 200
1 <= word.length <= 10^3
*/

// 类似于 数海岛 的 dfs
func exist(board [][]byte, word string) bool {

	var (
		dfs         func(i, j, k int) bool
		row, column int
		m           [][]bool
	)

	if len(word) == 0 {
		return true
	}
	row = len(board)
	if row == 0 {
		return false
	}
	column = len(board[0])
	m = make([][]bool, row)
	for i := 0; i < row; i++ {
		m[i] = make([]bool, column)
	}
	dir := [][]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	// 深度遍历
	dfs = func(i, j, k int) bool {
		m[i][j] = true
		defer func() {
			m[i][j] = false
		}()
		if board[i][j] != word[k] {
			return false
		}
		if len(word)-1 == k {
			return true
		}
		var flag bool
		for _, d := range dir {
			newX := i + d[0]
			newY := j + d[1]
			// 未越界且之前没访问过
			if newX >= 0 && newX < row && newY >= 0 && newY < column && !m[newX][newY] {
				if dfs(newX, newY, k+1) {
					flag = true
					break
				}
			}
		}
		return flag
	}

	var result bool
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			if dfs(i, j, 0) {
				result = true
				break
			}
		}
	}
	return result
}

func main() {

}
