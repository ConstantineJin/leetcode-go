package main

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	var h = hp{}
	for _, head := range lists {
		if head != nil {
			h = append(h, head)
		}
	}
	heap.Init(&h)           // 堆化
	var dummy = &ListNode{} // 哨兵节点，作为合并后链表头节点的前一个节点
	var cur = dummy
	for len(h) > 0 { // 循环直到堆为空
		var node = heap.Pop(&h).(*ListNode) // 剩余节点中的最小节点
		if node.Next != nil {               // 下一个节点不为空
			heap.Push(&h, node.Next) // 下一个节点有可能是最小节点，入堆
		}
		cur.Next = node // 合并到新链表中
		cur = cur.Next  // 准备合并下一个节点
	}
	return dummy.Next // 哨兵节点的下一个节点就是新链表的头节点
}

type hp []*ListNode

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].Val < h[j].Val } // 最小堆
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(*ListNode)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
