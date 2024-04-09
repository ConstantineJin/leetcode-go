package main

import "sort"

func maximumCount(nums []int) int {
	return max(sort.SearchInts(nums, 0), len(nums)-sort.SearchInts(nums, 1))
}
