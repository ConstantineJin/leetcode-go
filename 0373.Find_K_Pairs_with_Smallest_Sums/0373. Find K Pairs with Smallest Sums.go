package main

import "container/heap"

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	m, n := len(nums1), len(nums2)
	ans := make([][]int, 0, min(k, m*n))
	h := make(hp, min(k, m))
	for i := range h {
		h[i] = pair{nums1[i] + nums2[0], i, 0}
	}
	for len(h) > 0 && len(ans) < k {
		p := heap.Pop(&h).(pair)
		i, j := p.i, p.j
		ans = append(ans, []int{nums1[i], nums2[j]})
		if j+1 < n {
			heap.Push(&h, pair{nums1[i] + nums2[j+1], i, j + 1})
		}
	}
	return ans
}

type pair struct{ s, i, j int }
type hp []pair

func (h *hp) Len() int           { return len(*h) }
func (h *hp) Less(i, j int) bool { return (*h)[i].s < (*h)[j].s }
func (h *hp) Swap(i, j int)      { (*h)[i], (*h)[j] = (*h)[j], (*h)[i] }
func (h *hp) Push(v any)         { *h = append(*h, v.(pair)) }
func (h *hp) Pop() any           { defer func() { *h = (*h)[:len(*h)-1] }(); return (*h)[len(*h)-1] }
