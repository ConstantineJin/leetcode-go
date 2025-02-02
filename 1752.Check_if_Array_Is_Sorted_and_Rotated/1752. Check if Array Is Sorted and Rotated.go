package main

func check(nums []int) bool {
	n := len(nums)
	x := 0
	for i := 1; i < n; i++ {
		if nums[i] < nums[i-1] {
			x = i
			break
		}
	}
	if x == 0 {
		return true
	}
	for i := x + 1; i < n; i++ {
		if nums[i] < nums[i-1] {
			return false
		}
	}
	return nums[0] >= nums[n-1]
}
