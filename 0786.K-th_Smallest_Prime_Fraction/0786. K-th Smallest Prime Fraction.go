package main

import "container/heap"

// 优先队列
func kthSmallestPrimeFraction(arr []int, k int) []int {
	n := len(arr)
	h := make(hp, n-1)
	for j := 1; j < n; j++ {
		h[j-1] = frac{x: arr[0], y: arr[j], i: 0, j: j}
	}
	heap.Init(&h)
	for ; k > 1; k-- {
		f := heap.Pop(&h).(frac)
		if f.i+1 < f.j {
			heap.Push(&h, frac{x: arr[f.i+1], y: f.y, i: f.i + 1, j: f.j})
		}
	}
	return []int{h[0].x, h[0].y}
}

type frac struct{ x, y, i, j int }
type hp []frac

func (h *hp) Len() int           { return len(*h) }
func (h *hp) Less(i, j int) bool { return (*h)[i].x*(*h)[j].y < (*h)[i].y*(*h)[j].x }
func (h *hp) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *hp) Push(v any)         { *h = append(*h, v.(frac)) }
func (h *hp) Pop() any {
	defer func() { *h = (*h)[:len(*h)-1] }()
	return (*h)[len(*h)-1]
}
