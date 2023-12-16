package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	var pre, p = tail.Next, head
	for pre != tail {
		var nxt = p.Next
		p.Next = pre
		pre = p
		p = nxt
	}
	return tail, head
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var dummy = &ListNode{Next: head}
	var pre = dummy
	for head != nil {
		var tail = pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummy.Next
			}
		}
		var nxt = tail.Next
		head, tail = reverse(head, tail)
		pre.Next, tail.Next = head, nxt
		pre, head = tail, tail.Next
	}
	return dummy.Next
}
