package main

/**
给定一个二叉树，检查它是否是镜像对称的。
例如，二叉树[1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3

但是下面这个[1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3

进阶：
你可以运用递归和迭代两种方法解决这个问题吗？
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 思路1：递归
func isSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}

	if root.Left == nil && root.Right == nil {
		return true
	}

	p, q := root.Left, root.Right
	return check(p, q)

}
func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil && q != nil || p != nil && q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return check(p.Left, q.Right) && check(p.Right, q.Left)
}

// 思路2：迭代，双向队列，注意顺序
func isSymmetric2(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left == nil && root.Right != nil || root.Left != nil && root.Right == nil {
		return false
	}

	var s []*TreeNode
	s = append(s, root.Right)
	s = append([]*TreeNode{root.Left}, s...)
	for len(s) > 0 {
		last := len(s) - 1
		iNode, jNode := s[0], s[last]
		s = s[1:last]
		if iNode == nil && jNode == nil {
			continue
		}
		if iNode == nil && jNode != nil || iNode != nil && jNode == nil {
			return false
		}
		if iNode.Val != jNode.Val {
			return false
		}

		s = append([]*TreeNode{iNode.Left, iNode.Right}, s...)
		s = append(s, jNode.Left, jNode.Right)
	}
	return true
}

func main() {

}
