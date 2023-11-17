package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head} // 前哨节点
	p := dummy
	for p.Next != nil && p.Next.Next != nil {
		p1, p2 := p.Next, p.Next.Next
		p.Next, p1.Next, p2.Next, p = p2, p2.Next, p1, p1
	}
	return dummy.Next
}
