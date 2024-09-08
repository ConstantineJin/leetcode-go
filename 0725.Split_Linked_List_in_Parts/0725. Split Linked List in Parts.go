package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func splitListToParts(head *ListNode, k int) []*ListNode {
	if head == nil { // 如果原链表为空，直接返回
		return make([]*ListNode, k)
	}
	n := 1 // 链表的节点数
	for p := head; p.Next != nil; p = p.Next {
		n++
	}
	ans := make([]*ListNode, k)
	m, r := n/k, n%k // 平均每个部分至少有 m 个节点，其中前 r 个部分需要有 m+1 个节点
	p := head
	for i := 0; i < r; i++ {
		ans[i] = p
		for j := 0; j < m; j++ {
			p = p.Next
		}
		if p.Next == nil {
			return ans
		}
		// 断开链表
		q := p.Next
		p.Next = nil
		p = q
	}
	for i := r; i < k; i++ {
		ans[i] = p
		for j := 0; j < m-1; j++ {
			p = p.Next
		}
		if p.Next == nil {
			return ans
		}
		// 断开链表
		q := p.Next
		p.Next = nil
		p = q
	}
	return ans
}
