package main

import "sort"

func countWays(ranges [][]int) int {
	sort.Slice(ranges, func(i, j int) bool { return ranges[i][0] < ranges[j][0] })
	var ans, maxR = 1, -1
	for _, p := range ranges {
		if p[0] > maxR {
			ans = ans * 2 % (1e9 + 7)
		}
		maxR = max(maxR, p[1])
	}
	return ans
}
