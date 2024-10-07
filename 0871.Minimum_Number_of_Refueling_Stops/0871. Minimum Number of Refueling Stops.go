package main

import (
	"container/heap"
	"sort"
)

func minRefuelStops(target, startFuel int, stations [][]int) (ans int) {
	stations = append(stations, []int{target, 0}) // 将终点也视作一个加油站
	prePosition, curFuel := 0, startFuel
	h := &hp{}
	for _, station := range stations {
		position, fuel := station[0], station[1]
		curFuel -= position - prePosition
		for curFuel < 0 && h.Len() > 0 {
			curFuel += heap.Pop(h).(int)
			ans++
		}
		if curFuel < 0 {
			return -1
		}
		heap.Push(h, fuel)
		prePosition = position
	}
	return
}

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] } // 使用大顶堆维护经过的加油站的油量
func (h *hp) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() any          { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
