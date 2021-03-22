package main

/**
给你一棵二叉树的根节点 root ，请你返回 层数最深的叶子节点的和 。

示例 1：
输入：root = [1,2,3,4,5,null,6,7,null,null,null,null,8]
输出：15

示例 2：
输入：root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]
输出：19

提示：
树中节点数目在范围 [1, 104] 之间。
1 <= Node.val <= 100

链接：https://leetcode-cn.com/problems/deepest-leaves-sum
*/

/**
 * 层序遍历即可
 */
func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return root.Val
	}

	var (
		que    []*TreeNode
		result int
	)

	que = append(que, root)
	for len(que) > 0 {
		length := len(que)
		var node *TreeNode
		result = 0
		for i := 0; i < length; i++ {
			node = que[0]
			que = que[1:]
			result += node.Val
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
		}
	}

	return result
}

func main() {

}
