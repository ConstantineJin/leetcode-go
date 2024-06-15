package main

import (
	"container/heap"
	"sort"
)

type project struct{ profit, capital int }

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	n := len(profits)
	projects := make([]project, n)
	for i := range projects {
		projects[i] = project{profit: profits[i], capital: capital[i]}
	}
	sort.Slice(projects, func(i, j int) bool { return projects[i].capital < projects[j].capital }) // 将所有项目按照所需资本升序排序
	h := new(hp)                                                                                   // 建立利润的大根堆
	for i := 0; k > 0; k-- {
		for ; i < n && projects[i].capital <= w; i++ {
			heap.Push(h, projects[i].profit) // profit 是纯利润，因此将项目压入堆时不需要将 w 减去 capital
		}
		if h.Len() == 0 { // 没有足够的资本 w 进行下一轮投资
			break
		}
		w += heap.Pop(h).(int) // 贪心，每次从堆中取出最大值
	}
	return w
}

type hp sort.IntSlice

func (h *hp) Len() int           { return len(*h) }
func (h *hp) Less(i, j int) bool { return (*h)[i] > (*h)[j] }
func (h *hp) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *hp) Push(v any)         { *h = append(*h, v.(int)) }
func (h *hp) Pop() any           { defer func() { *h = (*h)[:len(*h)-1] }(); return (*h)[len(*h)-1] }
