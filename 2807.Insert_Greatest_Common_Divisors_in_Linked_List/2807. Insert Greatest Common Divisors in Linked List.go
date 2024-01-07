package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	for p := head; p.Next != nil; p = p.Next.Next {
		p.Next = &ListNode{Val: gcd(p.Val, p.Next.Val), Next: p.Next}
	}
	return head
}
