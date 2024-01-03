package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归
//func removeNodes(head *ListNode) *ListNode {
//	if head == nil {
//		return nil
//	}
//	head.Next = removeNodes(head.Next)
//	if head.Next != nil && head.Val < head.Next.Val {
//		return head.Next
//	} else {
//		return head
//	}
//}

// 反转链表
func reverse(head *ListNode) *ListNode {
	var dummy = &ListNode{}
	for head != nil {
		var p = head
		head = head.Next
		p.Next = dummy.Next
		dummy.Next = p
	}
	return dummy.Next
}

func removeNodes(head *ListNode) *ListNode {
	head = reverse(head)
	for p := head; p.Next != nil; {
		if p.Val > p.Next.Val {
			p.Next = p.Next.Next
		} else {
			p = p.Next
		}
	}
	return reverse(head)
}
