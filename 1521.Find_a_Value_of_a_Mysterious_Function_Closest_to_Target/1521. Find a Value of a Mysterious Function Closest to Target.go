package main

import "math"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func closestToTarget(arr []int, target int) int {
	ans := math.MaxInt
	for i, x := range arr {
		ans = min(ans, abs(x-target))
		for j := i - 1; j >= 0 && arr[j]|x != x; j-- {
			arr[j] &= x
			ans = min(ans, abs(arr[j]-target))
		}
	}
	return ans
}
