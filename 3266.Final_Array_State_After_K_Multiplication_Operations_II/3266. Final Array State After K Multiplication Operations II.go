package main

import (
	"container/heap"
	"sort"
)

const mod int = 1e9 + 7

func getFinalState(nums []int, k, multiplier int) []int {
	if multiplier == 1 {
		return nums
	}
	n := len(nums)
	var mx int
	h := make(hp, n)
	for i, x := range nums {
		mx = max(mx, x)
		h[i] = pair{x, i}
	}
	heap.Init(&h)
	for ; k > 0 && h[0].x < mx; k-- { // 模拟，直到堆顶是最大值
		h[0].x *= multiplier
		heap.Fix(&h, 0)
	}
	sort.Slice(h, func(i, j int) bool { return less(h[i], h[j]) })
	for i, p := range h {
		e := k / n
		if i < k%n {
			e++ // 前 k%n 个元素还要多操作一次
		}
		nums[p.i] = p.x % mod * pow(multiplier, e) % mod
	}
	return nums
}

type pair struct{ x, i int }

func less(a, b pair) bool { return a.x < b.x || a.x == b.x && a.i < b.i }

type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return less(h[i], h[j]) }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (hp) Push(any)             {}
func (hp) Pop() (_ any)         { return }

func pow(x, n int) int {
	res := 1
	for ; n > 0; n >>= 1 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
