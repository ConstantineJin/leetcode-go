package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func doubleIt(head *ListNode) *ListNode {
	head = &ListNode{Val: 0, Next: head}
	for p := head; p != nil; p = p.Next {
		p.Val = p.Val * 2 % 10
		if p.Next != nil && p.Next.Val > 4 {
			p.Val++
		}
	}
	if head.Val == 0 {
		return head.Next
	}
	return head
}
