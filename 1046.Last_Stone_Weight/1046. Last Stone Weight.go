package main

import (
	"container/heap"
	"sort"
)

func lastStoneWeight(stones []int) int {
	var h = &hp{stones}
	heap.Init(h)
	for h.Len() > 1 {
		var x, y = heap.Pop(h).(int), heap.Pop(h).(int)
		if x != y {
			heap.Push(h, x-y)
		}
	}
	if h.Len() > 0 {
		return h.IntSlice[0]
	}
	return 0
}

type hp struct{ sort.IntSlice }

func (h *hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(x any)         { h.IntSlice = append(h.IntSlice, x.(int)) }
func (h *hp) Pop() any {
	defer func() { h.IntSlice = h.IntSlice[:len(h.IntSlice)-1] }()
	return h.IntSlice[len(h.IntSlice)-1]
}
