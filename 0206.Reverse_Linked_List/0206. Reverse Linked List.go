package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归
//func reverseList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//	newHead := reverseList(head.Next)
//	head.Next, head.Next.Next = nil, head
//	return newHead
//}

// 迭代
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
