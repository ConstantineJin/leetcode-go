package main

import "sort"

func largestDivisibleSubset(nums []int) (ans []int) {
	sort.Ints(nums)
	var n = len(nums)
	var f = make([][]int, n)
	for i := 0; i < n; i++ {
		var flag bool
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && len(f[j]) >= len(f[i]) {
				f[i], flag = append(append([]int{}, f[j]...), nums[i]), true
			}
		}
		if !flag {
			f[i] = []int{nums[i]}
		}
	}
	for _, arr := range f {
		if len(arr) > len(ans) {
			ans = arr
		}
	}
	return
}
