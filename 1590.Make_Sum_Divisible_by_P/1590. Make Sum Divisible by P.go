package main

func minSubarray(nums []int, p int) int {
	n := len(nums)
	prefixSum := make([]int, n+1)
	for i, num := range nums {
		prefixSum[i+1] = (prefixSum[i] + num) % p
	}
	if prefixSum[n] == 0 {
		return 0
	}
	ans := n
	last := make(map[int]int)
	for i, s := range prefixSum {
		last[s] = i
		if j, ok := last[(s-prefixSum[n]+p)%p]; ok {
			ans = min(ans, i-j)
		}
	}
	if ans == n {
		return -1
	}
	return ans
}
