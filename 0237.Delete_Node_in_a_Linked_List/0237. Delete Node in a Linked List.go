package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val, node.Next = node.Next.Val, node.Next.Next // node不是最后一个节点，则node.Next不为空
}
