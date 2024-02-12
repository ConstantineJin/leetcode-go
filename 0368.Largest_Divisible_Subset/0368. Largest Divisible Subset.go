package main

import "sort"

func largestDivisibleSubset(nums []int) (ans []int) {
	sort.Ints(nums)
	var n = len(nums)
	var f = make([][]int, n)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && len(f[j]) >= len(f[i]) {
				f[i] = append(append([]int{}, f[j]...), nums[i])
			}
		}
		if len(f[i]) == 0 {
			f[i] = []int{nums[i]}
		}
		if len(f[i]) > len(ans) {
			ans = f[i]
		}
	}
	return
}
