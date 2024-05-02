package main

import (
	"container/heap"
	"math"
	"sort"
)

type worker struct{ quality, wage int }

func mincostToHireWorkers(quality []int, wage []int, k int) float64 {
	var n = len(quality)
	var workers = make([]worker, n)
	for i := range workers {
		workers[i] = worker{quality[i], wage[i]}
	}
	// 按wage/quality升序排序
	sort.Slice(workers, func(i, j int) bool { return workers[i].wage*workers[j].quality < workers[j].wage*workers[i].quality })
	var h = &hp{make([]int, k)}
	var sum int
	for i, wk := range workers[:k] {
		sum += wk.quality
		h.IntSlice[i] = wk.quality
	}
	heap.Init(h)                                                             // 建立大顶堆
	var ans = float64(sum*workers[k-1].wage) / float64(workers[k-1].quality) // 选取wage/quality最小的k个工人作为初始化解
	for _, wk := range workers[k:] {
		if wk.quality < h.IntSlice[0] { // 如果当前工人的quality小于堆顶的quality，则可以得到更小的sum，有可能得到更优答案
			sum -= h.IntSlice[0] - wk.quality
			h.IntSlice[0] = wk.quality
			heap.Fix(h, 0) // 更新堆顶
			ans = math.Min(ans, float64(sum*wk.wage)/float64(wk.quality))
		}
	}
	return ans
}

type hp struct{ sort.IntSlice }

func (h *hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
