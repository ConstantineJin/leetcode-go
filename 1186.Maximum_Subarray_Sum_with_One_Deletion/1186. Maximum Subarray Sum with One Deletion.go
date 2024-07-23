package main

import "math"

func maximumSum(arr []int) int {
	ans := math.MinInt / 2
	f0, f1 := ans, ans
	for _, x := range arr {
		f1 = max(f1+x, f0)
		f0 = max(f0, 0) + x
		ans = max(ans, f0, f1)
	}
	return ans
}
