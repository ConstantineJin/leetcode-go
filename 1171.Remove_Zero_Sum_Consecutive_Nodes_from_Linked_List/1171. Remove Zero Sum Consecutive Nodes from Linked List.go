package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeZeroSumSublists(head *ListNode) *ListNode {
	var dummy, last = &ListNode{Next: head}, make(map[int]*ListNode)
	var p = dummy
	var s int
	for p != nil {
		s += p.Val
		last[s] = p
		p = p.Next
	}
	s, p = 0, dummy
	for p != nil {
		s += p.Val
		p.Next = last[s].Next
		p = p.Next
	}
	return dummy.Next
}
