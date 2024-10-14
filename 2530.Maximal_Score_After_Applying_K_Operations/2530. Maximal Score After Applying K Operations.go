package main

import "container/heap"

func maxKelements(nums []int, k int) (ans int64) {
	h := (*hp)(&nums)
	heap.Init(h)
	for range k {
		x := heap.Pop(h).(int)
		ans += int64(x)
		heap.Push(h, (x+2)/3)
	}
	return ans
}

type hp []int

func (h hp) Len() int {
	return len(h)
}
func (h hp) Less(i, j int) bool {
	return h[i] > h[j]
}
func (h hp) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(x any)   { *h = append(*h, x.(int)) }
func (h *hp) Pop() any     { n := len(*h); x := (*h)[n-1]; *h = (*h)[:n-1]; return x }
