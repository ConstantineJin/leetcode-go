package main

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	slow := 2
	for fast := 2; fast < len(nums); fast++ {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
