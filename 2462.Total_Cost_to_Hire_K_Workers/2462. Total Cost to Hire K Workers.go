package main

import (
	"container/heap"
	"sort"
)

func totalCost(costs []int, k int, candidates int) (ans int64) {
	var n = len(costs)
	// 前后双指针能够相遇，即所有工人都可以被选择，直接选取最小的k个工人
	if candidates*2+k > n {
		sort.Ints(costs)
		for i := 0; i < k; i++ {
			ans += int64(costs[i])
		}
		return
	}
	// 维护两个最小堆，分别存储前candidates个工人和后candidates个工人的工资
	var pre, suf = hp{costs[:candidates]}, hp{costs[n-candidates:]}
	heap.Init(&pre)
	heap.Init(&suf)
	for i, j := candidates, n-candidates-1; k > 0; k-- {
		// 选择两个最小堆中较小的堆顶，如果相等选择前一个堆顶
		if pre.IntSlice[0] <= suf.IntSlice[0] {
			ans += int64(pre.replace(costs[i]))
			i++
		} else {
			ans += int64(suf.replace(costs[j]))
			j--
		}
	}
	return
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(x any) { h.IntSlice = append(h.IntSlice, x.(int)) }
func (h *hp) Pop() any {
	defer func() { h.IntSlice = h.IntSlice[:len(h.IntSlice)-1] }()
	return h.IntSlice[len(h.IntSlice)-1]
}
func (h *hp) replace(v int) int { top := h.IntSlice[0]; h.IntSlice[0] = v; heap.Fix(h, 0); return top }
