package main

/**
反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

说明:
1 ≤ m ≤ n ≤ 链表长度。

示例:

输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// 思路：k个一组反转链表的单例版本
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return nil
	}

	hair := new(ListNode)
	hair.Next = head

	var pre, node, tail, end *ListNode
	pre = hair
	node = hair
	tail = hair
	m = m - 1
	for m > 0 {
		pre = pre.Next
		m--
	}
	node = pre.Next
	for n > 0 {
		n--
		tail = tail.Next
	}
	end = tail.Next
	curHead, curTail := reverse(node, end)
	pre.Next = curHead
	curTail.Next = end
	return hair.Next
}

func reverse(head *ListNode, end *ListNode) (*ListNode, *ListNode) {
	var pre *ListNode
	node := head
	for node != end {
		next := node.Next
		node.Next = pre
		pre = node
		node = next
	}

	return pre, head
}

func main() {

}
