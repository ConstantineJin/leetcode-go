package main

import "strconv"

func summaryRanges(nums []int) (ans []string) {
	n, i, j := len(nums), 0, 0
	for i < n {
		for j = i + 1; j < n && nums[j] == nums[j-1]+1; j++ {
		}
		if j == i+1 {
			ans = append(ans, strconv.Itoa(nums[i]))
		} else {
			ans = append(ans, strconv.Itoa(nums[i])+"->"+strconv.Itoa(nums[j-1]))
		}
		i = j
	}
	return
}
