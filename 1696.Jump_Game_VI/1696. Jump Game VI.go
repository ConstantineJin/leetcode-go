package main

func maxResult(nums []int, k int) int {
	var n = len(nums)
	var f = make([]int, n)
	f[0] = nums[0]
	var q = []int{0} // 单调队列，从左到右严格递减
	for i := 1; i < n; i++ {
		if q[0] < i-k {
			q = q[1:]
		}
		f[i] = f[q[0]] + nums[i]
		for len(q) > 0 && f[i] >= f[q[len(q)-1]] {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	return f[n-1]
}
