package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil { // 空节点直接返回nil
		return nil
	} else if head.Next == nil { // 只有一个节点就直接返回
		return &TreeNode{Val: head.Val}
	}
	prev, slow, fast := &ListNode{}, head, head // slow与fast双指针找到链表的中点，prev用于断开链表
	for fast != nil && fast.Next != nil {
		prev, slow, fast = slow, slow.Next, fast.Next.Next
	}
	if prev != nil {
		prev.Next = nil // 在slow前断开链表
	}
	return &TreeNode{Val: slow.Val, Left: sortedListToBST(head), Right: sortedListToBST(slow.Next)} // 递归构建左子树和右子树
}
