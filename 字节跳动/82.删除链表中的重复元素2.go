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

// 思路：模拟，
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	hair := new(ListNode)
	hair.Next = head

	pre, cur := hair, head

	for cur != nil && cur.Next != nil {
		next := cur.Next
		if next.Val == cur.Val {
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

	return hair.Next
}

func main() {

}
