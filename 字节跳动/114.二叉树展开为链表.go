package main

/**
给你二叉树的根结点 root ，请你将它展开为一个单链表：
展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。

示例 1：
输入：root = [1,2,5,3,4,null,6]
输出：[1,null,2,null,3,null,4,null,5,null,6]

示例 2：
输入：root = []
输出：[]

示例 3：
输入：root = [0]
输出：[0]

提示：
树中结点数在范围 [0, 2000] 内
-100 <= Node.val <= 100

进阶：你可以使用原地算法（O(1) 额外空间）展开这棵树吗？
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
思路1：非进阶版，直接前序遍历用数组存 时空间复杂度 O(n)
思路2：进阶版，寻找前驱节点
*/
func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	node := root
	for node != nil {
		// 若存在左子树，则找到左子树的最右节点，即前驱
		if node.Left != nil {
			next := node.Left
			pre := next
			for pre.Right != nil {
				pre = pre.Right
			}
			// 将左子树整体移动到右边
			pre.Right = node.Right
			node.Right = next
			node.Left = nil
		}
		node = node.Right
	}
}

func main() {

}
