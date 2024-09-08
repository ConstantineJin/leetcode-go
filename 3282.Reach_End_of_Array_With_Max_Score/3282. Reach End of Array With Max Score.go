package main

func findMaximumScore(nums []int) (ans int64) {
	n := len(nums)
	i, mx := 0, nums[0]
	for j := 1; j < n; j++ {
		if nums[j] > mx {
			ans += int64((j - i) * mx)
			i, mx = j, nums[j]
		}
	}
	ans += int64((n - 1 - i) * mx)
	return
}
