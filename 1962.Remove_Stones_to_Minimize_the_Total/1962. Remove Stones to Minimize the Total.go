package main

import (
	"container/heap"
	"sort"
)

func minStoneSum(piles []int, k int) (ans int) {
	h := &hp{piles}
	heap.Init(h) // 原地堆化
	for ; k > 0 && piles[0] > 0; k-- {
		piles[0] -= piles[0] / 2 // 直接修改堆顶
		heap.Fix(h, 0)
	}
	for _, x := range piles {
		ans += x
	}
	return
}

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] } // 最大堆
func (hp) Pop() (_ any)         { return }
func (hp) Push(any)             {}
