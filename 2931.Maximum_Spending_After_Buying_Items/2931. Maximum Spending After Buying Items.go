package main

import "container/heap"

type pair struct{ shop, index, value int }

func maxSpending(values [][]int) int64 {
	m, n := len(values), len(values[0])
	h := &hp{}
	for i, value := range values {
		heap.Push(h, pair{i, n - 1, value[n-1]})
	}
	heap.Init(h)
	var ans int
	for d := 1; d <= m*n; d++ {
		top := heap.Pop(h).(pair)
		ans += top.value * d
		if top.index > 0 {
			heap.Push(h, pair{top.shop, top.index - 1, values[top.shop][top.index-1]})
		}
	}
	return int64(ans)
}

type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].value < h[j].value }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
