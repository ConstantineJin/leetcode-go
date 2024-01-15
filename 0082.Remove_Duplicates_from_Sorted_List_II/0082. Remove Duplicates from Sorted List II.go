package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	var dummy = &ListNode{Next: head}
	for p := dummy; p.Next != nil && p.Next.Next != nil; {
		if p.Next.Val == p.Next.Next.Val {
			var temp = p.Next.Val
			for p.Next != nil && p.Next.Val == temp {
				p.Next = p.Next.Next
			}
		} else {
			p = p.Next
		}
	}
	return dummy.Next
}
