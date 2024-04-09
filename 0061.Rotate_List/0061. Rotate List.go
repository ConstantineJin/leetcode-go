package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) (ans *ListNode) {
	if k == 0 || head == nil || head.Next == nil {
		return head
	}
	var n = 1
	for p := head; p.Next != nil; p = p.Next {
		n++
	}
	if k%n == 0 {
		return head
	}
	k = n - k%n
	var p, q *ListNode = head, nil
	for i := 1; i < k; i++ {
		p = p.Next
	}
	ans = p.Next
	for q = p; q.Next != nil; q = q.Next {
	}
	p.Next, q.Next = nil, head
	return
}
