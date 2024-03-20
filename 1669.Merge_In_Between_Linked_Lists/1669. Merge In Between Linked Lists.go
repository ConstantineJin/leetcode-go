package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	var i int
	var p, q *ListNode
	for q = list2; q.Next != nil; q = q.Next {
	}
	p = list1
	for ; i < b+1; i++ {
		p = p.Next
	}
	q.Next = p
	p = list1
	for i = 0; i < a-1; i++ {
		p = p.Next
	}
	p.Next = list2
	return list1
}
