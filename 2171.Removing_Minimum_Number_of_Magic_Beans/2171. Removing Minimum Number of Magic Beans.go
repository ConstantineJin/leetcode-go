package main

import "sort"

func minimumRemoval(beans []int) int64 {
	sort.Ints(beans)
	var sum int
	for _, bean := range beans {
		sum += bean
	}
	var n, ans = len(beans), sum
	for i, bean := range beans {
		ans = min(ans, sum-(n-i)*bean)
	}
	return int64(ans)
}
