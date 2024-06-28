package main

import "sort"

func maximumImportance(n int, roads [][]int) int64 {
	deg := make([]int, n)
	for _, road := range roads {
		deg[road[0]]++
		deg[road[1]]++
	}
	sort.Slice(deg, func(i, j int) bool { return deg[i] < deg[j] }) // 按邻居数升序排序
	var ans int
	for i, d := range deg {
		ans += (i + 1) * d
	}
	return int64(ans)
}
