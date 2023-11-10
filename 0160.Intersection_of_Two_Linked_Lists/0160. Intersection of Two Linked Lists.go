package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil { // 任一链表为空则两者不可能相交
		return nil
	}
	pa, pb := headA, headB // 双指针，初始指向两链表头
	for pa != pb {         // 如果两链表等长，则第一轮双指针就会在交点处交会；如果不等长，则会在第二轮交会
		if pa == nil { // pa遍历完A链表，转指向B链表头
			pa = headB
		} else { // 否则遍历下一个节点
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa // 返回交会节点
}
