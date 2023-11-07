package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum, carry := &ListNode{}, 0 // sum指向当前操作的结果链表节点，carry表示进位值
	var head = sum
	for l1 != nil || l2 != nil || carry != 0 { // l1, l2和carry只要有一个不为空/不为0就继续操作
		if l1 == nil && l2 == nil { // l1和l2都遍历完了
			if carry != 0 { // 如果进位值不等于0
				sum.Next = &ListNode{Val: carry} // 添加最后一个节点，值为仅剩的进位值
			}
			return head.Next // 返回
		} else { // l1和l2至少有一个没有遍历完
			sum.Next = &ListNode{}
			sum = sum.Next // 添加结果的下一个节点并使sum指向之
			if l1 == nil { // 如果此时l1已遍历完
				carry += l2.Val // 进位值加上当前l2节点的值，对10取余是当前结果节点的值，除以10是下一个进位值
				sum.Val, carry = carry%10, carry/10
				l2 = l2.Next // l2指向下一个节点
			} else if l2 == nil { // 如果此时l2已遍历完，操作同上
				carry += l1.Val
				sum.Val, carry = carry%10, carry/10
				l1 = l1.Next
			} else { // l1和l2都没有遍历完，操作同上
				carry += l1.Val + l2.Val
				sum.Val, carry = carry%10, carry/10
				l1, l2 = l1.Next, l2.Next
			}
		}
	}
	return head.Next // 返回头节点的Next
}
