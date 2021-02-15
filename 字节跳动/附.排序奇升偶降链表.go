package main

import "fmt"

/**
给定一个奇数位升序，偶数位降序的链表，将其重新排序。

输入: 1->8->3->6->5->4->7->2->NULL
输出: 1->2->3->4->5->6->7->8->NULL
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 思路：拆分链表，反转链表，合并链表

func sortOddEvenList(head *ListNode) *ListNode {
	odd, even := split(head)
	even = reversEven(even)
	return merge2List(odd, even)
}

func split(head *ListNode) (*ListNode, *ListNode) {
	oddNode, evenNode := head, head.Next
	oddPre := &ListNode{
		Next: oddNode,
	}
	evenPre := &ListNode{
		Next: evenNode,
	}

	for oddNode != nil && evenNode != nil {

		oddNode.Next = evenNode.Next
		oddNode = oddNode.Next
		if oddNode != nil {
			evenNode.Next = oddNode.Next
			evenNode = evenNode.Next
		}
	}

	return oddPre.Next, evenPre.Next
}

func reversEven(head *ListNode) *ListNode {

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

func merge2List(headA *ListNode, headB *ListNode) *ListNode {
	pre := new(ListNode)
	cur := pre
	for headA != nil && headB != nil {
		if headA.Val < headB.Val {
			cur.Next = headA
			cur = cur.Next
			headA = headA.Next
		} else {
			cur.Next = headB
			cur = cur.Next
			headB = headB.Next
		}
	}

	if headA != nil {
		cur.Next = headA
	}
	if headB != nil {
		cur.Next = headB
	}

	return pre.Next
}

func main() {
	pre := new(ListNode)
	cur := pre
	for i, j := 0, 3; i < 4 && j >= 0; {
		cur.Next = &ListNode{
			Val: 2*i + 1,
		}
		cur = cur.Next
		cur.Next = &ListNode{
			Val: 2*j + 2,
		}
		cur = cur.Next
		i++
		j--
	}
	cur = pre.Next
	for cur != nil {
		fmt.Print(cur.Val, "->")
		cur = cur.Next
	}
	fmt.Println()
	cur = sortOddEvenList(pre.Next)
	for cur != nil {
		fmt.Print(cur.Val, "->")
		cur = cur.Next
	}
}
