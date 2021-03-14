package main

/**
输入一个整数数组，判断该数组是不是某二叉搜索树的后序遍历结果。
如果是则返回 true，否则返回 false。假设输入的数组的任意两个数字都互不相同。
参考以下这颗二叉搜索树：

     5
    / \
   2   6
  / \
 1   3

示例 1：
输入: [1,6,3,2,5]
输出: false

示例 2：
输入: [1,3,2,6,5]
输出: true

提示：
数组长度 <= 1000
*/

/*
思路：子问题递归，类似 前序后序构造二叉树
	后序遍历，根节点为数组最后一位
	二叉搜索树：左子树 < 根 < 右子树
	则可以先找出第一个比 根大的元素下标 k
	则 [i:k] 为左子树， [k:n-1] 为右子树
	左右子树又是两个子问题
*/
func verifyPostorder(postorder []int) bool {
	length := len(postorder)
	if length < 2 {
		return true
	}

	var dfs func(postorder []int) bool
	dfs = func(postorder []int) bool {
		length := len(postorder)
		if length < 2 {
			return true
		}
		i := 0
		root := postorder[len(postorder)-1]
		for ; i < length-1 && postorder[i] < root; i++ {
		}

		left := postorder[:i]
		right := postorder[i : length-1]
		for _, v := range right {
			if v < root {
				return false
			}
		}
		return dfs(left) && dfs(right)
	}
	return dfs(postorder)
}

func main() {

}
