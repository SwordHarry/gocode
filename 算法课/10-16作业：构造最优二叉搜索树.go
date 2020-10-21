package main

import (
	"fmt"
	"math"
)

/**
编程实现最优二叉搜索树，输入输出自己定义
*/

// 节点
type BSTNode struct {
	name        string // 节点名
	left, right *BSTNode
}

// 创建二维数组
func makeFloat2DimensionArray(row, column int) [][]float64 {
	result := make([][]float64, row)
	for i := 0; i < row; i++ {
		temp := make([]float64, column)
		result[i] = append(result[i], temp...)
	}
	return result
}
func makeInt2DimensionArray(row, column int) [][]int {
	result := make([][]int, row)
	for i := 0; i < row; i++ {
		temp := make([]int, column)
		result[i] = append(result[i], temp...)
	}
	return result
}

func optimalBST(pList, qList []float64, n int) ([][]float64, [][]int) {
	// e[i,j] 表示包含关键字 ki ... kj 的最优二叉搜索树进行一次搜索的期望代价，最终希望计算出 e[1,n]
	e := makeFloat2DimensionArray(n+2, n+1)
	// w[i,j] 表示对于包含关键字 ki ... kj 的子树，所有的概率之和
	w := makeFloat2DimensionArray(n+2, n+1)
	// root[i,j] 表示包含关键字 ki ... kj 的子树的根
	root := makeInt2DimensionArray(n+1, n+1)

	for i := 1; i <= n+1; i++ {
		// 对于 j=i−1 的情况，由于子树只包含伪关键字 q[i-1]，所以期望搜索代价为 e[i,i−1]=q[i-1]
		e[i][i-1] = qList[i-1]
		w[i][i-1] = qList[i-1]
	}
	for l := 1; l <= n; l++ {
		for i := 1; i <= n-l+1; i++ {
			j := i + l - 1
			e[i][j] = math.MaxFloat64
			w[i][j] = w[i][j-1] + pList[j] + qList[j]
			// 算法改进：动态规划加速
			if i < j {
				for r := root[i][j-1]; r <= root[i+1][j]; r++ {
					t := e[i][r-1] + e[r+1][j] + w[i][j]
					if t < e[i][j] {
						e[i][j] = t
						root[i][j] = r
					}
				}
			} else {
				e[i][j] = e[i][j-1] + e[i+1][j] + w[i][j]
				root[i][j] = i
			}
		}
	}
	fmt.Println(e, "\n", w, "\n", root)
	return e, root
}

// 打印 最优二叉搜索树
func makeBST(i, j int, root [][]int) {
	fmt.Printf("p%d 为 根\n", root[i][j])
	printBST(i, j, root)
}

func printBST(i, j int, root [][]int) {
	if i <= j {
		r := root[i][j]

		if r > i {
			c := root[i][r-1]
			fmt.Printf("p%d 为 p%d 的左孩子\n", c, r)
			printBST(i, r-1, root)
		} else {
			fmt.Printf("q%d 为 p%d 的左孩子\n", r-1, r)
		}
		if r < j {
			c := root[r+1][j]
			fmt.Printf("p%d 为 p%d 的右孩子\n", c, r)
			printBST(r+1, j, root)
		} else {
			fmt.Printf("q%d 为 p%d 的左孩子\n", r+1, r)
		}
	}
}

func main() {
	// 节点概率 0 表示无此节点
	pList := []float64{0, 0.15, 0.10, 0.05, 0.10, 0.20}
	// 区间概率
	qList := []float64{0.05, 0.1, 0.05, 0.05, 0.05, 0.1}
	e, root := optimalBST(pList, qList, 5)
	fmt.Printf("最优二叉搜索树总概率：%.2f \n", e[1][5])
	makeBST(1, 5, root)
}
