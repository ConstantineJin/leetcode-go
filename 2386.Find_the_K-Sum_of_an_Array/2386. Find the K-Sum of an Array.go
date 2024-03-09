package main

import (
	"container/heap"
	"slices"
)

func kSum(nums []int, k int) int64 {
	var n = len(nums)
	var sum int
	for i, num := range nums {
		if num >= 0 {
			sum += num
		} else {
			nums[i] = -num
		}
	}
	slices.Sort(nums)
	h := hp{{0, 0}} // 空子序列
	for ; k > 1; k-- {
		p := heap.Pop(&h).(pair)
		i := p.i
		if i < n {
			// 在子序列的末尾添加 nums[i]
			heap.Push(&h, pair{p.sum + nums[i], i + 1}) // 下一个添加/替换的元素下标为 i+1
			if i > 0 {                                  // 替换子序列的末尾元素为 nums[i]
				heap.Push(&h, pair{p.sum + nums[i] - nums[i-1], i + 1})
			}
		}
	}
	return int64(sum - h[0].sum)
}

type pair struct{ sum, i int }
type hp []pair

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i].sum < h[j].sum }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v any)        { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
