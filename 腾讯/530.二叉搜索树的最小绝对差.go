package main

import "math"

/**
给你一棵所有节点为非负值的二叉搜索树，请你计算树中任意两节点的差的绝对值的最小值。

示例：
输入：

   1
    \
     3
    /
   2

输出：
1
解释：
最小绝对差为 1，其中 2 和 1 的差的绝对值为 1（或者 2 和 3）。

提示：
树中至少有 2 个节点。
本题与 783 https://leetcode-cn.com/problems/minimum-distance-between-bst-nodes/ 相同
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	var stack []*TreeNode
	var pre *TreeNode
	node := root
	result := math.MaxInt64
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		last := len(stack) - 1
		// 中序遍历
		node = stack[last]
		stack = stack[:last]
		if pre != nil {
			result = min(result, node.Val-pre.Val)
		}
		pre = node
		node = node.Right
	}

	return result
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {

}
