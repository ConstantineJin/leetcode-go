package main

import (
	"container/heap"
	"math"
	"sort"
)

func pickGifts(gifts []int, k int) (ans int64) {
	h := &hp{gifts}
	heap.Init(h) // 原地堆化
	for ; k > 0 && gifts[0] > 1; k-- {
		gifts[0] = int(math.Sqrt(float64(gifts[0]))) // 直接修改堆顶
		heap.Fix(h, 0)
	}
	for _, x := range gifts {
		ans += int64(x)
	}
	return
}

type hp struct{ sort.IntSlice }

func (h hp) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] } // 最大堆
func (hp) Pop() (_ any)         { return }
func (hp) Push(any)             {}
