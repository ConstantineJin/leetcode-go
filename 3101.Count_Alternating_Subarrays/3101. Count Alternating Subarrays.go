package main

func countAlternatingSubarrays(nums []int) (ans int64) {
	var cnt int64
	for i, num := range nums {
		if i > 0 && num == nums[i-1] {
			cnt = 1
		} else {
			cnt++
		}
		ans += cnt
	}
	return
}
