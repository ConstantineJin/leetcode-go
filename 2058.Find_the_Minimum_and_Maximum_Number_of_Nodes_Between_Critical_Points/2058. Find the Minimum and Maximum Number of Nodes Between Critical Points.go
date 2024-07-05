package main

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func nodesBetweenCriticalPoints(head *ListNode) []int {
	var i, j, prev, first, last int // i 表示当前节点的编号（头节点对应 i = 0），j 表示上一个关键节点的编号，prev 记录上一个节点的值，first 和 last 维护第一个和最后一个关键节点的编号
	minDis := math.MaxInt           // 两个关键节点之间的最小距离
	for p := head; p.Next != nil; p = p.Next {
		if prev > 0 && p.Next != nil {
			v, next := p.Val, p.Next.Val
			if v > prev && v > next || v < prev && v < next {
				if first == 0 {
					first = i
				}
				last = i
				if j > 0 {
					minDis = min(minDis, i-j)
				}
				j = i
			}
		}
		prev = p.Val
		i++
	}
	if first == last {
		return []int{-1, -1}
	}
	return []int{minDis, last - first}
}
