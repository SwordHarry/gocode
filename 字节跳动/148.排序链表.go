package main

/**
给你链表的头结点head，请将其按 升序 排列并返回 排序后的链表 。
进阶：
你可以在O(nlogn) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？
示例 1：
输入：head = [4,2,1,3]
输出：[1,2,3,4]

示例 2：
输入：head = [-1,5,3,4,0]
输出：[-1,0,3,4,5]

示例 3：
输入：head = []
输出：[]

提示：
链表中节点的数目在范围[0, 5 * 104]内
-105<= Node.val <= 105
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 思路：归并排序，分链表，然后进行合并
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	// 分链表
	mid := slow
	next := mid.Next
	mid.Next = nil
	left := sortList(head)
	right := sortList(next)

	// 合并链表
	pre := &ListNode{}
	cur := pre
	for left != nil && right != nil {
		if left.Val < right.Val {
			cur.Next = left
			left = left.Next
		} else {
			cur.Next = right
			right = right.Next
		}
		cur = cur.Next
	}

	if left != nil {
		cur.Next = left
	}
	if right != nil {
		cur.Next = right
	}
	return pre.Next
}

func main() {

}
