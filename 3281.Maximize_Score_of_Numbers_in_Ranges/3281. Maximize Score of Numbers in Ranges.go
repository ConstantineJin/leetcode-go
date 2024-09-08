package main

import (
	"math"
	"sort"
)

func maxPossibleScore(start []int, d int) int {
	sort.Ints(start)
	n := len(start)
	return sort.Search((start[n-1]+d-start[0])/(n-1), func(score int) bool { // 二分上界为平均值+1
		score++
		x := math.MinInt
		for _, s := range start {
			x += score
			if x > s+d {
				return true
			}
			x = max(x, s)
		}
		return false
	})
}
