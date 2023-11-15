package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	fast, slow := head.Next.Next, head.Next // 快慢指针
	for fast != nil {
		if fast == slow { // 快慢指针相遇，说明链表存在环
			return true
		}
		if fast.Next == nil { // 链表存在终点，说明没有环
			return false
		}
		fast = fast.Next.Next // 快指针每次前进两步
		slow = slow.Next      // 慢指针每次前进一步
	}
	return false // 走到终点，说明没有环
}
