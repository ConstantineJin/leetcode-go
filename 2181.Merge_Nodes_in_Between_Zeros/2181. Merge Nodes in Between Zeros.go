package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	p := new(ListNode)
	dummy := p
	for q := head; q.Next != nil; q = q.Next {
		if q.Val == 0 {
			newNode := new(ListNode)
			p.Next = newNode
			p = p.Next
		} else {
			p.Val += q.Val
		}
	}
	return dummy.Next
}
