package main

/**
给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。

示例 1:
输入: 1->2->3->4->5->NULL, k = 2
输出: 4->5->1->2->3->NULL
解释:
向右旋转 1 步: 5->1->2->3->4->NULL
向右旋转 2 步: 4->5->1->2->3->NULL

示例 2:
输入: 0->1->2->NULL, k = 4
输出: 2->0->1->NULL
解释:
向右旋转 1 步: 2->0->1->NULL
向右旋转 2 步: 1->2->0->NULL
向右旋转 3 步: 0->1->2->NULL
向右旋转 4 步: 2->0->1->NULL
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 思路：先形成环，然后走 k 次；注意数学 (n-k%n-1)
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	// 1. 形成环
	node := head
	next := node.Next
	n := 1
	for next != nil {
		node = next
		next = node.Next
		n++
	}
	node.Next = head
	node = head

	// 2. 找到新的头，然后断开；注意 n-k%n-1
	for i := 0; i < n-k%n-1; i++ {
		node = node.Next
	}

	head2 := node.Next
	node.Next = nil

	return head2
}

func main() {

}
