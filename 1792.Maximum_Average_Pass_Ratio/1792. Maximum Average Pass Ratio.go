package main

import "container/heap"

func maxAverageRatio(classes [][]int, extraStudents int) (ans float64) {
	h := hp(classes)
	heap.Init(&h)
	for ; extraStudents > 0; extraStudents-- {
		h[0][0]++
		h[0][1]++
		heap.Fix(&h, 0)
	}
	for _, c := range h {
		ans += float64(c[0]) / float64(c[1])
	}
	return ans / float64(len(classes))
}

type hp [][]int

func (h hp) Len() int { return len(h) }
func (h hp) Less(i, j int) bool {
	return (h[i][1]-h[i][0])*h[j][1]*(h[j][1]+1) > (h[j][1]-h[j][0])*h[i][1]*(h[i][1]+1)
}
func (h hp) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
func (hp) Push(_ any)      {}
func (hp) Pop() (_ any)    { return }
