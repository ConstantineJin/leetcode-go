package main

import (
	"container/heap"
	"slices"
)

type event struct{ start, end, value int }

func maxTwoEvents(events [][]int) (ans int) {
	h := hp{}                                                            // 小顶堆
	var maxVal int                                                       // 结束时间在当前活动开始时间以前的活动的价值最大值
	slices.SortFunc(events, func(a, b []int) int { return a[0] - b[0] }) // 按活动开始时间排序
	for _, e := range events {
		start, end, value := e[0], e[1], e[2]
		for len(h) > 0 && h[0].end < start { // 堆顶元素的结束时间严格小于当前活动的开始时间
			maxVal = max(maxVal, heap.Pop(&h).(event).value) // 更新结束时间在当前活动开始时间以前的活动的价值最大值
		}
		ans = max(ans, maxVal+value)            // 更新答案
		heap.Push(&h, event{start, end, value}) // 当前活动入堆
	}
	return
}

type hp []event

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].end < h[j].end }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(event)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
