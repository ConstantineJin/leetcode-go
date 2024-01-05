package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var p = head
	for p.Next != nil {
		if p.Next.Val == p.Val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return head
}
