package main

import (
	"slices"
	"sort"
)

func maximumBeauty(items [][]int, queries []int) []int {
	slices.SortFunc(items, func(a, b []int) int { return a[0] - b[0] }) // 按价格升序排序
	n := len(items)
	for i := 1; i < n; i++ {
		items[i][1] = max(items[i][1], items[i-1][1])
	}
	for i, query := range queries {
		if idx := sort.Search(n, func(i int) bool { return items[i][0] > query }) - 1; idx >= 0 {
			queries[i] = items[idx][1]
		} else {
			queries[i] = 0
		}
	}
	return queries
}
