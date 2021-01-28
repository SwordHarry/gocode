package main

/**
给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。
示例:
输入:[1,2,3,null,5,null,4]
输出:[1, 3, 4]
解释:

   1            <---
 /   \
2     3         <---
 \     \
  5     4       <---
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 思路：层序遍历返回每层最后一个元素
func rightSideView(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}

	node := root
	var queue []*TreeNode
	queue = append(queue, node)
	for len(queue) > 0 {
		length := len(queue)
		for i := 0; i < length; i++ {
			node = queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, node.Val)
	}
	return result
}

func main() {

}
