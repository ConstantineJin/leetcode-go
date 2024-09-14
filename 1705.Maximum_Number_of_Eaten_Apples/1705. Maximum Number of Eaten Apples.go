package main

import "container/heap"

type Apple struct{ count, rotten int }

func eatenApples(apples, days []int) (ans int) {
	h := hp{}
	for i, cnt := range apples {
		for h.Len() > 0 && h[0].rotten <= i {
			heap.Pop(&h)
		}
		if cnt > 0 {
			heap.Push(&h, Apple{count: cnt, rotten: i + days[i]})
		}
		if h.Len() > 0 {
			h[0].count--
			if h[0].count == 0 {
				heap.Pop(&h)
			}
			ans++
		}
	}
	for i := len(apples); h.Len() > 0; i++ {
		for h.Len() > 0 && h[0].rotten <= i {
			heap.Pop(&h)
		}
		if h.Len() == 0 {
			break
		}
		h[0].count--
		if h[0].count == 0 {
			heap.Pop(&h)
		}
		ans++
	}
	return
}

type hp []Apple

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].rotten < h[j].rotten }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(Apple)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
