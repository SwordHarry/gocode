package main

/**
根据一棵树的前序遍历与中序遍历构造二叉树。
注意:
你可以假设树中没有重复的元素。
例如，给出
前序遍历 preorder =[3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
返回如下的二叉树：
    3
   / \
  9  20
    /  \
   15   7
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	preLen := len(preorder)
	inLen := len(inorder)
	var root *TreeNode
	if preLen <= 0 || inLen <= 0 {
		return root
	}

	root = new(TreeNode)
	root.Val = preorder[0]
	var i int
	for ; i < inLen; i++ {
		if preorder[0] == inorder[i] {
			break
		}
	}
	leftInOrder := inorder[:i]
	leftLen := len(leftInOrder)
	root.Left = buildTree(preorder[1:1+leftLen], leftInOrder)
	root.Right = buildTree(preorder[1+leftLen:], inorder[i+1:]) // rightInOrder
	return root
}

func main() {

}
