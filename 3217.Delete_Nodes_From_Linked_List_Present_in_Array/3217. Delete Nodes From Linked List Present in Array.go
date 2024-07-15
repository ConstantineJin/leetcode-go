package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func modifiedList(nums []int, head *ListNode) *ListNode {
	mp := make(map[int]bool)
	for _, num := range nums {
		mp[num] = true
	}
	dummy := &ListNode{Next: head}
	prev := dummy
	for p := head; p != nil; p = p.Next {
		if !mp[p.Val] {
			prev.Next = p
			prev = prev.Next
		}
	}
	prev.Next = nil
	return dummy.Next
}
