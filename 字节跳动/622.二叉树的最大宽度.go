package main

/**
给定一个二叉树，编写一个函数来获取这个树的最大宽度。树的宽度是所有层中的最大宽度。
这个二叉树与满二叉树（full binary tree）结构相同，但一些节点为空。
每一层的宽度被定义为两个端点（该层最左和最右的非空节点，两端点间的null节点也计入长度）之间的长度。

示例 1:
输入:
           1
         /   \
        3     2
       / \     \
      5   3     9
输出: 4
解释: 最大值出现在树的第 3 层，宽度为 4 (5,3,null,9)。

示例 2:
输入:

          1
         /
        3
       / \
      5   3
输出: 2
解释: 最大值出现在树的第 3 层，宽度为 2 (5,3)。

示例 3:
输入:
          1
         / \
        3   2
       /
      5
输出: 2
解释: 最大值出现在树的第 2 层，宽度为 2 (3,2)。

示例 4:
输入:
          1
         / \
        3   2
       /     \
      5       9
     /         \
    6           7
输出: 8
解释: 最大值出现在树的第 4 层，宽度为 8 (6,null,null,null,null,null,null,7)。
注意: 答案在32位有符号整数的表示范围内。
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 思路：wfs 利用满二叉树的 i 2i+1 2i+2 的特性，计算宽度
func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	var (
		result int
		tQue   = []*TreeNode{root}
		iQue   = []int{0}
	)

	for len(tQue) > 0 {
		length := len(tQue)
		result = max(result, iQue[length-1]-iQue[0]+1)
		for i := 0; i < length; i++ {
			node := tQue[0]
			tQue = tQue[1:]
			count := iQue[0]
			iQue = iQue[1:]
			if node.Left != nil {
				tQue = append(tQue, node.Left)
				iQue = append(iQue, 2*count+1)
			}
			if node.Right != nil {
				tQue = append(tQue, node.Right)
				iQue = append(iQue, 2*count+2)
			}
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

}
