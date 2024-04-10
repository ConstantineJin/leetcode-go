package main

import "math"

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func findGCD(nums []int) int {
	var mn, mx = math.MaxInt, math.MinInt
	for _, num := range nums {
		mn, mx = min(mn, num), max(mx, num)
	}
	return gcd(mn, mx)
}
