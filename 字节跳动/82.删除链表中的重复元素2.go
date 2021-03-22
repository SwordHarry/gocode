package main

/**
给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中没有重复出现的数字。

示例 1:
输入: 1->2->3->3->4->4->5
输出: 1->2->5

示例 2:
输入: 1->1->1->2->3
输出: 2->3
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 思路：模拟，三指针 pre cur next
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	result := &ListNode{
		Next: head,
	}

	pre, cur := result, head

	// 分情况讨论：若 遇到重复节点 和 没遇到重复节点
	for cur != nil {
		next := cur.Next
		if next != nil && next.Val == cur.Val {
			for next != nil && cur.Val == next.Val {
				cur = cur.Next
				next = next.Next
			}
			pre.Next = next
		} else {
			pre = pre.Next
		}
		cur = cur.Next
	}

	return result.Next
}

func main() {

}
