package main

import "sort"

func divideArray(nums []int, k int) [][]int {
	var ans = make([][]int, len(nums)/3)
	for i := range ans {
		ans[i] = make([]int, 3)
	}
	sort.Ints(nums)
	for i := 0; i < len(nums); i += 3 {
		if nums[i+2]-nums[i] > k {
			return make([][]int, 0)
		}
		ans[i/3] = []int{nums[i], nums[i+1], nums[i+2]}
	}
	return ans
}
