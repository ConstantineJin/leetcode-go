package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	p, q := new(ListNode), new(ListNode)
	pHead, qHead := p, q
	for head != nil {
		if head.Val < x {
			p.Next = head
			p = p.Next
		} else {
			q.Next = head
			q = q.Next
		}
		head = head.Next
	}
	p.Next, q.Next = qHead.Next, nil
	return pHead.Next
}
