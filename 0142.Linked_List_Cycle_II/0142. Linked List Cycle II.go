package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow { // 快慢指针相遇时
			p := head // p指针从头以1为步长前进
			for p != slow {
				p, slow = p.Next, slow.Next
			}
			return p // p与slow相遇即为链表入环处
		}
	}
	return nil
}
