package main

import (
	"container/heap"
	"math"
)

type pair struct{ group, index, value int }

func smallestRange(nums [][]int) []int {
	ansLeft, ansRight := math.MaxInt, math.MinInt
	h := &hp{} // 小顶堆。保证在任意时刻每个数组有且只有一个元素在堆中
	for i, num := range nums {
		ansLeft, ansRight = min(ansLeft, num[0]), max(ansRight, num[0])
		heap.Push(h, pair{i, 0, num[0]}) // 初始将每个数组的第一个元素入堆
	}
	mx := ansRight // 堆中元素的最大值
	heap.Init(h)
	for {
		top := heap.Pop(h).(pair) // 最小元素出堆
		idx := top.index + 1
		if idx == len(nums[top.group]) { // 有数组已经遍历完了，答案区间的跨度不可能再发生变动，退出循环
			return []int{ansLeft, ansRight}
		}
		val := nums[top.group][idx]
		heap.Push(h, pair{top.group, idx, val}) // 出堆元素所在的数组的下一个元素入堆
		mx = max(mx, val)
		if mx-h.Top() < ansRight-ansLeft {
			ansLeft, ansRight = h.Top(), mx
		}
	}
}

type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].value < h[j].value }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h hp) Top() int           { return h[0].value }
