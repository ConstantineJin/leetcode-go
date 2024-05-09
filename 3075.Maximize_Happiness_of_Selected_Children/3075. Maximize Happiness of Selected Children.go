package main

import "sort"

func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Ints(happiness)
	n := len(happiness)
	var ans int
	for i := n - 1; i >= n-k; i-- {
		ans += max(happiness[i]-n+i+1, 0)
	}
	return int64(ans)
}
