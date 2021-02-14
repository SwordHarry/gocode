package main

/**
给定一棵二叉树，你需要计算它的直径长度。一棵二叉树的直径长度是任意两个结点路径长度中的最大值。这条路径可能穿过也可能不穿过根结点。
示例 :
给定二叉树
          1
         / \
        2   3
       / \
      4   5
返回 3, 它的长度是路径 [4,2,1,3] 或者[5,2,1,3]。
注意：两结点之间的路径长度是以它们之间边的数目表示。
*/

// 该题为 124.二叉树最大路径和 的简化版本
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
	var result int
	if root == nil {
		return result
	}

	// dfs
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		result = max(result, left+right)
		return max(left, right) + 1
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
