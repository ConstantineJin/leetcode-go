package main

import (
	"math/bits"
	"slices"
)

func canSortArray(nums []int) bool {
	n := len(nums)
	for i := 0; i < n; {
		start, cnt := i, bits.OnesCount(uint(nums[i]))
		i++
		for i < n && bits.OnesCount(uint(nums[i])) == cnt {
			i++
		}
		slices.Sort(nums[start:i])
	}
	return slices.IsSorted(nums)
}
