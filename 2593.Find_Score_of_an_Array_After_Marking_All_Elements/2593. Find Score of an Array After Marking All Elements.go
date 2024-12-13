package main

import "container/heap"

type pair struct{ index, value int }

func findScore(nums []int) (ans int64) {
	n := len(nums)
	marked := make([]bool, n)
	h := &hp{}
	for i, num := range nums {
		heap.Push(h, pair{i, num})
	}
	for h.Len() > 0 {
		top := heap.Pop(h).(pair)
		if !marked[top.index] {
			ans += int64(top.value)
			if top.index > 0 {
				marked[top.index-1] = true
			}
			if top.index < n-1 {
				marked[top.index+1] = true
			}
		}
	}
	return
}

type hp []pair

func (h hp) Len() int { return len(h) }
func (h hp) Less(i, j int) bool {
	return h[i].value < h[j].value || h[i].value == h[j].value && h[i].index < h[j].index
}
func (h hp) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)   { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any     { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
