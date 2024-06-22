package main

func maximumTripletValue(nums []int) int64 {
	n := len(nums)
	left, right := make([]int, n), make([]int, n)
	mx := nums[0]
	for i := 1; i < n; i++ {
		left[i] = mx - nums[i]
		mx = max(mx, nums[i])
	}
	mx = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		right[i] = mx
		mx = max(mx, nums[i])
	}
	var ans int
	for i := 1; i < n-1; i++ {
		ans = max(ans, left[i]*right[i])
	}
	return int64(ans)
}
