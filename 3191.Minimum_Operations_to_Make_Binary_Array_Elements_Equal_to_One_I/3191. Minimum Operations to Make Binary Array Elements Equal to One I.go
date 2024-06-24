package main

func minOperations(nums []int) (ans int) {
	n := len(nums)
	for i, num := range nums[:n-2] {
		if num == 0 {
			ans++
			nums[i+1] ^= 1
			nums[i+2] ^= 1
		}
	}
	if nums[n-2] == 0 || nums[n-1] == 0 {
		return -1
	}
	return
}
