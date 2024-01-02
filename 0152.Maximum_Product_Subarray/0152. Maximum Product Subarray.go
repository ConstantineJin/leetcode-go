package main

func maxProduct(nums []int) int {
	var maxF, minF, ans = nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		var mx, mn = maxF, minF
		maxF, minF = max(mx*nums[i], max(nums[i], mn*nums[i])), min(mn*nums[i], min(nums[i], mx*nums[i]))
		ans = max(maxF, ans)
	}
	return ans
}
