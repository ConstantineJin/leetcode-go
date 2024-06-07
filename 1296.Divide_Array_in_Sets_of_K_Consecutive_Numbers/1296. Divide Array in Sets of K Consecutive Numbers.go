package main

import "sort"

func isPossibleDivide(nums []int, k int) bool {
	n := len(nums)
	if n%k > 0 {
		return false
	}
	mp := make(map[int]int)
	for _, num := range nums {
		mp[num]++
	}
	sort.Ints(nums)
	for _, num := range nums {
		if mp[num] == 0 {
			continue
		}
		for i := num; i < num+k; i++ {
			if mp[i] <= 0 {
				return false
			}
			mp[i]--
		}
	}
	return true
}
