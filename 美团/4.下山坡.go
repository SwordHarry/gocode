package main

import "fmt"

/**
给定图，图表示各节点和边，节点有高度 h，求从任一节点出发，hi > hj，高度严格下降的最长路径，求出经过的节点数
第一行输入 n 个节点 和 m 条边，第二行输入节点高度，剩下行输入边的关系
示例：
5 4
3 2 3 4 5
1 2
2 3
3 4
4 5

输出：
4
解释：
5->4->3->2
*/

func getMostNode(graph [][]bool, heights []int) int {
	length := len(graph)
	var result int
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if i != j && graph[i][j] && heights[i] > heights[j] {
				result = max(result, dfs(graph, heights, j)+1)
			}
		}
	}
	return result
}

func dfs(graph [][]bool, heights []int, j int) int {
	result := 1
	length := len(graph)
	for k := 0; k < length; k++ {
		if j != k && graph[j][k] && heights[j] > heights[k] {
			result = dfs(graph, heights, k) + 1
		}
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	fmt.Println(getMostNode([][]bool{
		{false, true, false, false, false},
		{true, false, true, false, false},
		{false, true, false, true, false},
		{false, false, true, false, true},
		{false, false, false, true, false},
	}, []int{3, 2, 3, 4, 5}))
}
