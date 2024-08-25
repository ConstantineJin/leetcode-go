package main

func getFinalState(nums []int, k int, multiplier int) []int {
	for i := 0; i < k; i++ {
		mn, idx := nums[0], 0
		for j, num := range nums {
			if num < mn {
				mn = num
				idx = j
			}
		}
		nums[idx] *= multiplier
	}
	return nums
}
