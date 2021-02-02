package main

/**
给定一个二叉树和一个目标和，找到所有从根节点到叶子节点路径总和等于给定目标和的路径。
说明:叶子节点是指没有子节点的节点。
示例:
给定如下二叉树，以及目标和sum = 22，

              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
返回:
[
   [5,4,11,2],
   [5,8,4,5]
]
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 思路：非递归 后序遍历 求 路径和
func pathSum(root *TreeNode, targetSum int) [][]int {
	var result [][]int
	var stack []*TreeNode
	var temp int
	node := root
	var pre *TreeNode
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			temp += node.Val
			node = node.Left
		}

		last := len(stack) - 1
		cur := stack[last]
		right := cur.Right
		if right == nil || right == pre {
			if right == nil && cur.Left == nil && temp == targetSum {
				var tempResult []int
				for _, n := range stack {
					tempResult = append(tempResult, n.Val)
				}
				result = append(result, tempResult)
			}
			pre = cur
			stack = stack[:last]
			temp -= cur.Val
		} else {
			node = right
		}

	}

	return result
}

// 递归 dfs 求 路径和；用时更快
func pathSum2(root *TreeNode, targetSum int) [][]int {
	var result [][]int
	var temp []int
	// 闭包方便操作切片
	var dfs func(root *TreeNode, targetSum int)
	dfs = func(root *TreeNode, targetSum int) {
		if root == nil {
			return
		}

		temp = append(temp, root.Val)
		targetSum = targetSum - root.Val
		dfs(root.Left, targetSum)
		dfs(root.Right, targetSum)
		if root.Left == nil && root.Right == nil && targetSum == 0 {
			tempResult := make([]int, len(temp))
			copy(tempResult, temp)
			result = append(result, tempResult)
		}
		temp = temp[:len(temp)-1]
	}
	dfs(root, targetSum)

	return result
}

func main() {

}
