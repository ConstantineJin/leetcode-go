package main

func longestMonotonicSubarray(nums []int) int {
	ans, n := 1, len(nums)
	for i := 0; i < n-1; {
		if nums[i+1] == nums[i] {
			i++
			continue
		}
		i0 := i
		increase := nums[i] < nums[i+1]
		for i += 2; i < n && nums[i] != nums[i-1] && nums[i] > nums[i-1] == increase; i++ {
		}
		ans = max(ans, i-i0)
		i--
	}
	return ans
}
