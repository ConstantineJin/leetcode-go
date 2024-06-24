package main

func minOperations(nums []int) (ans int) {
	n := len(nums)
	for i := n - 1; i > 0; i-- {
		if nums[i]+nums[i-1] == 1 {
			ans++
		}
	}
	if nums[0] == 0 {
		ans++
	}
	return
}
