package main

func resultsArray(nums []int, k int) []int {
	n := len(nums)
	ans := make([]int, n-k+1)
	curLen := 1
	for i := 1; i < k; i++ {
		if nums[i]-nums[i-1] == 1 {
			curLen++
		} else {
			curLen = 1
		}
	}
	if curLen == k {
		ans[0] = nums[k-1]
	} else {
		ans[0] = -1
	}
	for i := k; i < n; i++ {
		if nums[i]-nums[i-1] == 1 {
			curLen++
		} else {
			curLen = 1
		}
		if curLen >= k {
			ans[i-k+1] = nums[i]
		} else {
			ans[i-k+1] = -1
		}
	}
	return ans
}
