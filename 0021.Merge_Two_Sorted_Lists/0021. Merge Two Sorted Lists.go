package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 递归
//func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
//	if list1 == nil {
//		return list2
//	} else if list2 == nil {
//		return list1
//	}
//	if list1.Val < list2.Val {
//		list1.Next = mergeTwoLists(list1.Next, list2)
//		return list1
//	} else {
//		list2.Next = mergeTwoLists(list1, list2.Next)
//		return list2
//	}
//}

// 迭代
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	p := &ListNode{}
	head := p
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			p.Next = list1
			list1 = list1.Next
		} else {
			p.Next = list2
			list2 = list2.Next
		}
		p = p.Next
	}
	if list1 != nil {
		p.Next = list1
	} else {
		p.Next = list2
	}
	return head.Next
}
