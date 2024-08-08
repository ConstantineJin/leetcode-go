package main

import "slices"

func minimumAddedInteger(nums1, nums2 []int) int {
	slices.Sort(nums1)
	slices.Sort(nums2)
	for i := 2; i > 0; i-- { // 倒序枚举从 nums1[2] 或 nums1[1] 或 nums1[0] 开始保留，因为 nums1[i] 越大答案越小，第一个满足的就是答案
		x := nums2[0] - nums1[i]
		j := 0
		for _, v := range nums1[i:] {
			if nums2[j] == v+x {
				j++
				if j == len(nums2) { // nums2 是 {nums1[i] + x} 的子序列
					return x
				}
			}
		}
	}
	return nums2[0] - nums1[0]
}
