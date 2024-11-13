package main

import "sort"

func countFairPairs(nums []int, lower, upper int) (ans int64) {
	sort.Ints(nums)
	for i, num := range nums {
		ans += int64(sort.SearchInts(nums[:i], upper-num+1) - sort.SearchInts(nums[:i], lower-num))
	}
	return ans
}
