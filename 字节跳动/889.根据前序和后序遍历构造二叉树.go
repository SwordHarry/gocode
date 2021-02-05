package main

/**
返回与给定的前序和后序遍历匹配的任何二叉树。
pre和post遍历中的值是不同的正整数。

示例：
输入：pre = [1,2,4,5,3,6,7], post = [4,5,2,6,7,3,1]
输出：[1,2,3,4,5,6,7]

提示：
1 <= pre.length == post.length <= 30
pre[]和post[]都是1, 2, ..., pre.length的排列
每个输入保证至少有一个答案。如果有多个答案，可以返回其中一个。
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructFromPrePost(pre []int, post []int) *TreeNode {
	preLen := len(pre)
	if preLen == 0 {
		return nil
	}
	if preLen == 1 {
		return &TreeNode{
			Val: pre[0],
		}
	}

	root := &TreeNode{
		Val: pre[0],
	}
	// 左子树
	i := findIndex(pre[1], post) + 1
	root.Left = constructFromPrePost(pre[1:i+1], post[:i])
	// 右子树
	root.Right = constructFromPrePost(pre[i+1:], post[i:len(post)-1])
	return root
}

func findIndex(val int, s []int) int {
	for i, v := range s {
		if val == v {
			return i
		}
	}
	return -1
}
func main() {

}
