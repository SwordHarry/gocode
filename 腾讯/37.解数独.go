package main

/**
编写一个程序，通过填充空格来解决数独问题。
一个数独的解法需遵循如下规则：
数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
空白格用 '.' 表示。

一个数独。
答案被标成红色。

提示：
给定的数独序列只包含数字 1-9 和字符 '.' 。
你可以假设给定的数独只有唯一解。
给定数独永远是 9x9 形式的。
*/

/*
思路：和 n 后问题本质是一致的，需要用辅助数组存 行列小九宫格 冲突
	额外需要注意的是，还要用辅助数组存 '.' 的位置，而不用每次从头开始遍历，还要一个终止条件 valid，不然 dfs 会一直覆盖原答案
*/
func solveSudoku(board [][]byte) {
	// 分别用来记录 行,列,小九宫格,所有'.'的位置,递归函数，是否继续迭代
	var (
		line   [9][9]bool
		column [9][9]bool
		block  [3][3][9]bool
		space  [][2]int
		dfs    func(pos int) bool
		valid  bool
	)

	// 1. 找出所有点的位置 以及 初始化填充区
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				space = append(space, [2]int{i, j})
			} else {
				num := board[i][j] - '1'
				line[i][num] = true
				column[j][num] = true
				block[i/3][j/3][num] = true
			}
		}
	}

	// 2. 递归
	dfs = func(pos int) bool {
		if pos == len(space) {
			valid = true
			return true
		}

		x, y := space[pos][0], space[pos][1]
		for i := 1; i <= 9; i++ {
			if !line[x][i-1] && !column[y][i-1] && !block[x/3][y/3][i-1] && !valid {
				line[x][i-1] = true
				column[y][i-1] = true
				block[x/3][y/3][i-1] = true
				board[x][y] = byte(i + '0')
				dfs(pos + 1)
				line[x][i-1] = false
				column[y][i-1] = false
				block[x/3][y/3][i-1] = false
			}
		}
		return false
	}

	dfs(0)
}

func main() {

}
