package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 转化为数组，时间复杂度O(n)，空间复杂度O(n)
//func isPalindrome(head *ListNode) bool {
//	arr := make([]int, 0)
//	for head != nil {
//		arr = append(arr, head.Val)
//		head = head.Next
//	}
//	i, j := 0, len(arr)-1
//	for i < j {
//		if arr[i] != arr[j] {
//			return false
//		}
//		i++
//		j--
//	}
//	return true
//}

// 双指针，时间复杂度O(n)， 空间复杂度O(1)
// reverseList 反转链表
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

// endOfFirstHalf 使用快慢指针法求得前半段链表的结束节点
func endOfFirstHalf(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	firstHalfEnd := endOfFirstHalf(head)
	secondHalfStart := reverseList(firstHalfEnd.Next) // 反转后半段链表
	p1, p2, res := head, secondHalfStart, true        // 双指针指向前半段与后半段的起始位置
	for p1 != nil && p2 != nil {
		if p1.Val != p2.Val {
			res = false // 不直接return false是为了将后半段链表还原至初始状态，如果不需要还原可以直接返回
			break
		}
		p1, p2 = p1.Next, p2.Next
	}
	firstHalfEnd.Next = reverseList(secondHalfStart) // 还原后半段链表至原始状态，无此行也可通过
	return res
}
