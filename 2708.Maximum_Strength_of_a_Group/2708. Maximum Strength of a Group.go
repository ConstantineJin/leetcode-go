package main

func maxStrength(nums []int) int64 {
	mn, mx := nums[0], nums[0]
	for _, num := range nums[1:] {
		mn, mx = min(mn, num, mn*num, mx*num), max(mx, num, mn*num, mx*num)
	}
	return int64(mx)
}
