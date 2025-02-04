package main

func maxAscendingSum(nums []int) (ans int) {
	n := len(nums)
	for i := 0; i < n; {
		sum := nums[i]
		for i++; i < n && nums[i-1] < nums[i]; i++ {
			sum += nums[i]
		}
		ans = max(ans, sum)
	}
	return
}
