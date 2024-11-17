package main

func shortestSubarray(nums []int, k int) int {
	n := len(nums)
	prefixSums := make([]int, n+1)
	for i, num := range nums {
		prefixSums[i+1] = prefixSums[i] + num
	}
	ans := n + 1
	var q []int
	for i, sum := range prefixSums {
		for len(q) > 0 && sum-prefixSums[q[0]] >= k {
			ans = min(ans, i-q[0])
			q = q[1:]
		}
		for len(q) > 0 && prefixSums[q[len(q)-1]] >= sum {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	if ans > n {
		return -1
	}
	return ans
}
