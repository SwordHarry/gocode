package main

/**
给定一个单链表L：L0→L1→…→Ln-1→Ln ，
将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例1:
给定链表 1->2->3->4, 重新排列为 1->4->2->3.

示例 2:
给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 思路：从中间拆分链表，然后反转后半部分，然后奇偶合并
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	pre, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		pre = pre.Next
		fast = fast.Next.Next
	}
	slow := pre.Next
	pre.Next = nil
	mid := reverse(slow)
	head = merge(head, mid)
}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	node := head
	for node != nil {
		next := node.Next
		node.Next = pre
		pre = node
		node = next
	}

	return pre
}

func merge(headA *ListNode, headB *ListNode) *ListNode {
	pre := new(ListNode)
	cur := pre
	var flag bool
	for headA != nil || headB != nil {
		if flag {
			cur.Next = headB
			if headB != nil {
				headB = headB.Next
			}
		} else {
			cur.Next = headA
			if headA != nil {
				headA = headA.Next
			}
		}
		cur = cur.Next
		flag = !flag
	}
	return pre.Next
}

func main() {

}
