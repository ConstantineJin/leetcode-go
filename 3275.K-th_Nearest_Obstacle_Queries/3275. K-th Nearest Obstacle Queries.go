package main

import (
	"container/heap"
	"sort"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func resultsArray(queries [][]int, k int) []int {
	ans := make([]int, 0, len(queries))
	h := hp{}
	for _, query := range queries {
		heap.Push(&h, abs(query[0])+abs(query[1]))
		if h.Len() > k {
			heap.Pop(&h)
		}
		if h.Len() < k {
			ans = append(ans, -1)
		} else {
			ans = append(ans, h.IntSlice[0])
		}
	}
	return ans
}

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any          { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
