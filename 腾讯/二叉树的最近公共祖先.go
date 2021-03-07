package main

import (
	"fmt"
)

type TreeNode struct {
	val         int
	left, right *TreeNode
}

func getAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.val == p.val || root.val == q.val {
		return root
	}
	// 递归
	l := getAncestor(root.left, p, q)
	r := getAncestor(root.right, p, q)

	// 分类讨论
	if l != nil && r != nil {
		return root
	} else if l != nil && r == nil {
		return l
	} else if l == nil && r != nil {
		return r
	} else {
		return nil
	}
}

// test
func main() {
	q := &TreeNode{
		val: 4,
	}
	p := &TreeNode{
		val: 5,
		left: &TreeNode{
			val: 6},
		right: &TreeNode{
			val: 2,
			left: &TreeNode{
				val: 7,
			},
			right: q,
		},
	}
	root := &TreeNode{
		val:  3,
		left: p,
		right: &TreeNode{
			val: 1,
			left: &TreeNode{
				val: 0,
			},
			right: &TreeNode{
				val: 8,
			},
		},
	}
	fmt.Println(getAncestor(root, p, q).val)
}
