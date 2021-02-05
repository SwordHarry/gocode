package main

/**
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例 1：
输入：head = [1,2,3,4]
输出：[2,1,4,3]

示例 2：
输入：head = []
输出：[]

示例 3：
输入：head = [1]
输出：[1]

提示：
链表中节点的数目在范围 [0, 100] 内
0 <= Node.val <= 100

进阶：你能在不修改链表节点值的情况下解决这个问题吗?（也就是说，仅修改节点本身。）
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 思路1：迭代，其实是 k 个一组反转链表的特殊版本，k = 2时的解法
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	hair := &ListNode{
		Next: head,
	}
	pre, tail := hair, hair
	for tail != nil {
		k := 2
		for k != 0 {
			tail = tail.Next
			k--
			if tail == nil {
				return hair.Next
			}
		}
		end := tail.Next
		node := pre.Next
		curHead, curTail := reverseListNode(node, tail)
		pre.Next = curHead
		curTail.Next = end
		// 滑动窗口
		pre = curTail
		tail = curTail
	}
	return hair.Next
}

func reverseListNode(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {

	var pre *ListNode
	node := head
	for pre != tail {
		next := node.Next
		node.Next = pre
		pre = node
		node = next
	}
	return tail, head
}

// 思路2：递归实现
func swapPairs2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	nextHead := head.Next
	head.Next = swapPairs2(nextHead.Next)
	nextHead.Next = head

	return nextHead
}

func main() {

}
