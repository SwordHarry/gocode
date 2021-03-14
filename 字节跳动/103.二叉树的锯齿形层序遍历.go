package main

/**
给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

例如：
给定二叉树[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回锯齿形层序遍历如下：

[
  [3],
  [20,9],
  [15,7]
]
*/

import "container/list"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/**
思路：方法一：本质上是层次遍历，遇到奇数层将结果反转即可
	方法二：也可以使用双向队列，但这样编码相对更复杂，容易出错，建议用方法一
*/
func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var (
		result [][]int
		que    []*TreeNode
		count  int
	)

	node := root
	que = append(que, root)

	for len(que) > 0 {
		length := len(que)
		var temp []int
		for i := 0; i < length; i++ {
			node = que[0]
			que = que[1:]
			if node.Left != nil {
				que = append(que, node.Left)
			}
			if node.Right != nil {
				que = append(que, node.Right)
			}
			temp = append(temp, node.Val)

		}
		if count%2 == 1 {
			reverse(temp)
		}
		count++
		result = append(result, temp)
	}
	return result
}

func reverse(nums []int) {
	length := len(nums)
	for i := 0; i < length/2; i++ {
		nums[i], nums[length-1-i] = nums[length-1-i], nums[i]
	}
}

func main() {

}
