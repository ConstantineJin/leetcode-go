package main

import "sort"

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findClosestElements(arr []int, k int, x int) []int {
	var n, idx = len(arr), sort.SearchInts(arr, x)
	var ans = max(0, idx-k)
	for i := max(0, idx-k); i < min(n-k, idx); i++ {
		if abs(arr[i+k]-x) < abs(arr[i]-x) {
			ans = i + 1
		}
	}
	return arr[ans : ans+k]
}
