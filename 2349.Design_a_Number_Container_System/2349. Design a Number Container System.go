package main

import (
	"container/heap"
	"sort"
)

type NumberContainers struct {
	m  map[int]int
	ms map[int]*hp
}

func Constructor() NumberContainers {
	return NumberContainers{map[int]int{}, map[int]*hp{}}
}

func (n NumberContainers) Change(index int, number int) {
	n.m[index] = number
	if n.ms[number] == nil {
		n.ms[number] = &hp{}
	}
	heap.Push(n.ms[number], index) // 直接添加新数据，后面 find 再删除旧的
}

func (n NumberContainers) Find(number int) int {
	h, ok := n.ms[number]
	if !ok {
		return -1
	}
	for h.Len() > 0 && n.m[h.IntSlice[0]] != number {
		heap.Pop(h)
	}
	if h.Len() == 0 {
		return -1
	}
	return h.IntSlice[0]
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
