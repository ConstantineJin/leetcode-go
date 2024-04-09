package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var dummy = &ListNode{Next: head}
	var p = dummy
	for i := 1; i < left; i++ {
		p = p.Next
	}
	var pre, cur *ListNode = nil, p.Next
	for i := left; i <= right; i++ {
		var nxt = cur.Next
		cur.Next = pre
		pre, cur = cur, nxt
	}
	p.Next.Next, p.Next = cur, pre // 反转完成后，按照原链表顺序看，pre指向反转部分的最后一个节点，cur指向反转部分后的第一个节点，pre.Next指向反转的第一个节点
	return dummy.Next
}
