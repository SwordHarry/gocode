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

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k < 2 {
		return head
	}
	hair := &ListNode{
		Val:  0,
		Next: head,
	}
	// 四指针，pre -> node -> ... -> tail -> end
	pre := hair
	node := head
	end := head
	for node != nil {
		tail := node
		tempK := k - 1
		// 切分链表段
		for tempK > 0 {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
			tempK--
		}
		end = tail.Next
		// 反转链表并拼接
		curHead, curTail := reverse(node, tail)
		pre.Next = curHead
		curTail.Next = end
		// 窗口滑动
		pre = curTail // 注意这里的逻辑，此时 分片的链表 已经被反转，pre 应该指向新分片链表的尾部
		node = end
	}

	return hair.Next
}

// 反转链表，返回反转后的链表头尾
func reverse(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
	hair := &ListNode{
		Val:  0,
		Next: head,
	}
	pre := hair
	node := head
	next := node.Next
	for pre != tail {
		node.Next = pre
		pre = node
		node = next
		if next != nil {
			next = next.Next
		}
	}
	return tail, hair.Next
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
