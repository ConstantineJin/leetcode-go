package main

import (
	"container/heap"
	"sort"
)

func magicTower(nums []int) (ans int) {
	var sum int
	for _, num := range nums {
		sum += num
	}
	if sum < 0 {
		return -1
	}
	var hp, h = 1, &minHeap{}
	for _, num := range nums {
		if num < 0 {
			heap.Push(h, num)
		}
		hp += num
		if hp < 1 {
			// 这意味着 num < 0，所以前面必然会把 num 入堆
			// 所以堆必然不是空的，并且堆顶 <= num
			hp -= heap.Pop(h).(int) // 反悔
			ans++
		}
	}
	return
}

type minHeap struct{ sort.IntSlice } // 继承 Len, Less, Swap
func (h *minHeap) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *minHeap) Pop() any          { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
