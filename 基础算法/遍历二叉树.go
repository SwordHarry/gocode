package main

import (
	"container/list"
)

// 前中后序遍历二叉树
type TreeNode struct {
	Val         int
	Left, Right *TreeNode
	visited     bool
}

func preorderTraver(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}

	node := root
	stack := list.New()
	for node != nil || stack.Len() > 0 {
		for node != nil {
			result = append(result, node.Val)
			stack.PushBack(node)
			node = node.Left
		}

		e := stack.Back()
		stack.Remove(e)
		node = e.Value.(*TreeNode).Right
	}
	return result
}

func inOrderTraver(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}

	node := root
	stack := list.New()
	for node != nil || stack.Len() > 0 {
		for node != nil {
			stack.PushBack(node)
			node = node.Left
		}

		e := stack.Back()
		stack.Remove(e)
		node = e.Value.(*TreeNode)
		result = append(result, node.Val)
		node = node.Right
	}
	return result
}

func postOrderTraver(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}

	node := root
	stack := list.New()
	var pre *TreeNode
	for node != nil || stack.Len() > 0 {
		for node != nil {
			stack.PushBack(node)
			node = node.Left
		}

		e := stack.Back()
		cur := e.Value.(*TreeNode)
		right := cur.Right
		// 右孩子为空 或 右孩子在上轮遍历过
		if right == nil || right == pre {
			result = append(result, cur.Val)
			stack.Remove(e)
			pre = cur
		} else {
			node = right
		}
	}

	return result
}

func main() {

}
