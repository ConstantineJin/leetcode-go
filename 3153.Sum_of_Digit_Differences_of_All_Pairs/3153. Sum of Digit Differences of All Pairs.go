package main

import "strconv"

func sumDigitDifferences(nums []int) int64 {
	m, n := len(strconv.Itoa(nums[0])), len(nums)
	var ans int
	for i := 0; i < m; i++ {
		var cnt [10]int
		for j, num := range nums {
			cnt[num%10]++
			nums[j] /= 10
		}
		for _, v := range cnt {
			ans += v * (n - v)
		}
	}
	return int64(ans) / 2
}
