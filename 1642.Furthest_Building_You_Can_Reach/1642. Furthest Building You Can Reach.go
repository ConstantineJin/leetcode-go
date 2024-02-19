package main

import (
	"container/heap"
	"sort"
)

type hp struct{ sort.IntSlice }

func (h *hp) Push(x any) {
	h.IntSlice = append(h.IntSlice, x.(int))
}

func (h *hp) Pop() any {
	e := h.IntSlice[h.Len()-1]
	h.IntSlice = h.IntSlice[:h.Len()-1]
	return e
}

func furthestBuilding(heights []int, bricks int, ladders int) int {
	// 小顶堆，存经过的差值
	// 假设每一次都是用万能梯子前进，每一次将差值入堆，要想真的从这个差值过，从堆里pop出前面最小的差值，blocks里必须满足这个差值，否则跳出
	h := &hp{}
	heap.Init(h)
	for i := 1; i < len(heights); i++ {
		diff := heights[i] - heights[i-1]
		if diff > 0 {
			// 经过的有差值的都记录在堆
			heap.Push(h, diff)
			ladders--
			if ladders < 0 {
				e := heap.Pop(h).(int)
				bricks -= e
				if bricks < 0 {
					return i - 1
				}
			}
		}
	}
	return len(heights) - 1
}
