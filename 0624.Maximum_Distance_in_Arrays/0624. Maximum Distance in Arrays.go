package main

import "math"

func maxDistance(arrays [][]int) (ans int) {
	mn, mx := math.MaxInt>>1, math.MinInt>>1
	for _, array := range arrays {
		first, last := array[0], array[len(array)-1]
		ans = max(ans, max(last-mn, mx-first))
		mn, mx = min(mn, first), max(mx, last)
	}
	return
}
