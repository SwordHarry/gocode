package main

/**
请判断一个链表是否为回文链表。
示例 1:
输入: 1->2
输出: false

示例 2:
输入: 1->2->2->1
输出: true

进阶：
你能否用O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	node2 := reverseList(slow)
	node := head
	for node != nil && node2 != nil {
		if node.Val != node2.Val {
			return false
		}
		node = node.Next
		node2 = node2.Next
	}

	return true
}

// 反转链表
func reverseList(head *ListNode) *ListNode {
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

func main() {

}
