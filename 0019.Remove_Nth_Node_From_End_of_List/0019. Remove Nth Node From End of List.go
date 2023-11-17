package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head} // 创建前哨节点
	fast, slow := head, dummy   // 双指针
	for i := 0; i < n; i++ {    // 快指针先走n步
		fast = fast.Next
	}
	for ; fast != nil; fast, slow = fast.Next, slow.Next { // 快指针走到终点时慢指针恰好走到倒数第n+1个节点
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
