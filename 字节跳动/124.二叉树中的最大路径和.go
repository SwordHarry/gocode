package main

import "math"

/**
路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。
同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
路径和 是路径中各节点值的总和。
给你一个二叉树的根节点 root ，返回其 最大路径和 。

示例 1：
输入：root = [1,2,3]
输出：6
解释：最优路径是 2 -> 1 -> 3 ，路径和为 2 + 1 + 3 = 6

示例 2：
输入：root = [-10,9,20,null,null,15,7]
输出：42
解释：最优路径是 15 -> 20 -> 7 ，路径和为 15 + 20 + 7 = 42

提示：
树中节点数目范围是 [1, 3 * 104]
-1000 <= Node.val <= 1000
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/**
此题路径的概念为：任意节点到任意节点，且每个节点至多访问一次，即不能回头
思路：递归 -> 根左右 或 根左 或 根右 或 根
在计算最大值时，算 根左右
返回给父节点时，只能返回 根左 或 根右 或 根
*/
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	result := math.MinInt64

	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := max(dfs(root.Left), 0)
		right := max(dfs(root.Right), 0)
		val := root.Val
		result = max(result, val+left+right)
		return max(val+left, val+right)
	}
	dfs(root)
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

}
