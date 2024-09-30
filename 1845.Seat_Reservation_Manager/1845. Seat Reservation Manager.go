package main

import (
	"container/heap"
	"sort"
)

type SeatManager struct{ sort.IntSlice } // 最小堆

func Constructor(n int) SeatManager {
	m := SeatManager{make(sort.IntSlice, n)}
	for i := range m.IntSlice {
		m.IntSlice[i] = i + 1
	}
	return m // 本身就是有序的，无需调用 heap.Init() 方法
}

func (m *SeatManager) Reserve() int { return heap.Pop(m).(int) }

func (m *SeatManager) Unreserve(seatNumber int) { heap.Push(m, seatNumber) }

func (m *SeatManager) Push(v any) { m.IntSlice = append(m.IntSlice, v.(int)) }
func (m *SeatManager) Pop() any {
	a := m.IntSlice
	v := a[len(a)-1]
	m.IntSlice = a[:len(a)-1]
	return v
}
