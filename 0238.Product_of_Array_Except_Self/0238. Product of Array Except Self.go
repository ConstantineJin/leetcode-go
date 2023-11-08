package main

func productExceptSelf(nums []int) []int {
	n := len(nums)
	left, right, res := make([]int, n), make([]int, n), make([]int, n) // 前缀积与后缀积
	left[0], right[n-1] = 1, 1
	for i := 1; i < n; i++ {
		left[i] = left[i-1] * nums[i-1]
	}
	for i := n - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}
	for i := 0; i < n; i++ {
		res[i] = left[i] * right[i]
	}
	return res
}
