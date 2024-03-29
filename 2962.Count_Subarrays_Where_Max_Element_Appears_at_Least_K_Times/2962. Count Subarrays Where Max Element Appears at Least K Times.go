package main

import "slices"

func countSubarrays(nums []int, k int) (ans int64) {
	var mx = slices.Max(nums)
	var i, cnt int
	for _, num := range nums {
		if num == mx {
			cnt++
		}
		for ; cnt == k; i++ {
			if nums[i] == mx {
				cnt--
			}
		}
		ans += int64(i)
	}
	return
}
