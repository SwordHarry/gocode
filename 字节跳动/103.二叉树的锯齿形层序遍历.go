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
// 自定义封装的双向队列
type twoWayQueue struct {
	q []*TreeNode
}

func (t *twoWayQueue) PositivePop() *TreeNode {
	r := t.q[0]
	t.q = t.q[1:]
	return r
}

func (t *twoWayQueue) PositivePush(node *TreeNode) {
	t.q = append(t.q, node)
}

func (t *twoWayQueue) NegativePop() *TreeNode {
	lastIndex := len(t.q) - 1
	r := t.q[lastIndex]
	t.q = t.q[:lastIndex]
	return r
}

func (t *twoWayQueue) NegativePush(node *TreeNode) {
	t.q = append([]*TreeNode{node}, t.q...)
}

func (t *twoWayQueue) Len() int {
	return len(t.q)
}
*/

// go 有内置队列 container/list，也可以自己封装一个 双向队列
func zigzagLevelOrder(root *TreeNode) [][]int {
	queue := list.New()
	var result [][]int
	if root == nil {
		return result
	}
	node := root
	queue.PushBack(node)
	flag := true
	for queue.Len() > 0 {
		le := queue.Len()
		var temp []int
		for i := 0; i < le; i++ {
			if flag {
				// left pop
				e := queue.Front()
				node = e.Value.(*TreeNode)
				queue.Remove(e)
				if node.Left != nil {
					queue.PushBack(node.Left)
				}
				if node.Right != nil {
					queue.PushBack(node.Right)
				}
			} else {
				// right pop
				e := queue.Back()
				node = e.Value.(*TreeNode)
				queue.Remove(e)
				if node.Right != nil {
					queue.PushFront(node.Right)
				}
				if node.Left != nil {
					queue.PushFront(node.Left)
				}
			}
			temp = append(temp, node.Val)
		}
		if len(temp) > 0 {
			result = append(result, temp)
		}
		flag = !flag
	}

	return result
}

func main() {

}
