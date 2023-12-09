package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	var reverse = func(head *ListNode) *ListNode {
		var prev *ListNode
		curr := head
		for i := 0; i < k; i++ {
			next := curr.Next
			curr.Next = prev
			prev = curr
			curr = next
		}
		return prev
	}

	return dummy.Next
}
