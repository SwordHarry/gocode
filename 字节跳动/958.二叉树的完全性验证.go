package main

/**
给定一个二叉树，确定它是否是一个完全二叉树。
百度百科中对完全二叉树的定义如下：
若设二叉树的深度为 h，除第 h 层外，其它各层 (1～h-1) 的结点数都达到最大个数，第 h 层所有的结点都连续集中在最左边，这就是完全二叉树。
（注：第 h 层可能包含 1~2h个节点。）
示例 1：
输入：[1,2,3,4,5,6]
输出：true
解释：最后一层前的每一层都是满的（即，结点值为 {1} 和 {2,3} 的两层），且最后一层中的所有结点（{4,5,6}）都尽可能地向左。

示例 2：
输入：[1,2,3,4,5,null,7]
输出：false
解释：值为 7 的结点没有尽可能靠向左侧。

提示：
树中将会有 1 到 100 个结点。
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 思路：对于完全二叉树，可以用数组进行表示，父节点为 n，则左孩子和右孩子分别为 2n, 2n+1
// 即可以通过节点总个数和下标的对应关系判断
func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	node := root
	var s []*TreeNode
	s = append(s, node)
	var c []int
	c = append(c, 1)
	count := 0
	for len(s) > 0 {
		node = s[0]
		s = s[1:]
		curNum := c[count]
		count++
		if curNum > count {
			return false
		}
		if node.Left != nil {
			s = append(s, node.Left)
			c = append(c, curNum*2)
		}
		if node.Right != nil {
			s = append(s, node.Right)
			c = append(c, curNum*2+1)
		}
	}
	return true
}

func main() {

}
