package main

import "slices"

func minimumTime(nums1, nums2 []int, x int) int {
	var sum1, sum2 int
	var n = len(nums1)
	var id, f = make([]int, n), make([]int, n+1)
	for i := range id {
		id[i] = i
		sum1 += nums1[i]
		sum2 += nums2[i]
	}
	// 对下标数组排序，避免破坏 nums1 和 nums2 的对应关系
	slices.SortFunc(id, func(i, j int) int { return nums2[i] - nums2[j] })
	for i, p := range id {
		var a, b = nums1[p], nums2[p]
		for j := i + 1; j > 0; j-- {
			f[j] = max(f[j], f[j-1]+a+b*j)
		}
	}
	for t, v := range f {
		if sum1+sum2*t-v <= x {
			return t
		}
	}
	return -1
}
