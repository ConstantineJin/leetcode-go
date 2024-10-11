package main

import (
	"container/heap"
	"slices"
	"sort"
)

type friend struct{ index, arrival, leaving int }

func smallestChair(times [][]int, targetFriend int) int {
	n := len(times)
	friends := make([]friend, n)
	for i, time := range times {
		friends[i] = friend{i, time[0], time[1]}
	}
	slices.SortFunc(friends, func(a, b friend) int { return a.arrival - b.arrival })
	emptyChairs, occupiedChairs := &EmptyHeap{}, &OccupiedHeap{}
	for i := 0; i < n; i++ {
		heap.Push(emptyChairs, i)
	}
	for _, f := range friends {
		for occupiedChairs.Len() > 0 && occupiedChairs.Top().leaving <= f.arrival { // 释放已经离开的朋友的椅子
			heap.Push(emptyChairs, heap.Pop(occupiedChairs).(OccupiedFriend).chair)
		}
		chair := heap.Pop(emptyChairs).(int) // 为当前朋友分配椅子
		if f.index == targetFriend {
			return chair
		}
		heap.Push(occupiedChairs, OccupiedFriend{f.leaving, chair}) // 将当前朋友的离开时间和椅子加入占用队列
	}
	return -1
}

type OccupiedFriend struct{ leaving, chair int }

type OccupiedHeap []OccupiedFriend

func (h OccupiedHeap) Len() int            { return len(h) }
func (h OccupiedHeap) Less(i, j int) bool  { return h[i].leaving < h[j].leaving }
func (h OccupiedHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *OccupiedHeap) Push(x any)         { *h = append(*h, x.(OccupiedFriend)) }
func (h *OccupiedHeap) Pop() any           { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h OccupiedHeap) Top() OccupiedFriend { return h[0] }

type EmptyHeap sort.IntSlice

func (h EmptyHeap) Len() int           { return len(h) }
func (h EmptyHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h EmptyHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *EmptyHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *EmptyHeap) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
