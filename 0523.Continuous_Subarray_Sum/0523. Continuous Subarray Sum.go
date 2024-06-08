package main

func checkSubarraySum(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}
	mp := map[int]int{0: -1}
	var remainder int
	for i, num := range nums {
		remainder = (remainder + num) % k
		if prevIndex, ok := mp[remainder]; ok {
			if i-prevIndex > 1 {
				return true
			}
		} else {
			mp[remainder] = i
		}
	}
	return false
}
