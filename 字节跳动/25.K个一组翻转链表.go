package main

import (
	"fmt"
)

/**
给你一个链表，每k个节点一组进行翻转，请你返回翻转后的链表。
k是一个正整数，它的值小于或等于链表的长度。
如果节点总数不是k的整数倍，那么请将最后剩余的节点保持原有顺序。
示例：
给你这个链表：1->2->3->4->5
当k= 2 时，应当返回: 2->1->4->3->5
当k= 3 时，应当返回: 3->2->1->4->5
说明：
你的算法只能使用常数的额外空间。
你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。
*/

/**
 * Definition for singly-linked list.
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// 思路：四指针加反转链表
func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := new(ListNode)
	hair.Next = head

	pre, cur, tail := hair, hair, hair
	var end *ListNode

	for cur != nil {
		cur = pre.Next
		for count := k; count > 0; count-- {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		end = tail.Next
		cur, tail = reverse(cur, tail)
		pre.Next = cur
		tail.Next = end
		pre = tail
	}

	return hair.Next
}

func reverse(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
	var pre *ListNode
	cur := head
	for pre != tail {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return tail, head
}

func main() {
	head := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	result := reverseKGroup(head, 2)
	for result != nil {
		fmt.Print(result.Val, " ")
		result = result.Next
	}
}
