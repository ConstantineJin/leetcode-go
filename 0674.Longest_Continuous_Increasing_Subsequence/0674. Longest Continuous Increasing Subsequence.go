package main

func findLengthOfLCIS(nums []int) (ans int) {
	if len(nums) == 1 {
		return 1
	}
	var temp = 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			temp++
			ans = max(ans, temp)
		} else {
			ans = max(ans, temp)
			temp = 1
		}
	}
	return
}
