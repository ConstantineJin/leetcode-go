package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	var fast, slow = head, head
	for ; fast != nil && fast.Next != nil; fast, slow = fast.Next.Next, slow.Next {
	}
	return slow
}
