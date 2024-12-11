package main

import "slices"

func semiOrderedPermutation(nums []int) int {
	n := len(nums)
	i, j := slices.Index(nums, 1), slices.Index(nums, n)
	if i > j {
		return i + n - j - 2
	}
	return i + n - j - 1
}
