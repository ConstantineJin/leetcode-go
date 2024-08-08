package main

import "slices"

func addedInteger(nums1, nums2 []int) int {
	return slices.Min(nums2) - slices.Min(nums1)
}
