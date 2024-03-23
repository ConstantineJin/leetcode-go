package main

type ListNode struct {
	Val  int
	Next *ListNode
}

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

func reorderList(head *ListNode) {
	var fast, slow = head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	var p, q = head, reverseList(slow.Next)
	slow.Next = nil
	for p != nil && q != nil {
		pp, qq := p.Next, q.Next
		p.Next, q.Next = q, pp
		p, q = pp, qq
	}
}
