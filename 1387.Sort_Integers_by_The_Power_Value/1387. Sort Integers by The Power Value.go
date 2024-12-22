package main

import (
	"cmp"
	"slices"
)

var cache = map[int]int{1: 0}

func dfs(x int) int {
	if v, ok := cache[x]; ok {
		return v
	}
	if x%2 == 0 {
		cache[x] = dfs(x/2) + 1
	} else {
		cache[x] = dfs(x*3+1) + 1
	}
	return cache[x]
}

func getKth(lo, hi, k int) int {
	arr := make([]int, hi-lo+1)
	for i := range arr {
		arr[i] = lo + i
	}
	slices.SortFunc(arr, func(a, b int) int { return cmp.Or(dfs(a)-dfs(b), a-b) })
	return arr[k-1]
}
