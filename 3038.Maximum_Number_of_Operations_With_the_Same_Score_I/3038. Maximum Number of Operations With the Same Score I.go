package main

func maxOperations(nums []int) (ans int) {
	for ; ans < len(nums)-len(nums)%2-1 && nums[ans]+nums[ans+1] == nums[0]+nums[1]; ans += 2 {
	}
	return ans / 2
}
