package main

import "math"

func maxProductDifference(nums []int) int {
	var a, b, c, d = 0, 0, math.MaxInt, math.MaxInt
	for _, num := range nums {
		a, b, c, d = max(a, num), max(b, min(a, num)), min(d, num), min(c, max(d, num))
	}
	return a*b - c*d
}
