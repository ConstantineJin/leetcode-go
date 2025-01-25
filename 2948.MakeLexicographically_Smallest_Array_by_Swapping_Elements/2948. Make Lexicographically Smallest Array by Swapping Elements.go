package main

import "slices"

func lexicographicallySmallestArray(nums []int, limit int) []int {
	n := len(nums)
	ids := make([]int, n)
	for i := range n {
		ids[i] = i
	}
	slices.SortFunc(ids, func(i, j int) int { return nums[i] - nums[j] })
	ans := make([]int, n)
	for i := 0; i < n; {
		start := i
		for i++; i < n && nums[ids[i]]-nums[ids[i-1]] <= limit; i++ {
		}
		subIds := slices.Clone(ids[start:i])
		slices.Sort(subIds)
		for j, idx := range subIds {
			ans[idx] = nums[ids[start+j]]
		}
	}
	return ans
}
