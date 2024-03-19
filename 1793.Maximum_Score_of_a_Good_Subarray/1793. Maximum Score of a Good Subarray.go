package main

func maximumScore(nums []int, k int) int {
	var n, i, j, mn, ans = len(nums), k, k, nums[k], nums[k]
	for l := 0; l < n-1; l++ {
		if j == n-1 || i > 0 && nums[i-1] > nums[j+1] {
			i--
			mn = min(mn, nums[i])
		} else {
			j++
			mn = min(mn, nums[j])
		}
		ans = max(ans, mn*(j-i+1))
	}
	return ans
}
